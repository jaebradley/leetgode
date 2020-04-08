package majorityelement

/*
https://leetcode.com/problems/majority-element/

Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:

Input: [3,2,3]
Output: 3

Example 2:

Input: [2,2,1,1,1,2,2]
Output: 2

Approach:

* Could use a hashmap and store the counts
* Other approach is the Boyer-Moore Majority Vote Algorithm
* The way I think about it is that you iterate through all the values in the array
	* For the first value, set the count to 1
	* For the next value, if it's the same as the first value, increment the count
	* If it's not the same as the first value, decrement the count
	* If the count is zero, set the current value as the current proposed majority element and increment the count
	* At the end, if the majority element actually exists, it should be the value that's set since there are more of those element than any other element
		* so there will be at least time where the majority element will the current proposed majority element and there will not be enough other elements
		  to decrement it
*/

func majorityElement(nums []int) int {
	var currentMajorityElement = nums[0]
	var currentMajorityElementCount = 1

	for _, element := range nums[1:] {
		if currentMajorityElementCount == 0 {
			currentMajorityElement = element
			currentMajorityElementCount++
		} else if element == currentMajorityElement {
			currentMajorityElementCount++
		} else if element != currentMajorityElement {
			currentMajorityElementCount--
		}
	}

	return currentMajorityElement
}
