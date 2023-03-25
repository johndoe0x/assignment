package main

import (
	"fmt"
	"sort"
)

func main() {
	var start1, end1 int
	var start2, end2 int
	var fixedPoint int

	fmt.Print("Enter fixedPoints in Circle C: ")
	fmt.Scan(&fixedPoint)
	fixedPoints := Interval{
		fixedPoint - fixedPoint,
		fixedPoint - 1,
	}

	fmt.Print("Enter Interval I's start1, end1: ")
	fmt.Scan(&start1, &end1)
	i := Interval{
		start1,
		end1,
	}
	if isPointEqual(i) {
		return
	}
	fmt.Print("Enter Interval I prims's start2, end2: ")
	fmt.Scan(&start2, &end2)
	iPrime := Interval{
		start2,
		end2,
	}

	if isPointEqual(iPrime) {
		return
	}

	fmt.Printf(" Interval I's input: [%d, %d), Interval I prime's input: [%d, %d), Circle C's Fixed Points: %d",
		i.start, i.end, iPrime.start, iPrime.end, fixedPoint)

	if !inCircleC(i, iPrime, fixedPoint) {
		return
	}

	var unionRes Interval
	if areIntervalsEqual(i, iPrime) {
		fmt.Printf("Union is [%d, %d)", i.start, i.end)
	}
	if isUnionSameCircleC(i, iPrime, fixedPoints) {
		fmt.Printf("Union is [%d, %d)", i.start, i.end)
		return
	}
	matched, unionRes := isPointMatched(i, iPrime, fixedPoints)
	if matched {
		fmt.Printf("Union is [%d, %d)", unionRes.start, unionRes.end)
		return
	}
	include, unionRes := isIntervalIncluded(i, iPrime, fixedPoints)
	if include {
		fmt.Printf("Union is [%d, %d)", unionRes.start, unionRes.end)
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

func inCircleC(intervalI, intervalIPrime Interval, fixedPoint int) bool {

	allIntervalsPoints := gatheringIntervals(intervalI, intervalIPrime)
	fixedPointStart := fixedPoint - fixedPoint
	fixedPointEnd := fixedPoint - 1
	for i := 0; i < 4; i++ {
		if allIntervalsPoints[i] > fixedPointEnd || allIntervalsPoints[i] < fixedPointStart {
			fmt.Println("one of the interval points is not on the Circle C")
			return false
		}
	}
	return true
}

func areIntervalsEqual(intervalI, intervalIPrime Interval) bool {
	if intervalI.start == intervalIPrime.start && intervalI.end == intervalIPrime.end ||
		intervalI.start == intervalIPrime.end && intervalI.end == intervalIPrime.start {
		fmt.Println("Interval I & Interval I prime is Equal")
		return true
	}
	return false
}

func isUnionSameCircleC(intervalI, intervalIPrime, fixedPoints Interval) bool {
	allIntervalsPoints := gatheringIntervals(intervalI, intervalIPrime)
	sort.Ints(allIntervalsPoints)

	if allIntervalsPoints[0] == fixedPoints.start &&
		allIntervalsPoints[3] == fixedPoints.end {
		fmt.Println("Intervals are same as Circle C")
		return true
	}
	if isGoAroundCircle(intervalI, fixedPoints) && !isGoAroundCircle(intervalIPrime, fixedPoints) {
		IStartOverlapped := intervalI.start < intervalIPrime.end
		IEndOverlapped := intervalI.end > intervalIPrime.start
		if IStartOverlapped && IEndOverlapped {
			return true
		}
	}
	if !isGoAroundCircle(intervalI, fixedPoints) && isGoAroundCircle(intervalIPrime, fixedPoints) {
		iPrimeStartOverlapped := intervalIPrime.start < intervalI.end
		iPrimeEndOverlapped := intervalIPrime.end > intervalI.start
		if iPrimeStartOverlapped && iPrimeEndOverlapped {
			return true
		}
	}
	return false
}

func isGoAroundCircle(interval, fixedPoint Interval) bool {
	return interval.end > interval.start && interval.end > fixedPoint.start && interval.start < fixedPoint.end
}

func isPointMatched(intervalI, intervalIPrime, fixedPoint Interval) (bool, Interval) {

	var intervalRes Interval
	intervalIEndMatched := intervalI.end == intervalIPrime.start
	intervalIPrimeEndMatched := intervalIPrime.end == intervalI.start

	if intervalIEndMatched {
		if !isGoAroundCircle(intervalI, fixedPoint) && !isGoAroundCircle(intervalIPrime, fixedPoint) {
			if intervalI.start < intervalIPrime.end {
				intervalRes.start = intervalI.start
				intervalRes.end = intervalIPrime.end
				return true, intervalRes
			}
		}
		if isGoAroundCircle(intervalI, fixedPoint) && !isGoAroundCircle(intervalIPrime, fixedPoint) {
			if intervalIPrime.start < intervalIPrime.end {
				intervalRes.start = intervalI.start
				intervalRes.end = intervalIPrime.end
				return true, intervalRes
			}
		}

		if !isGoAroundCircle(intervalI, fixedPoint) && isGoAroundCircle(intervalIPrime, fixedPoint) {
			if intervalI.start < intervalI.end {
				intervalRes.start = intervalI.start
				intervalRes.end = intervalIPrime.end
				return true, intervalRes
			}
		}

	}
	if intervalIPrimeEndMatched {
		if !isGoAroundCircle(intervalI, fixedPoint) && !isGoAroundCircle(intervalIPrime, fixedPoint) {
			if intervalIPrime.start < intervalI.end {
				intervalRes.start = intervalIPrime.start
				intervalRes.end = intervalI.end
				return true, intervalRes
			}

		}
		if isGoAroundCircle(intervalI, fixedPoint) && !isGoAroundCircle(intervalIPrime, fixedPoint) {
			if intervalIPrime.start < intervalIPrime.end {
				intervalRes.start = intervalIPrime.start
				intervalRes.end = intervalI.end
				return true, intervalRes
			}
		}

		if !isGoAroundCircle(intervalI, fixedPoint) && isGoAroundCircle(intervalIPrime, fixedPoint) {
			if intervalI.start < intervalI.end {
				intervalRes.start = intervalIPrime.start
				intervalRes.end = intervalI.end
				return true, intervalRes
			}
		}
	}

	return false, Interval{start: -1, end: -1}
}

func isIntervalIncluded(intervalI, intervalIPrime, fixedPoint Interval) (bool, Interval) {
	// Not go around circle.
	startRangeEqual := intervalI.start == intervalIPrime.start
	endRangeEqual := intervalI.end == intervalIPrime.end
	if !isGoAroundCircle(intervalI, fixedPoint) && !isGoAroundCircle(intervalIPrime, fixedPoint) {
		// if res is interval I
		iStartRange := intervalI.start < intervalIPrime.start && intervalI.start < intervalIPrime.end
		iEndRange := intervalI.end > intervalIPrime.start && intervalI.end > intervalIPrime.end
		if iStartRange && iEndRange && !startRangeEqual && !endRangeEqual {
			return true, intervalI
		}
		// Include and Start point is equal
		if iEndRange && startRangeEqual && intervalI.start < intervalIPrime.end && !endRangeEqual {
			return true, intervalI
		}
		// Include and end point is equal
		if iStartRange && endRangeEqual && intervalI.start < intervalIPrime.start && !startRangeEqual {
			return true, intervalI
		}
		// if res is interval I Prime
		iPrimeStartRange := intervalIPrime.start < intervalI.start && intervalIPrime.start < intervalI.end
		iPrimeEndRange := intervalIPrime.end > intervalI.start && intervalIPrime.end > intervalI.end
		if iPrimeStartRange && iPrimeEndRange && !startRangeEqual && !endRangeEqual {
			return true, intervalIPrime
		}
		// Include and Start point is equal
		if iPrimeEndRange && startRangeEqual && intervalIPrime.start < intervalI.end && !endRangeEqual {
			return true, intervalIPrime
		}
		// Include and end point is equal
		if iPrimeStartRange && endRangeEqual && intervalIPrime.start < intervalI.start && !startRangeEqual {
			return true, intervalIPrime
		}
	}
	if isGoAroundCircle(intervalI, fixedPoint) && isGoAroundCircle(intervalIPrime, fixedPoint) {
		iStartRange := intervalI.start < intervalIPrime.start && intervalI.start > intervalIPrime.end
		iEndRange := intervalI.end < intervalIPrime.start && intervalI.end > intervalIPrime.end
		// if res is interval I
		if iStartRange && iEndRange && !startRangeEqual && !endRangeEqual {
			return true, intervalI
		}
		// Include and Start point is equal
		if iEndRange && startRangeEqual && intervalI.end > intervalIPrime.end && !endRangeEqual {
			return true, intervalI
		}
		// Include and end point is equal
		if iStartRange && endRangeEqual && intervalI.start < intervalIPrime.start && !startRangeEqual {
			return true, intervalI
		}
		// if res is interval I Prime
		iPrimeStartRange := intervalIPrime.start < intervalI.start && intervalIPrime.start > intervalI.end
		iPrimeEndRange := intervalIPrime.end < intervalI.start && intervalIPrime.end > intervalI.end
		if iPrimeStartRange && iPrimeEndRange && !startRangeEqual && !endRangeEqual {
			return true, intervalIPrime
		}
		// Include and Start point is equal
		if iPrimeEndRange && startRangeEqual && intervalIPrime.end > intervalI.end && !endRangeEqual {
			return true, intervalIPrime
		}
		// Include and end point is equal
		if iPrimeStartRange && endRangeEqual && intervalIPrime.start < intervalI.start && !startRangeEqual {
			return true, intervalIPrime
		}

	}
	if !isGoAroundCircle(intervalI, fixedPoint) && isGoAroundCircle(intervalIPrime, fixedPoint) {

		// interval I's endpoint is less than pivot ( pivot is 6 than endpoint can be less than 6,5,...)
		iPrimeStartRangeLessThanPivot := intervalIPrime.start < intervalI.start && intervalIPrime.start < intervalI.end
		iPrimeEndRangeLessThanPivot := intervalIPrime.end < intervalI.start && intervalIPrime.end < intervalI.end
		if iPrimeStartRangeLessThanPivot && iPrimeEndRangeLessThanPivot && !startRangeEqual && !endRangeEqual {
			return true, intervalIPrime
		}

		// Include and Start point is equal
		if iPrimeEndRangeLessThanPivot && startRangeEqual && intervalIPrime.end < intervalI.end && !endRangeEqual {
			return true, intervalIPrime
		}
		// Include and end point is equal
		if iPrimeStartRangeLessThanPivot && endRangeEqual && intervalIPrime.start < intervalI.start && !startRangeEqual {
			return true, intervalIPrime
		}
		// interval I's endpoint is more than pivot ( pivot is 6 than endpoint can be more than 1,2,...)
		iPrimeStartRangeMoreThanPivot := intervalIPrime.start > intervalI.start && intervalIPrime.start > intervalI.end
		iPrimeEndRangeMoreThanPivot := intervalIPrime.end > intervalI.start && intervalIPrime.end > intervalI.end
		if iPrimeStartRangeMoreThanPivot && iPrimeEndRangeMoreThanPivot && !startRangeEqual && !endRangeEqual {
			return true, intervalIPrime
		}

		// Include and Start point is equal
		if iPrimeEndRangeMoreThanPivot && startRangeEqual && intervalIPrime.end > intervalI.end && !endRangeEqual {
			return true, intervalIPrime
		}
		// Include and end point is equal
		if iPrimeStartRangeMoreThanPivot && endRangeEqual && intervalIPrime.start > intervalI.start && !startRangeEqual {
			return true, intervalIPrime
		}

	}

	if isGoAroundCircle(intervalI, fixedPoint) && !isGoAroundCircle(intervalI, fixedPoint) {
		iStartRangeLessThanPivot := intervalI.start < intervalIPrime.start && intervalI.start < intervalIPrime.end
		iEndRangeLessThanPivot := intervalI.end < intervalIPrime.start && intervalI.end < intervalIPrime.end

		if iStartRangeLessThanPivot && iEndRangeLessThanPivot && !startRangeEqual && !endRangeEqual {
			return true, intervalI
		}
		// Include and Start point is equal
		if iEndRangeLessThanPivot && startRangeEqual && intervalIPrime.end > intervalI.end && !endRangeEqual {
			return true, intervalI
		}
		// Include and end point is equal
		if iEndRangeLessThanPivot && endRangeEqual && intervalIPrime.start > intervalI.start && !startRangeEqual {
			return true, intervalI
		}
		iStartRangeMoreThanPivot := intervalI.start > intervalIPrime.start && intervalI.start > intervalIPrime.end
		iEndRangeMoreThanPivot := intervalI.end > intervalIPrime.start && intervalI.end > intervalIPrime.end
		if iStartRangeMoreThanPivot && iEndRangeMoreThanPivot && !startRangeEqual && !endRangeEqual {
			return true, intervalI
		}

		// Include and Start point is equal
		if iEndRangeMoreThanPivot && startRangeEqual && intervalIPrime.end < intervalI.end && !endRangeEqual {
			return true, intervalIPrime
		}
		// Include and end point is equal
		if iStartRangeMoreThanPivot && endRangeEqual && intervalIPrime.start < intervalI.start && !startRangeEqual {
			return true, intervalIPrime
		}
	}

	return false, Interval{start: -1, end: -1}
}

func isOverlapped(intervalI, intervalIPrime, fixedPoint Interval) (bool, Interval) {
	startShouldNotMatched := intervalI.start != intervalIPrime.start
	iStartEndShouldNotMatched := intervalI.start != intervalIPrime.end
	iPrimeEndStartShouldNotMatched := intervalIPrime.start != intervalI.end
	if !isGoAroundCircle(intervalI, fixedPoint) && !isGoAroundCircle(intervalIPrime, fixedPoint) {
		// I < I-prime case
		if startShouldNotMatched && intervalI.end > intervalIPrime.start && intervalIPrime.end > intervalIPrime.start {
			return true,
				Interval{
					intervalI.start,
					intervalIPrime.end,
				}
		}
		// I < I-prime case
		if startShouldNotMatched && intervalI.end < intervalIPrime.start && intervalI.end > intervalI.start {
			return true,
				Interval{
					intervalIPrime.start,
					intervalI.end,
				}
		}
	}

	if isGoAroundCircle(intervalI, fixedPoint) && isGoAroundCircle(intervalIPrime, fixedPoint) {
		if intervalI.start < intervalIPrime.start && intervalI.end < intervalIPrime.end {
			return true,
				Interval{
					intervalI.start,
					intervalIPrime.end,
				}
		}
		if intervalI.start > intervalIPrime.start && intervalI.end > intervalIPrime.end {
			return true,
				Interval{
					intervalIPrime.start,
					intervalI.end,
				}
		}
	}

	if !isGoAroundCircle(intervalI, fixedPoint) && isGoAroundCircle(intervalIPrime, fixedPoint) {
		if iStartEndShouldNotMatched && intervalIPrime.end < intervalI.start {
			return true,
				Interval{
					intervalIPrime.start,
					intervalI.end,
				}
		}

		if iPrimeEndStartShouldNotMatched && intervalI.end > intervalIPrime.start {
			return true,
				Interval{
					intervalI.start,
					intervalIPrime.end,
				}
		}
	}
	if isGoAroundCircle(intervalI, fixedPoint) && !isGoAroundCircle(intervalIPrime, fixedPoint) {
		if iPrimeEndStartShouldNotMatched && intervalIPrime.start > intervalI.end {
			return true,
				Interval{
					intervalI.start,
					intervalIPrime.end,
				}
		}

		if iStartEndShouldNotMatched && intervalIPrime.end > intervalI.start {
			return true,
				Interval{
					intervalIPrime.start,
					intervalI.end,
				}
		}
	}

	return false, Interval{start: -1, end: -1}
}
