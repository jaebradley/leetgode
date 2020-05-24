package dailytemperatures

/*
https://leetcode.com/problems/daily-temperatures/

Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature. If there is no future day for which this is possible, put 0 instead.

For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].

Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100].

Approach:

* Create a stack that will hold temperatures
* Starting from the beginning, if the stack isn't empty check the current temperature against the last temperature added to the stack (really, it's the index of the last temperature added to the stack)
	* If the current temperature is greater than the peeked index, update the value of the peeked index in the output to be diff between current index and peeked index
	* In other words, the first time a greater temperature is found for top element, pop that element and get diff in index
		* The diff represents how long the index had to wait until a greater temperature
* If the stack is empty or the current temperature is not greater than the peeked temperature (i.e. temperature is <= last temperature that was added) add that temperature
	* So pushed temperatures are never increasing
*/

type stack struct {
	Values []int
}

func (S *stack) push(value int) {
	updatedValues := append(S.Values, value)
	S.Values = updatedValues
}

func (S *stack) pop() int {
	value, updatedValues := S.Values[len(S.Values)-1], S.Values[:len(S.Values)-1]
	S.Values = updatedValues
	return value
}

func (S stack) peek() int {
	return S.Values[len(S.Values)-1]
}

func (S stack) isEmpty() bool {
	return len(S.Values) == 0
}

func dailyTemperatures(T []int) []int {
	daysToWait := make([]int, len(T))
	s := stack{Values: []int{}}

	for i := 0; i < len(T); i++ {
		for !s.isEmpty() && T[i] > T[s.peek()] {
			lastIndex := s.pop()
			daysToWait[lastIndex] = i - lastIndex
		}
		s.push(i)
	}

	return daysToWait
}
