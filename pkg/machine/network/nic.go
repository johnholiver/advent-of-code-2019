package network

import (
	"container/list"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	computer_io "github.com/johnholiver/advent-of-code-2019/pkg/computer/io"
	computer_mem "github.com/johnholiver/advent-of-code-2019/pkg/computer/memory"
)

type Controller struct {
	P         *computer.Processor
	addr      int
	router    *Router
	queue     *list.List
	Idle      bool
	debugMode bool
}

func NewController(addr int, program string) *Controller {
	buildComputer := func(program string) *computer.Processor {
		i := computer_io.NewTape()
		i.Append(addr)
		i.Append(-1) //Start with empty queue
		p := computer.NewProcessor(i, nil, nil)
		m := computer_mem.NewRelative(p, program)
		p.Memory = m
		o := computer_io.NewInterruptingTape(p)
		p.Output = o
		return p
	}

	return &Controller{
		buildComputer(program),
		addr,
		nil,
		list.New(),
		false,
		false,
	}
}

func (c *Controller) String() string {
	return fmt.Sprintf("[%v]", c.addr)
}

func (c *Controller) Exec() {
	c.Idle = false
	for c.P.WaitingInput == false {
		c.ExecOneStep()
	}
	c.Idle = true
}

func (c *Controller) ExecOneStep() {
	c.Idle = false
	var stepInput *[2]int
	inputM := c.QueuePop()

	if inputM != nil {
		if c.debugMode {
			fmt.Printf("%v pop: %v\n", c, inputM)
		}
		payload := inputM.Payload.([2]int)
		stepInput = &payload
	}

	output, done := c.ProcessOne(stepInput)
	if output == nil || done {
		if c.P.WaitingInput {
			c.Idle = true
		}
		return
	}

	outputM := NewIntPairPacket(c.addr, output[0], output[1], output[2])
	if c.debugMode {
		fmt.Printf("%v queueing: %v\n", c, outputM)
	}
	if c.router == nil {
		if c.debugMode {
			fmt.Println("No router found")
		}
		return
	}
	c.router.Route(outputM)
}

func (c *Controller) ProcessOne(input *[2]int) ([]int, bool) {
	output := make([]int, 3)
	if input != nil {
		c.P.Input.Append(input[0])
		c.P.Input.Append(input[1])
		c.P.Input.Append(-1)
	}

	c.P.Process()
	if c.P.IsHalted {
		//Emergency break :D
		return output, true
	}

	if !c.P.Output.CanRead() {
		return nil, false
	}
	output[0] = c.P.Output.Read()
	c.P.Input.Append(-1)
	c.P.Process()
	output[1] = c.P.Output.Read()
	c.P.Input.Append(-1)
	c.P.Process()
	output[2] = c.P.Output.Read()
	return output, false
}

func (c *Controller) SetDebugMode(d bool) {
	c.debugMode = d
}

func (c *Controller) Address() int {
	return c.addr
}

func (c *Controller) SetRouter(r *Router) {
	c.router = r
}

func (c *Controller) QueuePush(m *Packet) {
	// Enqueue
	c.queue.PushBack(m)
}

func (c *Controller) QueuePop() *Packet {
	if c.queue.Len() == 0 {
		return nil
	}

	// Read first
	e := c.queue.Front()
	// Dequeue
	c.queue.Remove(e)
	return e.Value.(*Packet)
}
