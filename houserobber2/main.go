package houserobber2

/*
https://leetcode.com/problems/house-robber-ii/

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:

Input: [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
             because they are adjacent houses.
Example 2:

Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
						 Total amount you can rob = 1 + 3 = 4.

Approach:
	* To handle last value, check two different slices of the houses
		* Check from 0 to n - 1, and 1 to n
		* This way, never run into the case where checking a house that will "wrap around"
	* There are two states, robbing at current house or not robbing at current house
	* The max value of the robbing at last house is found by
		* calculating the current robbing value by adding the not robbing state value of previous house + current value
		* update the not robbing state value of previous house by looking at the previous rob value or the previous not rob value and selecting whichever is greater
		* selecting the last rob value or the last not robbing value, whichever is bigger
*/

func calculateMaxRobbery(nums []int) int {
	prevRobValue := 0
	prevNotRobValue := 0

	for _, num := range nums {
		currentRobValue := prevNotRobValue + num

		if prevRobValue > prevNotRobValue {
			prevNotRobValue = prevRobValue
		}
		prevRobValue = currentRobValue
	}

	if prevRobValue > prevNotRobValue {
		return prevRobValue
	}

	return prevNotRobValue
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	maxRobberyFromFirstIndex := calculateMaxRobbery(nums[:len(nums)-1])
	maxRobberFromSecondIndex := calculateMaxRobbery(nums[1:])

	if maxRobberyFromFirstIndex > maxRobberFromSecondIndex {
		return maxRobberyFromFirstIndex
	}

	return maxRobberFromSecondIndex
}
