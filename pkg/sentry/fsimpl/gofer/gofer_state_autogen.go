// automatically generated by stateify.

package gofer

import (
	"github.com/nicocha30/gvisor-ligolo/pkg/state"
)

func (l *dentryList) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.dentryList"
}

func (l *dentryList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *dentryList) beforeSave() {}

// +checklocksignore
func (l *dentryList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *dentryList) afterLoad() {}

// +checklocksignore
func (l *dentryList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *dentryEntry) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.dentryEntry"
}

func (e *dentryEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *dentryEntry) beforeSave() {}

// +checklocksignore
func (e *dentryEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *dentryEntry) afterLoad() {}

// +checklocksignore
func (e *dentryEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (fd *directoryFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.directoryFD"
}

func (fd *directoryFD) StateFields() []string {
	return []string{
		"fileDescription",
		"DirectoryFileDescriptionDefaultImpl",
		"off",
		"dirents",
	}
}

func (fd *directoryFD) beforeSave() {}

// +checklocksignore
func (fd *directoryFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.fileDescription)
	stateSinkObject.Save(1, &fd.DirectoryFileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.off)
	stateSinkObject.Save(3, &fd.dirents)
}

func (fd *directoryFD) afterLoad() {}

// +checklocksignore
func (fd *directoryFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.fileDescription)
	stateSourceObject.Load(1, &fd.DirectoryFileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.off)
	stateSourceObject.Load(3, &fd.dirents)
}

func (cache *stringFixedCache) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.stringFixedCache"
}

func (cache *stringFixedCache) StateFields() []string {
	return []string{
		"namesList",
		"size",
	}
}

func (cache *stringFixedCache) beforeSave() {}

// +checklocksignore
func (cache *stringFixedCache) StateSave(stateSinkObject state.Sink) {
	cache.beforeSave()
	stateSinkObject.Save(0, &cache.namesList)
	stateSinkObject.Save(1, &cache.size)
}

func (cache *stringFixedCache) afterLoad() {}

// +checklocksignore
func (cache *stringFixedCache) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &cache.namesList)
	stateSourceObject.Load(1, &cache.size)
}

func (d *dentryCache) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.dentryCache"
}

func (d *dentryCache) StateFields() []string {
	return []string{
		"dentries",
		"dentriesLen",
		"maxCachedDentries",
	}
}

func (d *dentryCache) beforeSave() {}

// +checklocksignore
func (d *dentryCache) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.dentries)
	stateSinkObject.Save(1, &d.dentriesLen)
	stateSinkObject.Save(2, &d.maxCachedDentries)
}

func (d *dentryCache) afterLoad() {}

// +checklocksignore
func (d *dentryCache) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.dentries)
	stateSourceObject.Load(1, &d.dentriesLen)
	stateSourceObject.Load(2, &d.maxCachedDentries)
}

func (fstype *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.FilesystemType"
}

func (fstype *FilesystemType) StateFields() []string {
	return []string{}
}

func (fstype *FilesystemType) beforeSave() {}

// +checklocksignore
func (fstype *FilesystemType) StateSave(stateSinkObject state.Sink) {
	fstype.beforeSave()
}

func (fstype *FilesystemType) afterLoad() {}

// +checklocksignore
func (fstype *FilesystemType) StateLoad(stateSourceObject state.Source) {
}

func (fs *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.filesystem"
}

func (fs *filesystem) StateFields() []string {
	return []string{
		"vfsfs",
		"mfp",
		"opts",
		"iopts",
		"clock",
		"devMinor",
		"root",
		"dentryCache",
		"syncableDentries",
		"specialFileFDs",
		"lastIno",
		"savedDentryRW",
		"released",
	}
}

func (fs *filesystem) beforeSave() {}

// +checklocksignore
func (fs *filesystem) StateSave(stateSinkObject state.Sink) {
	fs.beforeSave()
	stateSinkObject.Save(0, &fs.vfsfs)
	stateSinkObject.Save(1, &fs.mfp)
	stateSinkObject.Save(2, &fs.opts)
	stateSinkObject.Save(3, &fs.iopts)
	stateSinkObject.Save(4, &fs.clock)
	stateSinkObject.Save(5, &fs.devMinor)
	stateSinkObject.Save(6, &fs.root)
	stateSinkObject.Save(7, &fs.dentryCache)
	stateSinkObject.Save(8, &fs.syncableDentries)
	stateSinkObject.Save(9, &fs.specialFileFDs)
	stateSinkObject.Save(10, &fs.lastIno)
	stateSinkObject.Save(11, &fs.savedDentryRW)
	stateSinkObject.Save(12, &fs.released)
}

func (fs *filesystem) afterLoad() {}

// +checklocksignore
func (fs *filesystem) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fs.vfsfs)
	stateSourceObject.Load(1, &fs.mfp)
	stateSourceObject.Load(2, &fs.opts)
	stateSourceObject.Load(3, &fs.iopts)
	stateSourceObject.Load(4, &fs.clock)
	stateSourceObject.Load(5, &fs.devMinor)
	stateSourceObject.Load(6, &fs.root)
	stateSourceObject.Load(7, &fs.dentryCache)
	stateSourceObject.Load(8, &fs.syncableDentries)
	stateSourceObject.Load(9, &fs.specialFileFDs)
	stateSourceObject.Load(10, &fs.lastIno)
	stateSourceObject.Load(11, &fs.savedDentryRW)
	stateSourceObject.Load(12, &fs.released)
}

func (f *filesystemOptions) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.filesystemOptions"
}

func (f *filesystemOptions) StateFields() []string {
	return []string{
		"fd",
		"aname",
		"interop",
		"dfltuid",
		"dfltgid",
		"forcePageCache",
		"limitHostFDTranslation",
		"overlayfsStaleRead",
		"regularFilesUseSpecialFileFD",
	}
}

func (f *filesystemOptions) beforeSave() {}

// +checklocksignore
func (f *filesystemOptions) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.fd)
	stateSinkObject.Save(1, &f.aname)
	stateSinkObject.Save(2, &f.interop)
	stateSinkObject.Save(3, &f.dfltuid)
	stateSinkObject.Save(4, &f.dfltgid)
	stateSinkObject.Save(5, &f.forcePageCache)
	stateSinkObject.Save(6, &f.limitHostFDTranslation)
	stateSinkObject.Save(7, &f.overlayfsStaleRead)
	stateSinkObject.Save(8, &f.regularFilesUseSpecialFileFD)
}

func (f *filesystemOptions) afterLoad() {}

// +checklocksignore
func (f *filesystemOptions) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.fd)
	stateSourceObject.Load(1, &f.aname)
	stateSourceObject.Load(2, &f.interop)
	stateSourceObject.Load(3, &f.dfltuid)
	stateSourceObject.Load(4, &f.dfltgid)
	stateSourceObject.Load(5, &f.forcePageCache)
	stateSourceObject.Load(6, &f.limitHostFDTranslation)
	stateSourceObject.Load(7, &f.overlayfsStaleRead)
	stateSourceObject.Load(8, &f.regularFilesUseSpecialFileFD)
}

func (i *InteropMode) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.InteropMode"
}

func (i *InteropMode) StateFields() []string {
	return nil
}

func (i *InternalFilesystemOptions) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.InternalFilesystemOptions"
}

func (i *InternalFilesystemOptions) StateFields() []string {
	return []string{
		"UniqueID",
		"LeakConnection",
		"OpenSocketsByConnecting",
	}
}

func (i *InternalFilesystemOptions) beforeSave() {}

// +checklocksignore
func (i *InternalFilesystemOptions) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.UniqueID)
	stateSinkObject.Save(1, &i.LeakConnection)
	stateSinkObject.Save(2, &i.OpenSocketsByConnecting)
}

func (i *InternalFilesystemOptions) afterLoad() {}

// +checklocksignore
func (i *InternalFilesystemOptions) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.UniqueID)
	stateSourceObject.Load(1, &i.LeakConnection)
	stateSourceObject.Load(2, &i.OpenSocketsByConnecting)
}

func (i *inoKey) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.inoKey"
}

func (i *inoKey) StateFields() []string {
	return []string{
		"ino",
		"devMinor",
		"devMajor",
	}
}

func (i *inoKey) beforeSave() {}

// +checklocksignore
func (i *inoKey) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.ino)
	stateSinkObject.Save(1, &i.devMinor)
	stateSinkObject.Save(2, &i.devMajor)
}

func (i *inoKey) afterLoad() {}

// +checklocksignore
func (i *inoKey) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.ino)
	stateSourceObject.Load(1, &i.devMinor)
	stateSourceObject.Load(2, &i.devMajor)
}

func (d *dentry) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.dentry"
}

func (d *dentry) StateFields() []string {
	return []string{
		"vfsd",
		"refs",
		"fs",
		"parent",
		"name",
		"inoKey",
		"deleted",
		"cached",
		"cacheEntry",
		"syncableListEntry",
		"children",
		"negativeChildrenCache",
		"negativeChildren",
		"syntheticChildren",
		"dirents",
		"childrenSet",
		"ino",
		"mode",
		"uid",
		"gid",
		"blockSize",
		"atime",
		"mtime",
		"ctime",
		"btime",
		"size",
		"atimeDirty",
		"mtimeDirty",
		"nlink",
		"mappings",
		"cache",
		"dirty",
		"pf",
		"haveTarget",
		"target",
		"endpoint",
		"pipe",
		"locks",
		"watches",
		"impl",
	}
}

// +checklocksignore
func (d *dentry) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.vfsd)
	stateSinkObject.Save(1, &d.refs)
	stateSinkObject.Save(2, &d.fs)
	stateSinkObject.Save(3, &d.parent)
	stateSinkObject.Save(4, &d.name)
	stateSinkObject.Save(5, &d.inoKey)
	stateSinkObject.Save(6, &d.deleted)
	stateSinkObject.Save(7, &d.cached)
	stateSinkObject.Save(8, &d.cacheEntry)
	stateSinkObject.Save(9, &d.syncableListEntry)
	stateSinkObject.Save(10, &d.children)
	stateSinkObject.Save(11, &d.negativeChildrenCache)
	stateSinkObject.Save(12, &d.negativeChildren)
	stateSinkObject.Save(13, &d.syntheticChildren)
	stateSinkObject.Save(14, &d.dirents)
	stateSinkObject.Save(15, &d.childrenSet)
	stateSinkObject.Save(16, &d.ino)
	stateSinkObject.Save(17, &d.mode)
	stateSinkObject.Save(18, &d.uid)
	stateSinkObject.Save(19, &d.gid)
	stateSinkObject.Save(20, &d.blockSize)
	stateSinkObject.Save(21, &d.atime)
	stateSinkObject.Save(22, &d.mtime)
	stateSinkObject.Save(23, &d.ctime)
	stateSinkObject.Save(24, &d.btime)
	stateSinkObject.Save(25, &d.size)
	stateSinkObject.Save(26, &d.atimeDirty)
	stateSinkObject.Save(27, &d.mtimeDirty)
	stateSinkObject.Save(28, &d.nlink)
	stateSinkObject.Save(29, &d.mappings)
	stateSinkObject.Save(30, &d.cache)
	stateSinkObject.Save(31, &d.dirty)
	stateSinkObject.Save(32, &d.pf)
	stateSinkObject.Save(33, &d.haveTarget)
	stateSinkObject.Save(34, &d.target)
	stateSinkObject.Save(35, &d.endpoint)
	stateSinkObject.Save(36, &d.pipe)
	stateSinkObject.Save(37, &d.locks)
	stateSinkObject.Save(38, &d.watches)
	stateSinkObject.Save(39, &d.impl)
}

// +checklocksignore
func (d *dentry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.vfsd)
	stateSourceObject.Load(1, &d.refs)
	stateSourceObject.Load(2, &d.fs)
	stateSourceObject.Load(3, &d.parent)
	stateSourceObject.Load(4, &d.name)
	stateSourceObject.Load(5, &d.inoKey)
	stateSourceObject.Load(6, &d.deleted)
	stateSourceObject.Load(7, &d.cached)
	stateSourceObject.Load(8, &d.cacheEntry)
	stateSourceObject.Load(9, &d.syncableListEntry)
	stateSourceObject.Load(10, &d.children)
	stateSourceObject.Load(11, &d.negativeChildrenCache)
	stateSourceObject.Load(12, &d.negativeChildren)
	stateSourceObject.Load(13, &d.syntheticChildren)
	stateSourceObject.Load(14, &d.dirents)
	stateSourceObject.Load(15, &d.childrenSet)
	stateSourceObject.Load(16, &d.ino)
	stateSourceObject.Load(17, &d.mode)
	stateSourceObject.Load(18, &d.uid)
	stateSourceObject.Load(19, &d.gid)
	stateSourceObject.Load(20, &d.blockSize)
	stateSourceObject.Load(21, &d.atime)
	stateSourceObject.Load(22, &d.mtime)
	stateSourceObject.Load(23, &d.ctime)
	stateSourceObject.Load(24, &d.btime)
	stateSourceObject.Load(25, &d.size)
	stateSourceObject.Load(26, &d.atimeDirty)
	stateSourceObject.Load(27, &d.mtimeDirty)
	stateSourceObject.Load(28, &d.nlink)
	stateSourceObject.Load(29, &d.mappings)
	stateSourceObject.Load(30, &d.cache)
	stateSourceObject.Load(31, &d.dirty)
	stateSourceObject.Load(32, &d.pf)
	stateSourceObject.Load(33, &d.haveTarget)
	stateSourceObject.Load(34, &d.target)
	stateSourceObject.Load(35, &d.endpoint)
	stateSourceObject.Load(36, &d.pipe)
	stateSourceObject.Load(37, &d.locks)
	stateSourceObject.Load(38, &d.watches)
	stateSourceObject.Load(39, &d.impl)
	stateSourceObject.AfterLoad(d.afterLoad)
}

func (s *stringListElem) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.stringListElem"
}

func (s *stringListElem) StateFields() []string {
	return []string{
		"str",
		"stringEntry",
	}
}

func (s *stringListElem) beforeSave() {}

// +checklocksignore
func (s *stringListElem) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.str)
	stateSinkObject.Save(1, &s.stringEntry)
}

func (s *stringListElem) afterLoad() {}

// +checklocksignore
func (s *stringListElem) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.str)
	stateSourceObject.Load(1, &s.stringEntry)
}

func (d *dentryListElem) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.dentryListElem"
}

func (d *dentryListElem) StateFields() []string {
	return []string{
		"d",
		"dentryEntry",
	}
}

func (d *dentryListElem) beforeSave() {}

// +checklocksignore
func (d *dentryListElem) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.d)
	stateSinkObject.Save(1, &d.dentryEntry)
}

func (d *dentryListElem) afterLoad() {}

// +checklocksignore
func (d *dentryListElem) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.d)
	stateSourceObject.Load(1, &d.dentryEntry)
}

func (fd *fileDescription) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.fileDescription"
}

func (fd *fileDescription) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"LockFD",
	}
}

func (fd *fileDescription) beforeSave() {}

// +checklocksignore
func (fd *fileDescription) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.vfsfd)
	stateSinkObject.Save(1, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.LockFD)
}

func (fd *fileDescription) afterLoad() {}

// +checklocksignore
func (fd *fileDescription) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.vfsfd)
	stateSourceObject.Load(1, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.LockFD)
}

func (d *lisafsDentry) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.lisafsDentry"
}

func (d *lisafsDentry) StateFields() []string {
	return []string{
		"dentry",
	}
}

func (d *lisafsDentry) beforeSave() {}

// +checklocksignore
func (d *lisafsDentry) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.dentry)
}

func (d *lisafsDentry) afterLoad() {}

// +checklocksignore
func (d *lisafsDentry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.dentry)
}

func (fd *regularFileFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.regularFileFD"
}

func (fd *regularFileFD) StateFields() []string {
	return []string{
		"fileDescription",
		"off",
	}
}

func (fd *regularFileFD) beforeSave() {}

// +checklocksignore
func (fd *regularFileFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.fileDescription)
	stateSinkObject.Save(1, &fd.off)
}

func (fd *regularFileFD) afterLoad() {}

// +checklocksignore
func (fd *regularFileFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.fileDescription)
	stateSourceObject.Load(1, &fd.off)
}

func (d *dentryPlatformFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.dentryPlatformFile"
}

func (d *dentryPlatformFile) StateFields() []string {
	return []string{
		"dentry",
		"fdRefs",
		"hostFileMapper",
	}
}

func (d *dentryPlatformFile) beforeSave() {}

// +checklocksignore
func (d *dentryPlatformFile) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.dentry)
	stateSinkObject.Save(1, &d.fdRefs)
	stateSinkObject.Save(2, &d.hostFileMapper)
}

// +checklocksignore
func (d *dentryPlatformFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.dentry)
	stateSourceObject.Load(1, &d.fdRefs)
	stateSourceObject.Load(2, &d.hostFileMapper)
	stateSourceObject.AfterLoad(d.afterLoad)
}

func (s *savedDentryRW) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.savedDentryRW"
}

func (s *savedDentryRW) StateFields() []string {
	return []string{
		"read",
		"write",
	}
}

func (s *savedDentryRW) beforeSave() {}

// +checklocksignore
func (s *savedDentryRW) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.read)
	stateSinkObject.Save(1, &s.write)
}

func (s *savedDentryRW) afterLoad() {}

// +checklocksignore
func (s *savedDentryRW) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.read)
	stateSourceObject.Load(1, &s.write)
}

func (e *endpoint) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.endpoint"
}

func (e *endpoint) StateFields() []string {
	return []string{
		"dentry",
		"path",
	}
}

func (e *endpoint) beforeSave() {}

// +checklocksignore
func (e *endpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.dentry)
	stateSinkObject.Save(1, &e.path)
}

func (e *endpoint) afterLoad() {}

// +checklocksignore
func (e *endpoint) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.dentry)
	stateSourceObject.Load(1, &e.path)
}

func (l *specialFDList) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.specialFDList"
}

func (l *specialFDList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *specialFDList) beforeSave() {}

// +checklocksignore
func (l *specialFDList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *specialFDList) afterLoad() {}

// +checklocksignore
func (l *specialFDList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *specialFDEntry) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.specialFDEntry"
}

func (e *specialFDEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *specialFDEntry) beforeSave() {}

// +checklocksignore
func (e *specialFDEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *specialFDEntry) afterLoad() {}

// +checklocksignore
func (e *specialFDEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (fd *specialFileFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.specialFileFD"
}

func (fd *specialFileFD) StateFields() []string {
	return []string{
		"fileDescription",
		"specialFDEntry",
		"isRegularFile",
		"seekable",
		"queue",
		"off",
		"haveBuf",
		"buf",
		"hostFileMapper",
		"fileRefs",
	}
}

func (fd *specialFileFD) beforeSave() {}

// +checklocksignore
func (fd *specialFileFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.fileDescription)
	stateSinkObject.Save(1, &fd.specialFDEntry)
	stateSinkObject.Save(2, &fd.isRegularFile)
	stateSinkObject.Save(3, &fd.seekable)
	stateSinkObject.Save(4, &fd.queue)
	stateSinkObject.Save(5, &fd.off)
	stateSinkObject.Save(6, &fd.haveBuf)
	stateSinkObject.Save(7, &fd.buf)
	stateSinkObject.Save(8, &fd.hostFileMapper)
	stateSinkObject.Save(9, &fd.fileRefs)
}

// +checklocksignore
func (fd *specialFileFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.fileDescription)
	stateSourceObject.Load(1, &fd.specialFDEntry)
	stateSourceObject.Load(2, &fd.isRegularFile)
	stateSourceObject.Load(3, &fd.seekable)
	stateSourceObject.Load(4, &fd.queue)
	stateSourceObject.Load(5, &fd.off)
	stateSourceObject.Load(6, &fd.haveBuf)
	stateSourceObject.Load(7, &fd.buf)
	stateSourceObject.Load(8, &fd.hostFileMapper)
	stateSourceObject.Load(9, &fd.fileRefs)
	stateSourceObject.AfterLoad(fd.afterLoad)
}

func (l *stringList) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.stringList"
}

func (l *stringList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *stringList) beforeSave() {}

// +checklocksignore
func (l *stringList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *stringList) afterLoad() {}

// +checklocksignore
func (l *stringList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *stringEntry) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.stringEntry"
}

func (e *stringEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *stringEntry) beforeSave() {}

// +checklocksignore
func (e *stringEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *stringEntry) afterLoad() {}

// +checklocksignore
func (e *stringEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func init() {
	state.Register((*dentryList)(nil))
	state.Register((*dentryEntry)(nil))
	state.Register((*directoryFD)(nil))
	state.Register((*stringFixedCache)(nil))
	state.Register((*dentryCache)(nil))
	state.Register((*FilesystemType)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*filesystemOptions)(nil))
	state.Register((*InteropMode)(nil))
	state.Register((*InternalFilesystemOptions)(nil))
	state.Register((*inoKey)(nil))
	state.Register((*dentry)(nil))
	state.Register((*stringListElem)(nil))
	state.Register((*dentryListElem)(nil))
	state.Register((*fileDescription)(nil))
	state.Register((*lisafsDentry)(nil))
	state.Register((*regularFileFD)(nil))
	state.Register((*dentryPlatformFile)(nil))
	state.Register((*savedDentryRW)(nil))
	state.Register((*endpoint)(nil))
	state.Register((*specialFDList)(nil))
	state.Register((*specialFDEntry)(nil))
	state.Register((*specialFileFD)(nil))
	state.Register((*stringList)(nil))
	state.Register((*stringEntry)(nil))
}
