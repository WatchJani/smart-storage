package main

import (
	"math/rand"
	"time"
)

//try make more reader for faster reading
//select rite structure

func main() {
	trigger := make(chan int)
	var endOfCycle <-chan time.Time

	ram := make([]int, 25000) // make better

	//client simulation request
	go func() {
		var counter int
		for {
			rand := rand.Intn(10)
			time.Sleep(time.Duration(rand) * time.Microsecond)
			trigger <- rand //Random time req
			counter++

			if counter == 100000 {
				close(trigger)
				break
			}
		}
	}()

	//processing smart cache
	for {
		select {
		case data := <-trigger:

			//Add to our storage
			ram = append(ram, data)

			//if timer is not activated, set timer on 1 second
			if endOfCycle == nil {
				endOfCycle = time.After(1 * time.Second)
			}

			//check if time for cache is expire
		case <-endOfCycle:
			//delete everything from memory
			ram = make([]int, 25000) //make more efficient implementation, without new allocation

			//reset timer
			endOfCycle = nil
		}
	}
}
