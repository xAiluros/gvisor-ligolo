// automatically generated by stateify.

package icmp

import (
	"github.com/nicocha30/gvisor-ligolo/pkg/state"
)

func (p *icmpPacket) StateTypeName() string {
	return "pkg/tcpip/transport/icmp.icmpPacket"
}

func (p *icmpPacket) StateFields() []string {
	return []string{
		"icmpPacketEntry",
		"senderAddress",
		"packetInfo",
		"data",
		"receivedAt",
		"tosOrTClass",
		"ttlOrHopLimit",
	}
}

func (p *icmpPacket) beforeSave() {}

// +checklocksignore
func (p *icmpPacket) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	var receivedAtValue int64
	receivedAtValue = p.saveReceivedAt()
	stateSinkObject.SaveValue(4, receivedAtValue)
	stateSinkObject.Save(0, &p.icmpPacketEntry)
	stateSinkObject.Save(1, &p.senderAddress)
	stateSinkObject.Save(2, &p.packetInfo)
	stateSinkObject.Save(3, &p.data)
	stateSinkObject.Save(5, &p.tosOrTClass)
	stateSinkObject.Save(6, &p.ttlOrHopLimit)
}

func (p *icmpPacket) afterLoad() {}

// +checklocksignore
func (p *icmpPacket) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.icmpPacketEntry)
	stateSourceObject.Load(1, &p.senderAddress)
	stateSourceObject.Load(2, &p.packetInfo)
	stateSourceObject.Load(3, &p.data)
	stateSourceObject.Load(5, &p.tosOrTClass)
	stateSourceObject.Load(6, &p.ttlOrHopLimit)
	stateSourceObject.LoadValue(4, new(int64), func(y any) { p.loadReceivedAt(y.(int64)) })
}

func (e *endpoint) StateTypeName() string {
	return "pkg/tcpip/transport/icmp.endpoint"
}

func (e *endpoint) StateFields() []string {
	return []string{
		"DefaultSocketOptionsHandler",
		"transProto",
		"waiterQueue",
		"uniqueID",
		"net",
		"stats",
		"ops",
		"rcvReady",
		"rcvList",
		"rcvBufSize",
		"rcvClosed",
		"frozen",
		"ident",
	}
}

// +checklocksignore
func (e *endpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.DefaultSocketOptionsHandler)
	stateSinkObject.Save(1, &e.transProto)
	stateSinkObject.Save(2, &e.waiterQueue)
	stateSinkObject.Save(3, &e.uniqueID)
	stateSinkObject.Save(4, &e.net)
	stateSinkObject.Save(5, &e.stats)
	stateSinkObject.Save(6, &e.ops)
	stateSinkObject.Save(7, &e.rcvReady)
	stateSinkObject.Save(8, &e.rcvList)
	stateSinkObject.Save(9, &e.rcvBufSize)
	stateSinkObject.Save(10, &e.rcvClosed)
	stateSinkObject.Save(11, &e.frozen)
	stateSinkObject.Save(12, &e.ident)
}

// +checklocksignore
func (e *endpoint) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.DefaultSocketOptionsHandler)
	stateSourceObject.Load(1, &e.transProto)
	stateSourceObject.Load(2, &e.waiterQueue)
	stateSourceObject.Load(3, &e.uniqueID)
	stateSourceObject.Load(4, &e.net)
	stateSourceObject.Load(5, &e.stats)
	stateSourceObject.Load(6, &e.ops)
	stateSourceObject.Load(7, &e.rcvReady)
	stateSourceObject.Load(8, &e.rcvList)
	stateSourceObject.Load(9, &e.rcvBufSize)
	stateSourceObject.Load(10, &e.rcvClosed)
	stateSourceObject.Load(11, &e.frozen)
	stateSourceObject.Load(12, &e.ident)
	stateSourceObject.AfterLoad(e.afterLoad)
}

func (l *icmpPacketList) StateTypeName() string {
	return "pkg/tcpip/transport/icmp.icmpPacketList"
}

func (l *icmpPacketList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *icmpPacketList) beforeSave() {}

// +checklocksignore
func (l *icmpPacketList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *icmpPacketList) afterLoad() {}

// +checklocksignore
func (l *icmpPacketList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *icmpPacketEntry) StateTypeName() string {
	return "pkg/tcpip/transport/icmp.icmpPacketEntry"
}

func (e *icmpPacketEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *icmpPacketEntry) beforeSave() {}

// +checklocksignore
func (e *icmpPacketEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *icmpPacketEntry) afterLoad() {}

// +checklocksignore
func (e *icmpPacketEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func init() {
	state.Register((*icmpPacket)(nil))
	state.Register((*endpoint)(nil))
	state.Register((*icmpPacketList)(nil))
	state.Register((*icmpPacketEntry)(nil))
}
