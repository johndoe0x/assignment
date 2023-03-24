package main

import (
	"container/ring"
	"fmt"
)

type Point struct {
	ID int
}

func NewCircle(point int) *ring.Ring {
	circle := ring.New(point)
	for i := 1; i <= point; i++ {
		circle.Value = &Point{ID: i}
		circle = circle.Next()
	}
	return circle
}

func main() {
	var point, i1, i2, j1, j2 int

	fmt.Println("Enter the number of points in the circle:")
	fmt.Scan(&point)

	circle := NewCircle(point)

	fmt.Println("Enter the start and end points for the first interval (i1, j1):")
	fmt.Scan(&i1, &j1)

	fmt.Println("Enter the start and end points for the second interval (i2, j2):")
	fmt.Scan(&i2, &j2)

	union := findUnion(circle, point, i1, j1, i2, j2)

	if union == nil {
		fmt.Println("NO union existed")
	} else {
		fmt.Printf("The union of intervals is [%d, %d)\n", union.start, union.end)
	}
}

type Interval struct {
	start, end int
}

func findUnion(circle *ring.Ring, n, i1, j1, i2, j2 int) *Interval {
	if (i1 == j1 && i2 == j2) || (i1 == j2 && i2 == j1) {
		return &Interval{
			start: 1,
			end:   n,
		}
	}

	interval1, interval2 := findIntervalRings(circle, i1, i2)

	if interval1 == nil || interval2 == nil {
		return nil
	}

	unionStart, unionEnd := calculateUnion(interval1, j1-i1, interval2, j2-i2)

	if unionStart == nil || unionEnd == nil {
		return nil
	}

	return &Interval{
		start: unionStart.Value.(*Point).ID,
		end:   unionEnd.Value.(*Point).ID,
	}
}

func findIntervalRings(circle *ring.Ring, i1, i2 int) (interval1, interval2 *ring.Ring) {
	circle.Do(func(p interface{}) {
		point := p.(*Point)
		if point.ID == i1 {
			interval1 = circle
		} else if point.ID == i2 {
			interval2 = circle
		}
	})
	return
}

func calculateUnion(interval1 *ring.Ring, length1 int, interval2 *ring.Ring, length2 int) (unionStart, unionEnd *ring.Ring) {
	unionStart = interval1
	unionEnd = interval1.Move(length1)

	if interval2 == unionEnd {
		unionEnd = unionEnd.Move(length2)
	} else if interval2.Move(length2) == unionStart {
		unionStart = interval2
	} else {
		return nil, nil
	}

	return unionStart, unionEnd
}
