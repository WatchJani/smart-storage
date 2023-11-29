package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkSpeed(b *testing.B) {
	capacity := 100_000

	data := make([]int, 0, capacity)

	for k := 0; k < b.N; k++ {
		for i := 0; i < capacity; i++ {
			data = append(data, i)
		}
	}
}

func BenchmarkSpeedTest(b *testing.B) {
	capacity := 4_000_000

	data := make([]int, 0, capacity)
	sender := make(chan int)

	go func() {
		for i := 0; i < capacity; i++ {
			sender <- 5
		}

		close(sender)
	}()

	for k := 0; k < b.N; k++ {
		for number := range sender {
			data = append(data, number)
		}
	}
}

func BenchmarkStore(b *testing.B) {
	trigger := make(chan int)
	var endOfCycle <-chan time.Time

	ram := make([]int, 0, 500_000) // make better

	go func() {
		for i := 0; i < 500_000; i++ {
			trigger <- rand.Intn(1000)
		}
		close(trigger)
	}()

	for k := 0; k < b.N; k++ {
	end:
		for {
			select {
			case data, ok := <-trigger:
				if !ok {
					// Simulacija klijenta završena, izađi iz petlje
					break end
				}
				//Add to our storage
				ram = append(ram, data)

				//if timer is not activated, set timer on 1 second
				if endOfCycle == nil {
					endOfCycle = time.After(1 * time.Second)
				}

				//check if time for cache is expire
			case <-endOfCycle:
				//delete everything from memory
				// ram = make([]int, 0, 10_000_000) //make more efficient implementation, without new allocation

				//reset timer
				endOfCycle = nil
			}
		}
	}
}
