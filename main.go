package main

import (
	"fmt"
	"math/rand"
	"time"
)

//try make more reader for faster reading
//select rite structure

func main() {
	// trigger := make(chan int)
	// var endOfCycle <-chan time.Time

	// ram := make([]int, 0, 10_000_000) // make better

	// // client simulation request
	// // go func() {
	// // 	var counter int
	// // 	for {
	// // 		rand := rand.Intn(1000)
	// // 		// time.Sleep(10 * time.Microsecond)
	// // 		trigger <- rand //Random time req
	// // 		counter++

	// // 		if counter == 100_000 {
	// // 			close(trigger)
	// // 			break
	// // 		}
	// // 	}
	// // }()

	// go func() {
	// 	for i := 0; i < 100_000_000; i++ {
	// 		trigger <- rand.Intn(1000)
	// 	}
	// 	close(trigger)
	// }()

	//processing smart cache
	// end:
	// 	for {
	// 		select {
	// 		case data, ok := <-trigger:
	// 			if !ok {
	// 				// Simulacija klijenta završena, izađi iz petlje
	// 				break end
	// 			}
	// 			//Add to our storage
	// 			ram = append(ram, data)

	// 			//if timer is not activated, set timer on 1 second
	// 			if endOfCycle == nil {
	// 				endOfCycle = time.After(1 * time.Second)
	// 			}

	// 			//check if time for cache is expire
	// 		case <-endOfCycle:
	// 			//delete everything from memory
	// 			fmt.Println(len(ram))
	// 			ram = make([]int, 0, 10_000_000) //make more efficient implementation, without new allocation

	//			//reset timer
	//			endOfCycle = nil
	//		}
	//	}

	trigger := New()
	trigger.Send()
	trigger.Do(Counter, WriteCounter, time.Second)
}

var myCounter int

func ResetCounter(counter *int) {
	*counter = 0
}

func Counter(number int) {
	myCounter += number
}

func WriteCounter() {
	fmt.Println(myCounter)
	ResetCounter(&myCounter)
}

type Trigger struct {
	trigger    chan int
	endOfCycle <-chan time.Time
}

func New() *Trigger {
	return &Trigger{
		trigger: make(chan int),
	}
}

func (t *Trigger) Do(fn func(int), end func(), duration time.Duration) {
	for {
		select {
		case data := <-t.trigger:
			//call function every time when trigger is triggered
			fn(data)

			//if timer is not activated, set timer on 1 second
			if t.endOfCycle == nil {
				t.endOfCycle = time.After(duration)
			}

			//check if time for cache is expire
		case <-t.endOfCycle:
			//call function who do something what need to do when timer is expire
			end()
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
