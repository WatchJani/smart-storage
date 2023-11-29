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
	c.MiddleValue(float32(c.counter))
	c.diary[c.index] = c.counter //set new element in our history of traffic
	c.Index()
}

func (c *CounterTraffic) Index() {
	c.index++

	if c.index == c.length {
		c.index = 0
	}
}

func (c CounterTraffic) GetMiddleValue() float32 {
	return c.middleValue
}

// max performance for average value
func (c *CounterTraffic) MiddleValue(value float32) {
	c.middleValue += (value - float32(c.diary[c.index])) / float32(c.length)
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

	fmt.Println("average traffic", c.GetMiddleValue())
}
