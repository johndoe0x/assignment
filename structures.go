package main

import "container/ring"

func NewCircleByRingCollection(point int) *ring.Ring {
	circle := ring.New(point)
	for i := 1; i <= point; i++ {
		circle.Value = i
		circle = circle.Next()
	}
	return circle
}

type Interval struct {
	start, end int
}

type Circle struct {
	Center      [2]float64
	Radius      float64
	FixedPoints [][2]float64
}
