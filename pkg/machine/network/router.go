package network

import (
	"fmt"
)

type Router struct {
	nics          map[int]*Controller
	debugMode     bool
	LastBroadcast *Packet
	NetworkReset  *Packet
}

func NewRouter() *Router {
	return &Router{
		make(map[int]*Controller),
		false,
		nil,
		nil,
	}
}

func (r *Router) SetDebugMode(d bool) {
	r.debugMode = d
}

func (r *Router) AddNic(nic *Controller) {
	r.nics[nic.Address()] = nic
	nic.SetRouter(r)
}

func (r *Router) RemoveNic(i int) {
	delete(r.nics, i)
}

func (r *Router) Route(m *Packet) {
	if r.debugMode {
		fmt.Printf("Router pushing: %v\n", m)
	}

	if m.Dst == 255 {
		r.LastBroadcast = m
	}

	nic, ok := r.nics[m.Dst]
	if !ok {
		if r.debugMode {
			fmt.Printf("Router can't find NIC[%v] to deliver: %v\n", m.Dst, m)
		}
		return
	}
	nic.QueuePush(m)
}

func (r *Router) Run() {
	for i := 0; i < len(r.nics); i++ {
		nic := r.nics[i]
		if r.debugMode {
			fmt.Printf("Executing one step of %v\n", nic)
		}
		nic.ExecOneStep()
	}
	if r.IsIdle() {
		r.NetworkReset = r.LastBroadcast
		nic := r.nics[0]
		if r.debugMode {
			fmt.Printf("Router is Idle. Sending to NIC[%v]: %v\n", 0, r.NetworkReset)
		}
		nic.QueuePush(r.NetworkReset)
	}
}

func (r *Router) IsIdle() bool {
	for i := 0; i < len(r.nics); i++ {
		nic := r.nics[i]
		if !nic.Idle {
			return false
		}
	}
	return true
}
