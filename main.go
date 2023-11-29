package main

import (
	"fmt"
	"root/trigger"
	"time"
)

func main() {
	trigger := trigger.New()

	counter := NewCounterTraffic(10)

	trigger.Send()
	trigger.Do(counter.Trigger, counter.Expire, time.Second)
}

type CounterTraffic struct {
	index       int
	counter     int
	diary       []int
	length      int
	middleValue float32
}

func NewCounterTraffic(length int) *CounterTraffic {
	return &CounterTraffic{
		diary:  make([]int, length),
		length: length,
	}
}

func (c *CounterTraffic) Add() {
	c.diary[c.index] = c.counter //set new element in our history of traffic
	fmt.Println(c.diary)
	c.Index()
}

func (c *CounterTraffic) Index() {
	c.index++

	if c.index == c.length {
		c.index = 0
	}
}

// fix performance for this function
func (c *CounterTraffic) MiddleValue() float32 {
	middle := 0
	for _, value := range c.diary {
		middle += value
	}

	return float32(middle) / float32(c.length)
}

func (c *CounterTraffic) Counter() {
	c.counter++
}

func (c *CounterTraffic) ResetCounter() {
	c.counter = 0
}

func (c *CounterTraffic) Trigger(int) {
	c.Counter()
}

func (c *CounterTraffic) Expire() {
	c.Add()
	c.ResetCounter()

	fmt.Println("average traffic", c.MiddleValue())
}
