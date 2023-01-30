// automatically generated by stateify.

//go:build amd64 && amd64 && amd64
// +build amd64,amd64,amd64

package arch

import (
	"github.com/nicocha30/gvisor-ligolo/pkg/state"
)

func (c *Context64) StateTypeName() string {
	return "pkg/sentry/arch.Context64"
}

func (c *Context64) StateFields() []string {
	return []string{
		"State",
	}
}

func (c *Context64) beforeSave() {}

// +checklocksignore
func (c *Context64) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.State)
}

func (c *Context64) afterLoad() {}

// +checklocksignore
func (c *Context64) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.State)
}

func init() {
	state.Register((*Context64)(nil))
}
