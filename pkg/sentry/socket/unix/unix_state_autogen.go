// automatically generated by stateify.

package unix

import (
	"github.com/nicocha30/gvisor-ligolo/pkg/state"
)

func (r *socketRefs) StateTypeName() string {
	return "pkg/sentry/socket/unix.socketRefs"
}

func (r *socketRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *socketRefs) beforeSave() {}

// +checklocksignore
func (r *socketRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *socketRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func (s *Socket) StateTypeName() string {
	return "pkg/sentry/socket/unix.Socket"
}

func (s *Socket) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"LockFD",
		"SendReceiveTimeout",
		"socketRefs",
		"ep",
		"stype",
		"abstractName",
		"abstractNamespace",
	}
}

func (s *Socket) beforeSave() {}

// +checklocksignore
func (s *Socket) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.vfsfd)
	stateSinkObject.Save(1, &s.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &s.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &s.LockFD)
	stateSinkObject.Save(4, &s.SendReceiveTimeout)
	stateSinkObject.Save(5, &s.socketRefs)
	stateSinkObject.Save(6, &s.ep)
	stateSinkObject.Save(7, &s.stype)
	stateSinkObject.Save(8, &s.abstractName)
	stateSinkObject.Save(9, &s.abstractNamespace)
}

func (s *Socket) afterLoad() {}

// +checklocksignore
func (s *Socket) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.vfsfd)
	stateSourceObject.Load(1, &s.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &s.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &s.LockFD)
	stateSourceObject.Load(4, &s.SendReceiveTimeout)
	stateSourceObject.Load(5, &s.socketRefs)
	stateSourceObject.Load(6, &s.ep)
	stateSourceObject.Load(7, &s.stype)
	stateSourceObject.Load(8, &s.abstractName)
	stateSourceObject.Load(9, &s.abstractNamespace)
}

func init() {
	state.Register((*socketRefs)(nil))
	state.Register((*Socket)(nil))
}
