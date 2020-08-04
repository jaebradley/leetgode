package besttimetobuyandsellstockwithtransactionfee

/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

Your are given an array of integers prices, for which the i-th element is the price of a given stock on day i; and a non-negative integer fee representing a transaction fee.

You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction. You may not buy more than 1 share of a stock at a time (ie. you must sell the stock share before you buy again.)

Return the maximum profit you can make.

Example 1:
Input: prices = [1, 3, 2, 8, 4, 9], fee = 2
Output: 8
Explanation: The maximum profit can be achieved by:
Buying at prices[0] = 1
Selling at prices[3] = 8
Buying at prices[4] = 4
Selling at prices[5] = 9
The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
Note:

0 < prices.length <= 50000.
0 < prices[i] < 50000.
0 <= fee < 50000.

Approach:
* At any given interval, can choose to hold, buy (if haven't purchased), or sell (if have purchased)
* To calculate max of buying at time period i it's the max of selling in previous time period - cost of buying at present time period OR holding (previous time period's buy value)
	* max value @ i = max(sell[i - 1] - cost[i], buy[i - 1])
* To calculate max of selling at time period i, it's the max of the buying value from previous time period + selling in current time period, or holding previous sell value
	* max value @ i = max(buy[i - 1] + price[i]  - fee, sell[i - 1])
* To get max value, get the selling value in the last time period

*/

func maxProfit(prices []int, fee int) int {
	dp := make([][2]int, len(prices))

	for i, price := range prices {
		if i == 0 {
			dp[i][0] = -price
			dp[i][1] = 0
		} else {
			prevBuy := dp[i-1][0]
			prevSell := dp[i-1][1]

			currentBuy := prevSell - price

			if currentBuy > prevBuy {
				dp[i][0] = currentBuy
			} else {
				dp[i][0] = prevBuy
			}

			currentSell := prevBuy + price - fee

			if currentSell > prevSell {
				dp[i][1] = currentSell
			} else {
				dp[i][1] = prevSell
			}
		}
	}

	return dp[len(dp)-1][1]
}
