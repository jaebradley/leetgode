package findallnumbersdisappearedinanarray

/*
https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]

Approach:

For each value in the input array, use that value to reference an index in the array.

For the value at the index, reset it to a negative value.

Iterate through values again, and for indices that do not have a negative value,
these are the values that do not appear in the array.

For any missing values, their index value would not be in the input array.
So the value at their index value will be positive.
*/

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		seenIndex := abs(num) - 1
		seenValue := nums[seenIndex]

		if seenValue > 0 {
			nums[seenIndex] = -seenValue
		}
	}

	results := make([]int, 0)

	for index, num := range nums {
		if num > 0 {
			results = append(results, index+1)
		}
	}

	return results
}
