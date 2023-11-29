package trigger

import (
	"math/rand"
	"time"
)

type Trigger struct {
	trigger    chan int
	endOfCycle <-chan time.Time
}

func New() *Trigger {
	return &Trigger{
		trigger: make(chan int),
	}
}

func (t *Trigger) Do(trigger func(int), expire func(), duration time.Duration) {
	for {
		select {
		case data := <-t.trigger:
			//call function every time when trigger is triggered
			trigger(data)

			//if timer is not activated, set timer on 1 second
			if t.endOfCycle == nil {
				t.endOfCycle = time.After(duration)
			}

			//check if time for cache is expire
		case <-t.endOfCycle:
			//call function who do something what need to do when timer is expire
			expire()
			//reset timer
			t.endOfCycle = nil
		}
	}
}

func (t *Trigger) Send() {
	go func() {
		for {
			num := rand.Intn(1000)
			time.Sleep(time.Duration(num) * time.Millisecond)

			t.trigger <- num
		}
	}()
}
