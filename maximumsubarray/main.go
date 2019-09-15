package maximumsubarray

/*
https://leetcode.com/problems/maximum-subarray/

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

Approach:

* Start at index 0
* Start summing until index where value at index is greater than current sum
* At index greater than current sum, restart start index at that index
* Keep track of sum and update max when new max is found
*/

func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]
	for index := 1; index < len(nums); index++ {
		currentValue := nums[index]
		nextSum := currentSum + currentValue
		if currentValue > nextSum {
			currentSum = currentValue
		} else {
			currentSum = nextSum
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
