package unix

import (
	"fmt"

	"github.com/nicocha30/gvisor-ligolo/pkg/atomicbitops"
	"github.com/nicocha30/gvisor-ligolo/pkg/refs"
)

// enableLogging indicates whether reference-related events should be logged (with
// stack traces). This is false by default and should only be set to true for
// debugging purposes, as it can generate an extremely large amount of output
// and drastically degrade performance.
const socketenableLogging = false

// obj is used to customize logging. Note that we use a pointer to T so that
// we do not copy the entire object when passed as a format parameter.
var socketobj *Socket

// Refs implements refs.RefCounter. It keeps a reference count using atomic
// operations and calls the destructor when the count reaches zero.
//
// NOTE: Do not introduce additional fields to the Refs struct. It is used by
// many filesystem objects, and we want to keep it as small as possible (i.e.,
// the same size as using an int64 directly) to avoid taking up extra cache
// space. In general, this template should not be extended at the cost of
// performance. If it does not offer enough flexibility for a particular object
// (example: b/187877947), we should implement the RefCounter/CheckedObject
// interfaces manually.
//
// +stateify savable
type socketRefs struct {
	// refCount is composed of two fields:
	//
	//	[32-bit speculative references]:[32-bit real references]
	//
	// Speculative references are used for TryIncRef, to avoid a CompareAndSwap
	// loop. See IncRef, DecRef and TryIncRef for details of how these fields are
	// used.
	refCount atomicbitops.Int64
}

// InitRefs initializes r with one reference and, if enabled, activates leak
// checking.
func (r *socketRefs) InitRefs() {
	r.refCount.Store(1)
	refs.Register(r)
}

// RefType implements refs.CheckedObject.RefType.
func (r *socketRefs) RefType() string {
	return fmt.Sprintf("%T", socketobj)[1:]
}

// LeakMessage implements refs.CheckedObject.LeakMessage.
func (r *socketRefs) LeakMessage() string {
	return fmt.Sprintf("[%s %p] reference count of %d instead of 0", r.RefType(), r, r.ReadRefs())
}

// LogRefs implements refs.CheckedObject.LogRefs.
func (r *socketRefs) LogRefs() bool {
	return socketenableLogging
}

// ReadRefs returns the current number of references. The returned count is
// inherently racy and is unsafe to use without external synchronization.
func (r *socketRefs) ReadRefs() int64 {
	return r.refCount.Load()
}

// IncRef implements refs.RefCounter.IncRef.
//
//go:nosplit
func (r *socketRefs) IncRef() {
	v := r.refCount.Add(1)
	if socketenableLogging {
		refs.LogIncRef(r, v)
	}
	if v <= 1 {
		panic(fmt.Sprintf("Incrementing non-positive count %p on %s", r, r.RefType()))
	}
}

// TryIncRef implements refs.TryRefCounter.TryIncRef.
//
// To do this safely without a loop, a speculative reference is first acquired
// on the object. This allows multiple concurrent TryIncRef calls to distinguish
// other TryIncRef calls from genuine references held.
//
//go:nosplit
func (r *socketRefs) TryIncRef() bool {
	const speculativeRef = 1 << 32
	if v := r.refCount.Add(speculativeRef); int32(v) == 0 {

		r.refCount.Add(-speculativeRef)
		return false
	}

	v := r.refCount.Add(-speculativeRef + 1)
	if socketenableLogging {
		refs.LogTryIncRef(r, v)
	}
	return true
}

// DecRef implements refs.RefCounter.DecRef.
//
// Note that speculative references are counted here. Since they were added
// prior to real references reaching zero, they will successfully convert to
// real references. In other words, we see speculative references only in the
// following case:
//
//	A: TryIncRef [speculative increase => sees non-negative references]
//	B: DecRef [real decrease]
//	A: TryIncRef [transform speculative to real]
//
//go:nosplit
func (r *socketRefs) DecRef(destroy func()) {
	v := r.refCount.Add(-1)
	if socketenableLogging {
		refs.LogDecRef(r, v)
	}
	switch {
	case v < 0:
		panic(fmt.Sprintf("Decrementing non-positive ref count %p, owned by %s", r, r.RefType()))

	case v == 0:
		refs.Unregister(r)

		if destroy != nil {
			destroy()
		}
	}
}

func (r *socketRefs) afterLoad() {
	if r.ReadRefs() > 0 {
		refs.Register(r)
	}
}
