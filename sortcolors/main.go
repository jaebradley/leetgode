package sortcolors

/*
https://leetcode.com/problems/sort-colors/

Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example:

Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Follow up:

A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
Could you come up with a one-pass algorithm using only constant space?

Approach:

* Keep track of three locations
	* The index of the last sorted 0
	* The index of the last sorted 1
	* The index of the first sorted 2
* Initially, these values are 0, 0, and length - 1, respectively
* While the last_sorted_one_index < first_sorted_two_index
	* If last_sorted_one_index is a 0, swap with last_sorted_zero_index, increment last_sorted_zero_index, and increment last_sorted_one_index as it's not the last sorted one any more
	* If it's a 1, increment last_sorted_one_index
	* If it's a 2, swap value with first_sorted_two_index, decrement the first_sorted_two_index
*/

func sortColors(nums []int) {
	var lastSortedZeroIndex int
	var lastSortedOneIndex int
	firstSortedTwoIndex := len(nums) - 1

	for lastSortedOneIndex <= firstSortedTwoIndex {
		if nums[lastSortedOneIndex] == 0 {
			nums[lastSortedOneIndex], nums[lastSortedZeroIndex] = nums[lastSortedZeroIndex], nums[lastSortedOneIndex]
			lastSortedZeroIndex++
			lastSortedOneIndex++
		} else if nums[lastSortedOneIndex] == 1 {
			lastSortedOneIndex++
		} else if nums[lastSortedOneIndex] == 2 {
			nums[lastSortedOneIndex], nums[firstSortedTwoIndex] = nums[firstSortedTwoIndex], nums[lastSortedOneIndex]
			firstSortedTwoIndex--
		}
	}
}
