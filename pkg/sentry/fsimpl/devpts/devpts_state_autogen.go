// automatically generated by stateify.

package devpts

import (
	"github.com/nicocha30/gvisor-ligolo/pkg/state"
)

func (fstype *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/devpts.FilesystemType"
}

func (fstype *FilesystemType) StateFields() []string {
	return []string{
		"initErr",
		"fs",
		"root",
	}
}

func (fstype *FilesystemType) beforeSave() {}

// +checklocksignore
func (fstype *FilesystemType) StateSave(stateSinkObject state.Sink) {
	fstype.beforeSave()
	stateSinkObject.Save(0, &fstype.initErr)
	stateSinkObject.Save(1, &fstype.fs)
	stateSinkObject.Save(2, &fstype.root)
}

func (fstype *FilesystemType) afterLoad() {}

// +checklocksignore
func (fstype *FilesystemType) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fstype.initErr)
	stateSourceObject.Load(1, &fstype.fs)
	stateSourceObject.Load(2, &fstype.root)
}

func (fs *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/devpts.filesystem"
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

func (i *rootInode) StateTypeName() string {
	return "pkg/sentry/fsimpl/devpts.rootInode"
}

func (i *rootInode) StateFields() []string {
	return []string{
		"implStatFS",
		"InodeAlwaysValid",
		"InodeAttrs",
		"InodeDirectoryNoNewChildren",
		"InodeNotSymlink",
		"InodeTemporary",
		"InodeWatches",
		"OrderedChildren",
		"rootInodeRefs",
		"locks",
		"master",
		"replicas",
		"nextIdx",
	}
}

func (i *rootInode) beforeSave() {}

// +checklocksignore
func (i *rootInode) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.implStatFS)
	stateSinkObject.Save(1, &i.InodeAlwaysValid)
	stateSinkObject.Save(2, &i.InodeAttrs)
	stateSinkObject.Save(3, &i.InodeDirectoryNoNewChildren)
	stateSinkObject.Save(4, &i.InodeNotSymlink)
	stateSinkObject.Save(5, &i.InodeTemporary)
	stateSinkObject.Save(6, &i.InodeWatches)
	stateSinkObject.Save(7, &i.OrderedChildren)
	stateSinkObject.Save(8, &i.rootInodeRefs)
	stateSinkObject.Save(9, &i.locks)
	stateSinkObject.Save(10, &i.master)
	stateSinkObject.Save(11, &i.replicas)
	stateSinkObject.Save(12, &i.nextIdx)
}

func (i *rootInode) afterLoad() {}

// +checklocksignore
func (i *rootInode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.implStatFS)
	stateSourceObject.Load(1, &i.InodeAlwaysValid)
	stateSourceObject.Load(2, &i.InodeAttrs)
	stateSourceObject.Load(3, &i.InodeDirectoryNoNewChildren)
	stateSourceObject.Load(4, &i.InodeNotSymlink)
	stateSourceObject.Load(5, &i.InodeTemporary)
	stateSourceObject.Load(6, &i.InodeWatches)
	stateSourceObject.Load(7, &i.OrderedChildren)
	stateSourceObject.Load(8, &i.rootInodeRefs)
	stateSourceObject.Load(9, &i.locks)
	stateSourceObject.Load(10, &i.master)
	stateSourceObject.Load(11, &i.replicas)
	stateSourceObject.Load(12, &i.nextIdx)
}

func (i *implStatFS) StateTypeName() string {
	return "pkg/sentry/fsimpl/devpts.implStatFS"
}

func (i *implStatFS) StateFields() []string {
	return []string{}
}

func (i *implStatFS) beforeSave() {}

// +checklocksignore
func (i *implStatFS) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
}

func (i *implStatFS) afterLoad() {}

// +checklocksignore
func (i *implStatFS) StateLoad(stateSourceObject state.Source) {
}

func (l *lineDiscipline) StateTypeName() string {
	return "pkg/sentry/fsimpl/devpts.lineDiscipline"
}

func (l *lineDiscipline) StateFields() []string {
	return []string{
		"size",
		"inQueue",
		"outQueue",
		"termios",
		"column",
		"masterWaiter",
		"replicaWaiter",
	}
}

func (l *lineDiscipline) beforeSave() {}

// +checklocksignore
func (l *lineDiscipline) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.size)
	stateSinkObject.Save(1, &l.inQueue)
	stateSinkObject.Save(2, &l.outQueue)
	stateSinkObject.Save(3, &l.termios)
	stateSinkObject.Save(4, &l.column)
	stateSinkObject.Save(5, &l.masterWaiter)
	stateSinkObject.Save(6, &l.replicaWaiter)
}

func (l *lineDiscipline) afterLoad() {}

// +checklocksignore
func (l *lineDiscipline) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.size)
	stateSourceObject.Load(1, &l.inQueue)
	stateSourceObject.Load(2, &l.outQueue)
	stateSourceObject.Load(3, &l.termios)
	stateSourceObject.Load(4, &l.column)
	stateSourceObject.Load(5, &l.masterWaiter)
	stateSourceObject.Load(6, &l.replicaWaiter)
}

func (o *outputQueueTransformer) StateTypeName() string {
	return "pkg/sentry/fsimpl/devpts.outputQueueTransformer"
}

func (o *outputQueueTransformer) StateFields() []string {
	return []string{}
}

func (o *outputQueueTransformer) beforeSave() {}

// +checklocksignore
func (o *outputQueueTransformer) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
}

func (o *outputQueueTransformer) afterLoad() {}

// +checklocksignore
func (o *outputQueueTransformer) StateLoad(stateSourceObject state.Source) {
}

func (i *inputQueueTransformer) StateTypeName() string {
	return "pkg/sentry/fsimpl/devpts.inputQueueTransformer"
}

func (i *inputQueueTransformer) StateFields() []string {
	return []string{}
}

func (i *inputQueueTransformer) beforeSave() {}

// +checklocksignore
func (i *inputQueueTransformer) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
}

func (i *inputQueueTransformer) afterLoad() {}

// +checklocksignore
func (i *inputQueueTransformer) StateLoad(stateSourceObject state.Source) {
}

func (mi *masterInode) StateTypeName() string {
	return "pkg/sentry/fsimpl/devpts.masterInode"
}

func (mi *masterInode) StateFields() []string {
	return []string{
		"implStatFS",
		"InodeAttrs",
		"InodeNoopRefCount",
		"InodeNotDirectory",
		"InodeNotSymlink",
		"InodeWatches",
		"locks",
		"root",
	}
}

func (mi *masterInode) beforeSave() {}

// +checklocksignore
func (mi *masterInode) StateSave(stateSinkObject state.Sink) {
	mi.beforeSave()
	stateSinkObject.Save(0, &mi.implStatFS)
	stateSinkObject.Save(1, &mi.InodeAttrs)
	stateSinkObject.Save(2, &mi.InodeNoopRefCount)
	stateSinkObject.Save(3, &mi.InodeNotDirectory)
	stateSinkObject.Save(4, &mi.InodeNotSymlink)
	stateSinkObject.Save(5, &mi.InodeWatches)
	stateSinkObject.Save(6, &mi.locks)
	stateSinkObject.Save(7, &mi.root)
}

func (mi *masterInode) afterLoad() {}

// +checklocksignore
func (mi *masterInode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &mi.implStatFS)
	stateSourceObject.Load(1, &mi.InodeAttrs)
	stateSourceObject.Load(2, &mi.InodeNoopRefCount)
	stateSourceObject.Load(3, &mi.InodeNotDirectory)
	stateSourceObject.Load(4, &mi.InodeNotSymlink)
	stateSourceObject.Load(5, &mi.InodeWatches)
	stateSourceObject.Load(6, &mi.locks)
	stateSourceObject.Load(7, &mi.root)
}

func (mfd *masterFileDescription) StateTypeName() string {
	return "pkg/sentry/fsimpl/devpts.masterFileDescription"
}

func (mfd *masterFileDescription) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"LockFD",
		"inode",
		"t",
	}
}

func (mfd *masterFileDescription) beforeSave() {}

// +checklocksignore
func (mfd *masterFileDescription) StateSave(stateSinkObject state.Sink) {
	mfd.beforeSave()
	stateSinkObject.Save(0, &mfd.vfsfd)
	stateSinkObject.Save(1, &mfd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &mfd.LockFD)
	stateSinkObject.Save(3, &mfd.inode)
	stateSinkObject.Save(4, &mfd.t)
}

func (mfd *masterFileDescription) afterLoad() {}

// +checklocksignore
func (mfd *masterFileDescription) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &mfd.vfsfd)
	stateSourceObject.Load(1, &mfd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &mfd.LockFD)
	stateSourceObject.Load(3, &mfd.inode)
	stateSourceObject.Load(4, &mfd.t)
}

func (q *queue) StateTypeName() string {
	return "pkg/sentry/fsimpl/devpts.queue"
}

func (q *queue) StateFields() []string {
	return []string{
		"readBuf",
		"waitBuf",
		"waitBufLen",
		"readable",
		"transformer",
	}
}

func (q *queue) beforeSave() {}

// +checklocksignore
func (q *queue) StateSave(stateSinkObject state.Sink) {
	q.beforeSave()
	stateSinkObject.Save(0, &q.readBuf)
	stateSinkObject.Save(1, &q.waitBuf)
	stateSinkObject.Save(2, &q.waitBufLen)
	stateSinkObject.Save(3, &q.readable)
	stateSinkObject.Save(4, &q.transformer)
}

func (q *queue) afterLoad() {}

// +checklocksignore
func (q *queue) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &q.readBuf)
	stateSourceObject.Load(1, &q.waitBuf)
	stateSourceObject.Load(2, &q.waitBufLen)
	stateSourceObject.Load(3, &q.readable)
	stateSourceObject.Load(4, &q.transformer)
}

func (ri *replicaInode) StateTypeName() string {
	return "pkg/sentry/fsimpl/devpts.replicaInode"
}

func (ri *replicaInode) StateFields() []string {
	return []string{
		"implStatFS",
		"InodeAttrs",
		"InodeNoopRefCount",
		"InodeNotDirectory",
		"InodeNotSymlink",
		"InodeWatches",
		"locks",
		"root",
		"t",
	}
}

func (ri *replicaInode) beforeSave() {}

// +checklocksignore
func (ri *replicaInode) StateSave(stateSinkObject state.Sink) {
	ri.beforeSave()
	stateSinkObject.Save(0, &ri.implStatFS)
	stateSinkObject.Save(1, &ri.InodeAttrs)
	stateSinkObject.Save(2, &ri.InodeNoopRefCount)
	stateSinkObject.Save(3, &ri.InodeNotDirectory)
	stateSinkObject.Save(4, &ri.InodeNotSymlink)
	stateSinkObject.Save(5, &ri.InodeWatches)
	stateSinkObject.Save(6, &ri.locks)
	stateSinkObject.Save(7, &ri.root)
	stateSinkObject.Save(8, &ri.t)
}

func (ri *replicaInode) afterLoad() {}

// +checklocksignore
func (ri *replicaInode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &ri.implStatFS)
	stateSourceObject.Load(1, &ri.InodeAttrs)
	stateSourceObject.Load(2, &ri.InodeNoopRefCount)
	stateSourceObject.Load(3, &ri.InodeNotDirectory)
	stateSourceObject.Load(4, &ri.InodeNotSymlink)
	stateSourceObject.Load(5, &ri.InodeWatches)
	stateSourceObject.Load(6, &ri.locks)
	stateSourceObject.Load(7, &ri.root)
	stateSourceObject.Load(8, &ri.t)
}

func (rfd *replicaFileDescription) StateTypeName() string {
	return "pkg/sentry/fsimpl/devpts.replicaFileDescription"
}

func (rfd *replicaFileDescription) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"LockFD",
		"inode",
	}
}

func (rfd *replicaFileDescription) beforeSave() {}

// +checklocksignore
func (rfd *replicaFileDescription) StateSave(stateSinkObject state.Sink) {
	rfd.beforeSave()
	stateSinkObject.Save(0, &rfd.vfsfd)
	stateSinkObject.Save(1, &rfd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &rfd.LockFD)
	stateSinkObject.Save(3, &rfd.inode)
}

func (rfd *replicaFileDescription) afterLoad() {}

// +checklocksignore
func (rfd *replicaFileDescription) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &rfd.vfsfd)
	stateSourceObject.Load(1, &rfd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &rfd.LockFD)
	stateSourceObject.Load(3, &rfd.inode)
}

func (r *rootInodeRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/devpts.rootInodeRefs"
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

func (tm *Terminal) StateTypeName() string {
	return "pkg/sentry/fsimpl/devpts.Terminal"
}

func (tm *Terminal) StateFields() []string {
	return []string{
		"n",
		"ld",
		"masterKTTY",
		"replicaKTTY",
	}
}

func (tm *Terminal) beforeSave() {}

// +checklocksignore
func (tm *Terminal) StateSave(stateSinkObject state.Sink) {
	tm.beforeSave()
	stateSinkObject.Save(0, &tm.n)
	stateSinkObject.Save(1, &tm.ld)
	stateSinkObject.Save(2, &tm.masterKTTY)
	stateSinkObject.Save(3, &tm.replicaKTTY)
}

func (tm *Terminal) afterLoad() {}

// +checklocksignore
func (tm *Terminal) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &tm.n)
	stateSourceObject.Load(1, &tm.ld)
	stateSourceObject.Load(2, &tm.masterKTTY)
	stateSourceObject.Load(3, &tm.replicaKTTY)
}

func init() {
	state.Register((*FilesystemType)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*rootInode)(nil))
	state.Register((*implStatFS)(nil))
	state.Register((*lineDiscipline)(nil))
	state.Register((*outputQueueTransformer)(nil))
	state.Register((*inputQueueTransformer)(nil))
	state.Register((*masterInode)(nil))
	state.Register((*masterFileDescription)(nil))
	state.Register((*queue)(nil))
	state.Register((*replicaInode)(nil))
	state.Register((*replicaFileDescription)(nil))
	state.Register((*rootInodeRefs)(nil))
	state.Register((*Terminal)(nil))
}
