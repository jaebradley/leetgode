package singlenumber

/*
https://leetcode.com/problems/single-number/

Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,1]
Output: 1
Example 2:

Input: [4,1,2,1,2]
Output: 4

Approach:

Using XOR flips bits.
XOR is also commutative.
Also A XOR A = 0.
So starting with 0, XORing each value should eventually get down to the value that doesn't have a pair.
*/

func singleNumber(nums []int) int {
	result := 0
	for _, value := range nums {
		result ^= value
	}
	return result
}
