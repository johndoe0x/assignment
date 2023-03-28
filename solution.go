package main

import (
	"fmt"
)

type Interval struct {
	start, end int
	isGoAround bool
}

func main() {
	var start1, end1 int
	var start2, end2 int
	var fixedPoint int

	fmt.Print("Enter fixedPoints in Circle C: ")
	fmt.Scan(&fixedPoint)
	fixedPoints := Interval{
		fixedPoint - fixedPoint,
		fixedPoint - 1,
		true,
	}

	fmt.Print("Enter Interval I's start1, end1: ")
	fmt.Scan(&start1, &end1)
	i := Interval{
		start1 % fixedPoint,
		end1 % fixedPoint,
		start1-end1 > 0,
	}
	if isPointEqual(i) {
		return
	}
	fmt.Print("Enter Interval I prims's start2, end2: ")
	fmt.Scan(&start2, &end2)
	iPrime := Interval{
		start2 % fixedPoint,
		end2 % fixedPoint,
		start2-end2 > 0,
	}

	if isPointEqual(iPrime) {
		return
	}

	fmt.Printf(" Interval I's input: [%d, %d), Interval I prime's input: [%d, %d), Circle C's Fixed Points: %d \n",
		i.start, i.end, iPrime.start, iPrime.end, fixedPoint)

	var unionRes Interval
	equal, res := isSame(i, iPrime)
	if equal {
		fmt.Printf("Union is [%d, %d)", res.start, res.end)
		return
	}
	if isUnionSameCircleC(i, iPrime, fixedPoints) {
		fmt.Printf("Union is same as circle [%d, %d)", fixedPoints.start, fixedPoints.end)
		return
	}

	matched, unionRes := hasPointMatchedUnion(i, iPrime)
	if matched {
		fmt.Printf("Union is [%d, %d)", unionRes.start, unionRes.end)
		return
	}

	include, unionRes := isIntervalIncluded(i, iPrime, fixedPoints)
	if include {
		fmt.Printf("intervals are includes . Union is [%d, %d)", unionRes.start, unionRes.end)
		return
	}

	overlapped, unionRes := isOverlapped(i, iPrime, fixedPoints)
	if overlapped {
		fmt.Printf("Union is [%d, %d)", unionRes.start, unionRes.end)
		return
	}

	fmt.Print("Union in-exists. ")
	return

}

func gatheringIntervals(intervalI, intervalIPrime Interval) []int {
	res := make([]int, 4)
	res[0] = intervalI.start
	res[1] = intervalI.end

	res[2] = intervalIPrime.start
	res[3] = intervalIPrime.end
	return res
}

func isPointEqual(interval Interval) bool {
	if interval.start == interval.end {
		fmt.Println("Pi and Pj can't be equal")
		return true
	}
	return false
}

func isSame(interval1, interval2 Interval) (bool, Interval) {
	if interval1 == interval2 {
		fmt.Println("Both interval is same")
		return true, interval1
	}
	return false, interval1
}

func isApointMatched(interval1, interval2 Interval) (bool, int) {
	arr := make([]int, 2)
	arr[0] = interval2.start
	arr[1] = interval2.end
	for _, v := range arr {
		if v == interval1.start || v == interval1.end {
			fmt.Println("point is matched")
			return true, v
		}
	}
	return false, -1
}

func isUnionSameCircleC(intervalI, intervalIPrime, fixedPoints Interval) bool {
	matched, _ := isApointMatched(intervalI, intervalIPrime)
	if matched {
		if intervalI.start == intervalIPrime.end && intervalI.end == intervalIPrime.start {
			return true
		}
		if intervalI.end == intervalIPrime.start && !intervalI.isGoAround && intervalIPrime.isGoAround {
			return intervalI.start < intervalIPrime.end && intervalI.end > intervalIPrime.end

		}

		if intervalI.start == intervalIPrime.end && intervalI.isGoAround && !intervalIPrime.isGoAround {
			return intervalI.end < intervalIPrime.start && intervalI.end < intervalIPrime.end
		}
	} else {

		if intervalI.isGoAround && !intervalIPrime.isGoAround {
			//in this  case intervalI  has to set inside of intervalPrime.
			if intervalI.start > intervalIPrime.start && intervalI.end > intervalIPrime.start &&
				//check less than biggest fixed point
				intervalI.start < fixedPoints.end && intervalI.end < fixedPoints.end {
				return true
			} else if intervalIPrime.start < intervalI.end && intervalIPrime.end < intervalI.end &&
				//check bigger than biggest fixed point
				intervalIPrime.start > fixedPoints.start && intervalIPrime.end > fixedPoints.start {
				return true
			}
		} else if !intervalI.isGoAround && intervalIPrime.isGoAround {
			//in this  case intervalPrime has to set inside of intervalI.
			if intervalIPrime.start > intervalI.start && intervalIPrime.end > intervalI.start &&
				//check less than the biggest fixed point
				intervalIPrime.start < fixedPoints.end && intervalIPrime.end < fixedPoints.end {
				return true
			} else if intervalIPrime.start < intervalI.end && intervalIPrime.end < intervalI.end &&
				//check bigger than biggest fixed point
				intervalIPrime.start > fixedPoints.start && intervalIPrime.end > fixedPoints.start {
				return true
			}
		} else if intervalI.isGoAround && intervalIPrime.isGoAround {
			whoIsBigger := intervalI.start-intervalI.end > intervalIPrime.start-intervalI.end
			// here is the two case, one is one interval is go around circle but less than the biggest fixed point && bigger than another interval's start
			// here is the two case, one is one interval is go around circle but bigger than the smallest fixed point && smaller tha anotjer interval's end

			// if true interval prime has less radian than interval prime
			if whoIsBigger {
				return intervalI.start < intervalIPrime.start && intervalI.end > intervalIPrime.end
				// if false  interval i has less radian than interval i
			} else {
				return intervalIPrime.start < intervalI.start && intervalI.end < intervalIPrime.end
			}

		}
	}
	return false
}
func MinOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func MaxOf(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

func hasPointMatchedUnion(intervalI, intervalIPrime Interval) (bool, Interval) {
	matched, v := isApointMatched(intervalI, intervalIPrime)
	if matched {

		if v == intervalI.start && v == intervalIPrime.start {
			if intervalI.end > intervalIPrime.end {
				return true, intervalI
			} else {
				return true, intervalIPrime
			}

		} else if v == intervalI.end && v == intervalIPrime.end {
			if intervalI.start > intervalIPrime.end {
				return true, intervalIPrime
			} else {
				return true, intervalI
			}
		} else if v == intervalI.end && v == intervalIPrime.start || v == intervalI.start && v == intervalIPrime.end {
			if !intervalI.isGoAround && !intervalIPrime.isGoAround {

				return true, Interval{
					MinOf(intervalIPrime.end, intervalI.start),
					MaxOf(intervalIPrime.end, intervalI.start),
					false,
				}
			} else {
				return true, Interval{
					MaxOf(intervalIPrime.end, intervalI.start),
					MinOf(intervalIPrime.end, intervalI.start),
					true,
				}
			}
		}
	}

	return false, Interval{start: -1, end: -1}
}

func checkInclusion(interval1, interval2 Interval) bool {
	for i := interval1.start + 1; i < interval1.end; i++ {
		incldStart := i == interval2.start
		incldEnd := i == interval2.end
		if incldStart && incldEnd {
			return true
		}
	}
	return false
}

func isIntervalIncluded(intervalI, intervalIPrime, fixedPoint Interval) (bool, Interval) {
	matched, _ := isApointMatched(intervalI, intervalIPrime)
	checkPointMatchedUnion, v := hasPointMatchedUnion(intervalI, intervalIPrime)
	if !matched && !checkPointMatchedUnion {

		if !intervalI.isGoAround && !intervalIPrime.isGoAround {
			if checkInclusion(intervalI, intervalIPrime) {
				return true, intervalI
			} else if checkInclusion(intervalIPrime, intervalI) {
				return true, intervalIPrime
			}
		} else {
			if intervalI.isGoAround && intervalIPrime.isGoAround {
				if intervalI.start < intervalIPrime.start && intervalIPrime.end < intervalI.end {
					return true, intervalI
				} else if intervalIPrime.start < intervalI.start && intervalI.end < intervalIPrime.end {
					return true, intervalIPrime
				}

			} else if intervalI.isGoAround && !intervalIPrime.isGoAround {
				if intervalIPrime.start > intervalI.start && intervalIPrime.end > intervalI.start &&
					intervalIPrime.start < fixedPoint.end && intervalIPrime.end < fixedPoint.end {
					return true, intervalI
				} else if intervalIPrime.start < intervalI.end && intervalIPrime.end < intervalI.end &&
					intervalIPrime.start > fixedPoint.start && intervalIPrime.end < fixedPoint.start {
					return true, intervalI
				} else {
					fmt.Println("not inclusion")
					return false, Interval{start: -1, end: -1}
				}

			} else if !intervalI.isGoAround && intervalIPrime.isGoAround {
				if intervalI.start > intervalIPrime.start && intervalI.end > intervalIPrime.start &&
					intervalI.start < fixedPoint.end && intervalI.end < fixedPoint.end {
					return true, intervalIPrime
				} else if intervalI.start < intervalIPrime.end && intervalI.end < intervalIPrime.end &&
					intervalI.start > fixedPoint.start && intervalI.end < fixedPoint.start {
					return true, intervalIPrime
				} else {
					fmt.Println("not inclusion")
					return false, Interval{start: -1, end: -1}
				}
			}
		}
	} else {
		fmt.Println("Error Point matched exists.")
		return false, v
	}
	return false, Interval{start: -1, end: -1}
}
func isOverlapped(intervalI, intervalIPrime, fixedPoint Interval) (bool, Interval) {
	if !intervalI.isGoAround && !intervalIPrime.isGoAround {
		if intervalI.end > intervalIPrime.start && intervalI.end < intervalIPrime.end {
			return true, Interval{start: intervalI.start, end: intervalIPrime.end, isGoAround: false}
		}

	} else if intervalI.isGoAround && !intervalIPrime.isGoAround {
		// overlapped in start

		// overlapped in end
	} else if !intervalI.isGoAround && intervalIPrime.isGoAround {
		// overlapped in start

		// overlapped in end
	}
	return false, Interval{start: -1, end: -1}
}
