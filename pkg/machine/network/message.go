package network

import "fmt"

type Packet struct {
	Src     int
	Dst     int
	Payload interface{}
}

func NewIntPairPacket(src, dst, x, y int) *Packet {
	payload := [2]int{x, y}
	return &Packet{
		src,
		dst,
		payload,
	}
}

func (m *Packet) String() string {
	return fmt.Sprintf("%v -> %v: %v", m.Src, m.Dst, m.Payload)
}
