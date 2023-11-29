package trigger

import "time"

type TriggerV2[T any] struct {
	trigger    chan T
	endOfCycle <-chan time.Time
}

func NewTrigger[T any]() *TriggerV2[T] {
	return &TriggerV2[T]{
		trigger: make(chan T),
	}
}

func (t *TriggerV2[T]) Do(trigger func(T), expire func(), duration time.Duration) {
	for {
		select {
		case input := <-t.trigger:
			//call function every time when trigger is triggered
			trigger(input)

			//if timer is not activated, set timer on duration
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
