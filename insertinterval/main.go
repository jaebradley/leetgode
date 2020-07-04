package insertinterval

/*
https://leetcode.com/problems/insert-interval/

Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

Approach:

* Non-overlapping means every start does not collide with any other start or end and any end does not overlap with many start / any other end
* Find interval before start of new interval
* From that interval, keep track of minimum start index for intervals that overlap
* Stop when reach an interval that does not overlap minimum start index and maximum end index of any overlapping intervals
*/

func insert(intervals [][]int, newInterval []int) [][]int {
	index := 0
	intervalStartValue := newInterval[0]
	intervalEndValue := newInterval[1]
	updatedIntervals := [][]int{}

	for index < len(intervals) && intervals[index][1] < newInterval[0] {
		updatedIntervals = append(updatedIntervals, intervals[index])
		index++
	}

	for index < len(intervals) && intervals[index][0] <= newInterval[1] {
		currentInterval := intervals[index]
		if currentInterval[0] < intervalStartValue {
			intervalStartValue = currentInterval[0]
		}

		if currentInterval[1] > intervalEndValue {
			intervalEndValue = currentInterval[1]
		}
		index++
	}

	updatedIntervals = append(updatedIntervals, []int{intervalStartValue, intervalEndValue})

	for index < len(intervals) {
		updatedIntervals = append(updatedIntervals, intervals[index])
		index++
	}

	return updatedIntervals
}
