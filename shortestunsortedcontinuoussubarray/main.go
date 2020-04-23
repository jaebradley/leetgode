package shortestunsortedcontinuoussubarray

import (
	"math"
)

/*
https://leetcode.com/problems/shortest-unsorted-continuous-subarray/

Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order, too.

You need to find the shortest such subarray and output its length.

Example 1:

Input: [2, 6, 4, 8, 10, 9, 15]
Output: 5
Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
Note:

Then length of the input array is in range [1, 10,000].
The input array may contain duplicates, so ascending order here means <=.

Approach:

* Start from beginning of array - keep track of the last index where the value at that index is less than the maximum value seen
	up to that point.
	This value will be the end value.
* Start from end of array - keep track of the last index where the value at that index is greater than the min seen up to that point.
	This value will be the start value.
* Set startIndex and endIndex to -1 and -2 respectively to handle zero state (i.e. no nums)
*/

func findUnsortedSubarray(nums []int) int {
	startIndex := -1
	endIndex := -2

	maximum := math.MinInt64
	minimum := math.MaxInt64

	for index, num := range nums {
		if num > maximum {
			maximum = num
		}

		if num < maximum {
			endIndex = index
		}
	}

	for j := len(nums) - 1; j >= 0; j = j - 1 {
		num := nums[j]
		if num < minimum {
			minimum = num
		}

		if num > minimum {
			startIndex = j
		}
	}

	return endIndex - startIndex + 1
}
