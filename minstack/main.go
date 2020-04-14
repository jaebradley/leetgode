package minstack

/*
https://leetcode.com/problems/min-stack/

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.


Example:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.

Approach:

* Build a Stack as a singly linked-list built from Nodes
* There's a head Node
* When elements get pushed or popped, the minimum value is changed if the pushed element has a lower value than the minimum or if the popped element has a value equal to the minimum
	* In the pushed case, the minimum value is changed
	* In the popped case, the minimum value is set to the new head Node minimum value
* Whenever an element is pushed, the Node also has a reference to the current minimum value at that time
*/

// MinStack is a Stack that also looks up the minimum value in the Stack in O(1)
type MinStack struct {
	head *node
}

type node struct {
	next    *node
	value   int
	minimum int
}

func min(first int, second int) int {
	if first < second {
		return first
	}

	return second
}

// Constructor initialize your data structure here.
func Constructor() MinStack {
	return MinStack{head: nil}
}

// Push adds value to head of stack in O(1)
func (stack *MinStack) Push(x int) {
	nextMinimum := x

	if stack.head != nil {
		nextMinimum = min(nextMinimum, stack.head.minimum)
	}

	nextHead := node{
		next:    stack.head,
		value:   x,
		minimum: nextMinimum,
	}

	stack.head = &nextHead
}

// Pop removes value from head of stack in O(1)
func (stack *MinStack) Pop() {
	if stack.head != nil {
		stack.head = stack.head.next
	}
}

// Top checks the head of stack's value in O(1)
func (stack *MinStack) Top() int {
	var value int

	if stack.head != nil {
		value = stack.head.value
	}

	return value
}

// GetMin returns the minimum value in stack
func (stack *MinStack) GetMin() int {
	var minimum int

	if stack.head != nil {
		minimum = stack.head.minimum
	}

	return minimum
}
