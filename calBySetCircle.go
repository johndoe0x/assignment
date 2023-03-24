package main

import (
	"fmt"
	"math"
)

func NewCircle(center [2]float64, radius float64, FixedPoints int) *Circle {
	circle := Circle{Center: center, Radius: radius}
	angleDivisor := FixedPoints
	for k := 0; k < FixedPoints; k++ {
		theta := float64(k) * 2.0 * math.Pi / float64(angleDivisor)
		x := radius*math.Cos(theta) + center[0]
		y := radius*math.Sin(theta) + center[1]
		circle.FixedPoints = append(circle.FixedPoints, [2]float64{x, y})
	}
	return &circle
}

func (c *Circle) fgetInterval(startIndex, endIndex int) [][2]float64 {
	numPoints := len(c.FixedPoints)
	point1 := c.FixedPoints[startIndex%numPoints]
	point2 := c.FixedPoints[endIndex%numPoints]
	dx, dy := point2[0]-point1[0], point2[1]-point1[1]
	theta1 := math.Atan2(dy, dx)
	if startIndex == numPoints-1 && endIndex == 0 {

		theta2 := math.Pi*2 + theta1
		return [][2]float64{{theta1, theta2}} // in radians
	} else {
		theta2 := math.Atan2(-dy, -dx)
		start := theta1 // in radians
		end := theta2   // in radians
		if math.Abs(end-start-math.Pi*2) < 1e-9 {
			return [][2]float64{{0, math.Pi * 2}} // entire circle
		} else if math.Abs(end-start) < 1e-9 {
			return [][2]float64{} // empty interval
		} else if start <= end {
			return [][2]float64{{start, end}}
		} else {
			return [][2]float64{{start, math.Pi * 2}, {0, end}} // half-open interval
		}
	}
}

func (c *Circle) getInterval(startIndex, endIndex int) [][2]float64 {
	numPoints := len(c.FixedPoints)
	point1 := c.FixedPoints[startIndex%numPoints]
	point2 := c.FixedPoints[endIndex%numPoints]
	dx, dy := point2[0]-point1[0], point2[1]-point1[1]
	theta1 := math.Atan2(dy, dx)

	// Check if we need to reverse the interval direction
	if math.IsNaN(theta1) && dx == 0 {
		// special case for closing the loop
		// if numPoint is 12
		// point1 := c.FixedPoints[11%12] = 11
		// point2 := c.FixedPoints[endIndex%numPoints] = 0
		// then theta1 Atan2(y<0, 0) = -Pi/2
		if dy > 0 {
			theta1 = math.Pi / 2
		} else if dy < 0 {
			theta1 = -math.Pi / 2
		} else {
			// This is NaN case handling.
			return [][2]float64{}
		}
	}
	if startIndex == numPoints-1 && endIndex == 0 { // special case for closing the loop
		theta2 := math.Pi*2 + theta1
		return [][2]float64{{theta1, theta2}} // in radians
	} else {
		theta2 := math.Atan2(-dy, -dx)
		start := theta1 // in radians
		end := theta2   // in radians
		if math.Abs(end-start-math.Pi*2) < 1e-9 {
			return [][2]float64{{0, math.Pi * 2}} // entire circle
		} else if math.Abs(end-start) < 1e-9 {
			return [][2]float64{} // empty interval
		} else if start <= end {
			return [][2]float64{{start, end}}
		} else {
			return [][2]float64{{start, math.Pi * 2}, {0, end}} // half-open interval
		}
	}
}

func isOnCircle(I, Iprime [][2]float64, radius float64) (bool, bool) {
	isOnI, isOnIprime := true, true
	for _, interval := range I {
		if interval[0] < 0 || interval[1] > 2*math.Pi*radius {
			isOnI = false
			break
		}
	}
	for _, interval := range Iprime {
		if interval[0] < 0 || interval[1] > 2*math.Pi*radius {
			isOnIprime = false
			break
		}
	}
	return isOnI, isOnIprime
}

func getAllPossibleUnions(givenCircle *Circle) [][2]float64 {
	// Collect pairwise intersections of all [Pi, Pj) segments
	intersections := make([][2]float64, 0)
	for i := 0; i < len(givenCircle.FixedPoints); i++ {
		for j := i + 1; j < len(givenCircle.FixedPoints); j++ {
			interval := givenCircle.getInterval(i, j)
			if len(interval) > 0 {
				intersections = append(intersections, interval...)
			}
		}
	}

	// Merge all pairwise intersections
	union := mergeIntervals(intersections)
	return union
}

func canMakeUnion(I, Iprime [][2]float64, givenUnions [][2]float64) bool {

	// Check if union is an interval
	if len(givenUnions) == 1 {
		start, end := math.Mod(givenUnions[0][0]*180/math.Pi, 360), math.Mod(givenUnions[0][1]*180/math.Pi, 360)
		if start <= end {
			// Check if union is a subset of I and I'
			for _, interval := range I {
				if interval[0] > start || interval[1] < end {
					return false
				}
			}
			for _, interval := range Iprime {
				if interval[0] > start || interval[1] < end {
					return false
				}
			}
			return true
		} else {
			// Check if union is a superset of I and I'
			for _, interval := range I {
				if interval[0] < start && interval[1] > end {
					return true
				}
			}
			for _, interval := range Iprime {
				if interval[0] < start && interval[1] > end {
					return true
				}
			}
		}
	}
	// union is 0 or more than 2 is false.
	return false
}

func mergeIntervals(intervals [][2]float64) [][2]float64 {
	if len(intervals) == 0 {
		return intervals
	}
	merged := [][2]float64{intervals[0]}
	for _, current := range intervals[1:] {
		if current[0] <= merged[len(merged)-1][1] {
			merged[len(merged)-1][1] = math.Max(merged[len(merged)-1][1], current[1])
		} else {
			merged = append(merged, current)
		}
	}
	return merged
}

func main() {
	center := [2]float64{0, 0}
	radius := 5.0
	numFixedPoints := 12
	circle := NewCircle(center, radius, numFixedPoints)
	I := [][2]float64{{2.0, 4.0}, {6.0, 9.0}}
	Iprime := [][2]float64{{1.0, 3.0}, {5.0, 7.0}}

	Iinterval, IprimeInterval := isOnCircle(I, Iprime, radius)

	if Iinterval || IprimeInterval {
		fmt.Println("Error: given intervals are not on the circle")
		return
	}

	unions := getAllPossibleUnions(circle)

	if len(unions) == 1 {
		start, end := math.Mod(unions[0][0]*180/math.Pi, 360), math.Mod(unions[0][1]*180/math.Pi, 360)
		if start <= end {
			fmt.Printf("The union of the intervals is (%.2f, %.2f).\n", start, end)
		} else {
			fmt.Printf("The union of the intervals is (%.2f, 360) U (0, %.2f).\n", start, end)
		}
	} else {
		fmt.Println("The union of the intervals is not an interval.")
	}
}
