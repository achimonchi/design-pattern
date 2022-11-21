package singleton

type Counter struct {
	Count int
}

var counter *Counter

func GetCounter() *Counter {

	if counter == nil {
		counter = &Counter{
			Count: 0,
		}
	}
	return counter
}

func (c *Counter) Add() {
	c.Count++
}
