package coinchange

import "math"

/*
https://leetcode.com/problems/coin-change/

You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Note:
You may assume that you have an infinite number of each kind of coin.

Approach:
	* Create an array from 0 to total amount
	* Set the value for the coin count in the array to max int except for the first index
	* Iterate from 1 to the amount
	* For each coin, subtract the coin value from the current amount
	* If the current amount is >= 0, check to see what the previous coin count was for that previous value
	* This is how coin counts can "stack" - the minimum coin count for smaller values can be used to calculate the minimum coin counts at larger values
	* If the value at the end of the array is still max int, then there's no coin combination to reach the amount and return -1
*/

func coinChange(coins []int, amount int) int {
	counts := make([]int, amount+1)

	for i := 1; i < len(counts); i++ {
		counts[i] = math.MaxInt32
	}

	for i := 1; i < len(counts); i++ {
		for _, coin := range coins {
			previousValue := i - coin

			if previousValue >= 0 {
				coinCountAtPreviousValue := counts[previousValue]
				nextCoinCount := coinCountAtPreviousValue + 1

				if nextCoinCount < counts[i] {
					counts[i] = nextCoinCount
				}
			}
		}
	}

	if counts[len(counts)-1] == math.MaxInt32 {
		return -1
	}

	return counts[len(counts)-1]
}
