package houserobber

/*
https://leetcode.com/problems/house-robber/

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:

Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
						 Total amount you can rob = 2 + 9 + 1 = 12.

Approach:

* Max amount at house[i] is
* Max amount if didn't steal at house[i - 1] + amount at house[i]
* Max amount between values at house[i - 1]

[
	0 1 2 4 3
	0 0 1 2 4
]

[
	0 2 7 11 10 12
	0 0 2  7 11 11
]
*/

func max(first int, second int) int {
	if first > second {
		return first
	}

	return second
}

func rob(nums []int) int {
	previousHouseStealValue := 0
	previousHouseDidntStealValue := 0

	for _, num := range nums {
		nextHouseStealValue := previousHouseDidntStealValue + num
		nextPreviousHouseDidntStealValue := max(previousHouseStealValue, previousHouseDidntStealValue)
		previousHouseStealValue = nextHouseStealValue
		previousHouseDidntStealValue = nextPreviousHouseDidntStealValue
	}

	return max(previousHouseStealValue, previousHouseDidntStealValue)
}
