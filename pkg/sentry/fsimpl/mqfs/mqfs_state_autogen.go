// automatically generated by stateify.

package mqfs

import (
	"github.com/nicocha30/gvisor-ligolo/pkg/state"
)

func (ft *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/mqfs.FilesystemType"
}

func (ft *FilesystemType) StateFields() []string {
	return []string{}
}

func (ft *FilesystemType) beforeSave() {}

// +checklocksignore
func (ft *FilesystemType) StateSave(stateSinkObject state.Sink) {
	ft.beforeSave()
}

func (ft *FilesystemType) afterLoad() {}

// +checklocksignore
func (ft *FilesystemType) StateLoad(stateSourceObject state.Source) {
}

func (fs *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/mqfs.filesystem"
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

func (q *queueInode) StateTypeName() string {
	return "pkg/sentry/fsimpl/mqfs.queueInode"
}

func (q *queueInode) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"queue",
	}
}

func (q *queueInode) beforeSave() {}

// +checklocksignore
func (q *queueInode) StateSave(stateSinkObject state.Sink) {
	q.beforeSave()
	stateSinkObject.Save(0, &q.DynamicBytesFile)
	stateSinkObject.Save(1, &q.queue)
}

func (q *queueInode) afterLoad() {}

// +checklocksignore
func (q *queueInode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &q.DynamicBytesFile)
	stateSourceObject.Load(1, &q.queue)
}

func (fd *queueFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/mqfs.queueFD"
}

func (fd *queueFD) StateFields() []string {
	return []string{
		"FileDescriptionDefaultImpl",
		"DynamicBytesFileDescriptionImpl",
		"LockFD",
		"vfsfd",
		"inode",
		"queue",
	}
}

func (fd *queueFD) beforeSave() {}

// +checklocksignore
func (fd *queueFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(1, &fd.DynamicBytesFileDescriptionImpl)
	stateSinkObject.Save(2, &fd.LockFD)
	stateSinkObject.Save(3, &fd.vfsfd)
	stateSinkObject.Save(4, &fd.inode)
	stateSinkObject.Save(5, &fd.queue)
}

func (fd *queueFD) afterLoad() {}

// +checklocksignore
func (fd *queueFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(1, &fd.DynamicBytesFileDescriptionImpl)
	stateSourceObject.Load(2, &fd.LockFD)
	stateSourceObject.Load(3, &fd.vfsfd)
	stateSourceObject.Load(4, &fd.inode)
	stateSourceObject.Load(5, &fd.queue)
}

func (r *RegistryImpl) StateTypeName() string {
	return "pkg/sentry/fsimpl/mqfs.RegistryImpl"
}

func (r *RegistryImpl) StateFields() []string {
	return []string{
		"root",
		"fs",
		"mount",
	}
}

func (r *RegistryImpl) beforeSave() {}

// +checklocksignore
func (r *RegistryImpl) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.root)
	stateSinkObject.Save(1, &r.fs)
	stateSinkObject.Save(2, &r.mount)
}

func (r *RegistryImpl) afterLoad() {}

// +checklocksignore
func (r *RegistryImpl) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.root)
	stateSourceObject.Load(1, &r.fs)
	stateSourceObject.Load(2, &r.mount)
}

func (i *rootInode) StateTypeName() string {
	return "pkg/sentry/fsimpl/mqfs.rootInode"
}

func (i *rootInode) StateFields() []string {
	return []string{
		"rootInodeRefs",
		"InodeAlwaysValid",
		"InodeAttrs",
		"InodeDirectoryNoNewChildren",
		"InodeNotSymlink",
		"InodeTemporary",
		"InodeWatches",
		"OrderedChildren",
		"locks",
	}
}

func (i *rootInode) beforeSave() {}

// +checklocksignore
func (i *rootInode) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.rootInodeRefs)
	stateSinkObject.Save(1, &i.InodeAlwaysValid)
	stateSinkObject.Save(2, &i.InodeAttrs)
	stateSinkObject.Save(3, &i.InodeDirectoryNoNewChildren)
	stateSinkObject.Save(4, &i.InodeNotSymlink)
	stateSinkObject.Save(5, &i.InodeTemporary)
	stateSinkObject.Save(6, &i.InodeWatches)
	stateSinkObject.Save(7, &i.OrderedChildren)
	stateSinkObject.Save(8, &i.locks)
}

func (i *rootInode) afterLoad() {}

// +checklocksignore
func (i *rootInode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.rootInodeRefs)
	stateSourceObject.Load(1, &i.InodeAlwaysValid)
	stateSourceObject.Load(2, &i.InodeAttrs)
	stateSourceObject.Load(3, &i.InodeDirectoryNoNewChildren)
	stateSourceObject.Load(4, &i.InodeNotSymlink)
	stateSourceObject.Load(5, &i.InodeTemporary)
	stateSourceObject.Load(6, &i.InodeWatches)
	stateSourceObject.Load(7, &i.OrderedChildren)
	stateSourceObject.Load(8, &i.locks)
}

func (r *rootInodeRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/mqfs.rootInodeRefs"
}

func (r *rootInodeRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *rootInodeRefs) beforeSave() {}

// +checklocksignore
func (r *rootInodeRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *rootInodeRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func init() {
	state.Register((*FilesystemType)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*queueInode)(nil))
	state.Register((*queueFD)(nil))
	state.Register((*RegistryImpl)(nil))
	state.Register((*rootInode)(nil))
	state.Register((*rootInodeRefs)(nil))
}
