package validparentheses

/*
https://leetcode.com/problems/valid-parentheses/

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true

Approach:

* Use a stack to capture open brackets
* When a closed bracket is seen, pop off stack
* If stack is empty and closed bracket is seen, return false
* If open bracket matches closed bracket, continue
* If open bracket does not match closed bracket, return false
* Return if stack is empty at end
*/

func isValid(s string) bool {
	openParentheses := []rune{}

	for _, value := range s {
		if value == '(' || value == '[' || value == '{' {
			openParentheses = append(openParentheses, value)
		} else if value == ')' || value == ']' || value == '}' {
			if len(openParentheses) < 1 {
				return false
			}

			lastOpenParentheses := openParentheses[len(openParentheses)-1]
			if (lastOpenParentheses == '(' && value == ')') ||
				(lastOpenParentheses == '[' && value == ']') ||
				(lastOpenParentheses == '{' && value == '}') {
				openParentheses = openParentheses[:len(openParentheses)-1]
			} else {
				return false
			}
		}
	}

	return len(openParentheses) == 0
}
