package timer

import (
	"fmt"
	"time"
)

type Timer struct {
	Name  string
	start time.Time
	stop  time.Time
}

func New(name string) *Timer {
	return &Timer{name, time.Now(), time.Now()}
}

func (t *Timer) Start() *Timer {
	t.start = time.Now()
	return t
}

func (t *Timer) Stop() *Timer {
	t.stop = time.Now()
	return t
}

func (t *Timer) String() string {
	prefix := "Timer:\t"
	dur := t.Elapsed(t.stop)
	if dur <= 0 {
		return fmt.Sprintf("%v %v not initialized", prefix, t.Name)
	}
	return fmt.Sprintf("%v %v took %v", prefix, t.Name, t.Elapsed(t.stop))
}

func (t Timer) Elapsed(stop time.Time) time.Duration {
	return stop.Sub(t.start)
}
