// Copyright 2020 The gVisor Authors.
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

package fuse

import (
	"github.com/nicocha30/gvisor-ligolo/pkg/abi/linux"
	"github.com/nicocha30/gvisor-ligolo/pkg/context"
	"github.com/nicocha30/gvisor-ligolo/pkg/sentry/fsimpl/devtmpfs"
	"github.com/nicocha30/gvisor-ligolo/pkg/sentry/vfs"
)

// Register registers the FUSE device with vfsObj.
func Register(vfsObj *vfs.VirtualFilesystem) error {
	if err := vfsObj.RegisterDevice(vfs.CharDevice, linux.MISC_MAJOR, fuseDevMinor, fuseDevice{}, &vfs.RegisterDeviceOptions{
		GroupName: "misc",
	}); err != nil {
		return err
	}

	return nil
}

// CreateDevtmpfsFile creates a device special file in devtmpfs.
func CreateDevtmpfsFile(ctx context.Context, dev *devtmpfs.Accessor) error {
	if err := dev.CreateDeviceFile(ctx, "fuse", vfs.CharDevice, linux.MISC_MAJOR, fuseDevMinor, 0666 /* mode */); err != nil {
		return err
	}

	return nil
}
