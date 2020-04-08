package movezeroes

/*
https://leetcode.com/problems/move-zeroes/

Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.

Approach:

* Use two pointers
* One that keeps track of last zero and one that keeps track of last non-zero
* If zero index value is a zero and non-zero index value is not a zero, swap values at index
* If current zero index is not a zero, increment it
* Increment current non-zero index to continue considering values in array
*/

func moveZeroes(nums []int) {
	var zeroIndex = 0
	var nonZeroIndex = 0

	for nonZeroIndex < len(nums) {
		if nums[zeroIndex] == 0 && nums[nonZeroIndex] != 0 {
			nums[nonZeroIndex], nums[zeroIndex] = nums[zeroIndex], nums[nonZeroIndex]
		}

		if nums[zeroIndex] != 0 {
			zeroIndex++
		}

		nonZeroIndex++
	}
}
