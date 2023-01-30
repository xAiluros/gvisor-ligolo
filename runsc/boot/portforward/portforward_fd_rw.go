// Copyright 2022 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package portforward

import (
	"io"

	"github.com/nicocha30/gvisor-ligolo/pkg/context"
	"github.com/nicocha30/gvisor-ligolo/pkg/errors/linuxerr"
	"github.com/nicocha30/gvisor-ligolo/pkg/sentry/vfs"
	"github.com/nicocha30/gvisor-ligolo/pkg/usermem"
	"github.com/nicocha30/gvisor-ligolo/pkg/waiter"
)

// fileDescriptionReadWriter implements io.ReadWriter and allows reading and
// writing to a vfs.FileDescription.
type fileDescriptionReadWriter struct {
	// ctx is the context for the socket reader.
	ctx context.Context

	// file is the file to read and write from.
	file *vfs.FileDescription
}

// Read implements io.Reader.Read. It performs a blocking read on the fd.
func (r *fileDescriptionReadWriter) Read(buf []byte) (int, error) {
	var (
		notifyCh  chan struct{}
		waitEntry waiter.Entry
	)
	n, err := r.file.Read(r.ctx, usermem.BytesIOSequence(buf), vfs.ReadOptions{})
	for linuxerr.Equals(linuxerr.ErrWouldBlock, err) {
		if notifyCh == nil {
			waitEntry, notifyCh = waiter.NewChannelEntry(waiter.ReadableEvents | waiter.WritableEvents | waiter.EventHUp | waiter.EventErr)
			// Register for when the endpoint is readable or disconnected.
			r.file.EventRegister(&waitEntry)
			defer r.file.EventUnregister(&waitEntry)
		}
		<-notifyCh
		n, err = r.file.Read(r.ctx, usermem.BytesIOSequence(buf), vfs.ReadOptions{})
	}

	// host fd FileDescriptions use recvmsg which returns zero when the
	// peer has shutdown. When that happens return EOF.
	if n == 0 && err == nil {
		return 0, io.EOF
	}
	return int(n), err
}

// Write implements io.Writer.Write. It performs a blocking write on the fd.
func (r *fileDescriptionReadWriter) Write(buf []byte) (int, error) {
	var notifyCh chan struct{}
	var waitEntry waiter.Entry
	n, err := r.file.Write(r.ctx, usermem.BytesIOSequence(buf), vfs.WriteOptions{})
	for linuxerr.Equals(linuxerr.ErrWouldBlock, err) {
		if notifyCh == nil {
			waitEntry, notifyCh = waiter.NewChannelEntry(waiter.WritableEvents | waiter.EventHUp | waiter.EventErr)
			// Register for when the endpoint is writable or disconnected.
			r.file.EventRegister(&waitEntry)
			defer r.file.EventUnregister(&waitEntry)
		}
		<-notifyCh
		n, err = r.file.Write(r.ctx, usermem.BytesIOSequence(buf), vfs.WriteOptions{})
	}
	return int(n), err
}
