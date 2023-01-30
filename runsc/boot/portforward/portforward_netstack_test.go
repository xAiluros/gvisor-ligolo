// Copyright 2023 The gVisor Authors.
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
	"bytes"
	"io"
	"sync"
	"testing"

	"github.com/nicocha30/gvisor-ligolo/pkg/cleanup"
	"github.com/nicocha30/gvisor-ligolo/pkg/sentry/contexttest"
	"github.com/nicocha30/gvisor-ligolo/pkg/sentry/vfs"
	"github.com/nicocha30/gvisor-ligolo/pkg/tcpip"
	"github.com/nicocha30/gvisor-ligolo/pkg/waiter"
)

type baseTCPEndpointImpl struct {
	closed   bool
	readBuf  bytes.Buffer
	writeBuf bytes.Buffer
	mu       sync.Mutex
	wq       *waiter.Queue
}

// read reads data from the buffer that "Write" writes to.
func (b *baseTCPEndpointImpl) read(n int) ([]byte, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return nil, io.EOF
	}
	ret := b.writeBuf.Next(n)
	b.wq.Notify(waiter.WritableEvents)
	return ret, nil
}

// write writes data to the read buffer that "Read" reads from.
func (b *baseTCPEndpointImpl) write(buf []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return 0, io.EOF
	}
	n, err := b.readBuf.Write(buf)
	b.wq.Notify(waiter.ReadableEvents)
	return n, err
}

func (b *baseTCPEndpointImpl) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.closed = true
}

func (b *baseTCPEndpointImpl) Read(w io.Writer, _ tcpip.ReadOptions) (tcpip.ReadResult, tcpip.Error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return tcpip.ReadResult{}, &tcpip.ErrClosedForReceive{}
	}
	buf := b.readBuf.Next(b.readBuf.Len())
	n, err := w.Write(buf)
	if err != nil {
		return tcpip.ReadResult{}, &tcpip.ErrInvalidEndpointState{}
	}
	return tcpip.ReadResult{
		Count: n,
		Total: n,
	}, nil
}

func (b *baseTCPEndpointImpl) Write(payload tcpip.Payloader, _ tcpip.WriteOptions) (int64, tcpip.Error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return 0, &tcpip.ErrClosedForSend{}
	}
	buf := make([]byte, payload.Len())
	n, err := payload.Read(buf)
	if err != nil {
		return 0, &tcpip.ErrInvalidEndpointState{}
	}
	n, err = b.writeBuf.Write(buf[:n])
	if err != nil {
		return int64(n), &tcpip.ErrConnectionRefused{}
	}
	return int64(n), nil
}

func (b *baseTCPEndpointImpl) Shutdown(shutdown tcpip.ShutdownFlags) tcpip.Error {
	return nil
}

func newNetstackPortForwardConnWithMock(impl mockTCPEndpointImpl, fd *vfs.FileDescription, wq *waiter.Queue) *netstackPortForwardConn {
	ep := &mockTCPEndpoint{impl}
	return &netstackPortForwardConn{
		ep:       ep,
		wq:       wq,
		fd:       fd,
		toDone:   make(chan struct{}),
		fromDone: make(chan struct{}),
		cu:       cleanup.Cleanup{},
	}
}

func TestNetstackPortforward(t *testing.T) {
	for _, tc := range []struct {
		name     string
		requests map[string]string
	}{
		{
			name: "single",
			requests: map[string]string{
				"PING": "PONG",
			},
		},
		{
			name: "multiple",
			requests: map[string]string{
				"PING":       "PONG",
				"HELLO":      "GOODBYE",
				"IMPRESSIVE": "MOST IMPRESSIVE",
			},
		},
		{
			name: "empty",
			requests: map[string]string{
				"EMPTY":       "",
				"NOT":         "EMPTY",
				"OTHER EMPTY": "",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			doNetstackTest(t, tc.requests)
		})
	}
}

func doNetstackTest(t *testing.T, responses map[string]string) {
	ctx := contexttest.Context(t)
	appEndpoint := &mockApplicationFDImpl{}
	defer appEndpoint.Release(ctx)
	fd, err := newMockFileDescription(ctx, appEndpoint)
	if err != nil {
		t.Fatalf("newMockFileDescription: %v", err)
	}

	wq := waiter.Queue{}
	impl := &baseTCPEndpointImpl{wq: &wq}
	conn := newNetstackPortForwardConnWithMock(impl, fd, &wq)
	if err := conn.start(ctx); err != nil {
		t.Fatalf("conn.start: %v", err)
	}
	defer conn.close(ctx)

	harness := portforwarderTestHarness{
		app:  appEndpoint,
		shim: impl,
	}

	for req, resp := range responses {
		if _, err := harness.shimWrite([]byte(req)); err != nil {
			t.Fatalf("failed to write to shim: %v", err)
		}

		got, err := harness.appRead(len(req))
		if err != nil {
			t.Fatalf("failed to read from app: %v", err)
		}

		if string(got) != req {
			t.Fatalf("app mismatch: got: %s want: %s", string(got), req)
		}

		if _, err := harness.appWrite([]byte(resp)); err != nil {
			t.Fatalf("failed to write to app: %v", err)
		}

		got, err = harness.shimRead(len(resp))
		if err != nil {
			t.Fatalf("failed to read from shim: %v", err)
		}

		if string(got) != resp {
			t.Fatalf("shim mismatch: got: %s want: %s", string(got), resp)
		}
	}
}
