package twosum

/*
https://leetcode.com/problems/two-sum/

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

Approach:

* Create map
* For each value, key is diff of target - value and value is index
* If find a matching value, return key of first value and current value
*/

func twoSum(nums []int, target int) []int {
	seenValues := map[int]int{}
	for index, value := range nums {
		if seenIndex, ok := seenValues[value]; ok {
			return []int{seenIndex, index}
		}
		seenValues[target-value] = index
	}

	return []int{}
}
