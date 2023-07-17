// automatically generated by stateify.

package pipefs

import (
	"github.com/nicocha30/gvisor-ligolo/pkg/state"
)

func (f *filesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/pipefs.filesystemType"
}

func (f *filesystemType) StateFields() []string {
	return []string{}
}

func (f *filesystemType) beforeSave() {}

// +checklocksignore
func (f *filesystemType) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
}

func (f *filesystemType) afterLoad() {}

// +checklocksignore
func (f *filesystemType) StateLoad(stateSourceObject state.Source) {
}

func (fs *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/pipefs.filesystem"
}

func (fs *filesystem) StateFields() []string {
	return []string{
		"Filesystem",
		"devMinor",
	}
}

func (fs *filesystem) beforeSave() {}

// +checklocksignore
func (fs *filesystem) StateSave(stateSinkObject state.Sink) {
	fs.beforeSave()
	stateSinkObject.Save(0, &fs.Filesystem)
	stateSinkObject.Save(1, &fs.devMinor)
}

func (fs *filesystem) afterLoad() {}

// +checklocksignore
func (fs *filesystem) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fs.Filesystem)
	stateSourceObject.Load(1, &fs.devMinor)
}

func (i *inode) StateTypeName() string {
	return "pkg/sentry/fsimpl/pipefs.inode"
}

func (i *inode) StateFields() []string {
	return []string{
		"InodeNotDirectory",
		"InodeNotSymlink",
		"InodeNoopRefCount",
		"InodeWatches",
		"locks",
		"pipe",
		"ino",
		"uid",
		"gid",
		"ctime",
	}
}

func (i *inode) beforeSave() {}

// +checklocksignore
func (i *inode) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.InodeNotDirectory)
	stateSinkObject.Save(1, &i.InodeNotSymlink)
	stateSinkObject.Save(2, &i.InodeNoopRefCount)
	stateSinkObject.Save(3, &i.InodeWatches)
	stateSinkObject.Save(4, &i.locks)
	stateSinkObject.Save(5, &i.pipe)
	stateSinkObject.Save(6, &i.ino)
	stateSinkObject.Save(7, &i.uid)
	stateSinkObject.Save(8, &i.gid)
	stateSinkObject.Save(9, &i.ctime)
}

func (i *inode) afterLoad() {}

// +checklocksignore
func (i *inode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.InodeNotDirectory)
	stateSourceObject.Load(1, &i.InodeNotSymlink)
	stateSourceObject.Load(2, &i.InodeNoopRefCount)
	stateSourceObject.Load(3, &i.InodeWatches)
	stateSourceObject.Load(4, &i.locks)
	stateSourceObject.Load(5, &i.pipe)
	stateSourceObject.Load(6, &i.ino)
	stateSourceObject.Load(7, &i.uid)
	stateSourceObject.Load(8, &i.gid)
	stateSourceObject.Load(9, &i.ctime)
}

func init() {
	state.Register((*filesystemType)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*inode)(nil))
}