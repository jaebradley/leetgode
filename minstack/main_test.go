package minstack

import "testing"

func TestPushSingleElement(t *testing.T) {
	stack := Constructor()
	stack.Push(1)

	top := stack.Top()

	if top != 1 {
		t.Fatalf("Expected top value in stack to be 1 and not %d", top)
	}
}

func TestPushingTwoElements(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)

	top := stack.Top()

	if top != 2 {
		t.Fatalf("Expected top value in stack to be 2 and not %d", top)
	}
}

func TestGetMinOfEmptyStack(t *testing.T) {
	stack := Constructor()
	min := stack.GetMin()

	if min != 0 {
		t.Fatalf("Expected minimum value to be zero for empty stack and not %d", min)
	}
}

func TestGetMinOfSingleElementStack(t *testing.T) {
	stack := Constructor()

	stack.Push(1)

	min := stack.GetMin()

	if min != 1 {
		t.Fatalf("Expected minimum value to be 1 and not %d", min)
	}
}

func TestGetMinOfTwoElementStackWhenFirstElementIsMinimum(t *testing.T) {
	stack := Constructor()

	stack.Push(1)
	stack.Push(2)

	min := stack.GetMin()

	if min != 1 {
		t.Fatalf("Expected minimum value to be 1 and not %d", min)
	}
}

func TestGetMinOfTwoElementStackWhenSecondElementIsMinimum(t *testing.T) {
	stack := Constructor()

	stack.Push(2)
	stack.Push(1)

	min := stack.GetMin()

	if min != 1 {
		t.Fatalf("Expected minimum value to be 1 and not %d", min)
	}
}

func TestGetMinWhenTwoElementStackPopsCurrentMinimum(t *testing.T) {
	stack := Constructor()

	stack.Push(2)
	stack.Push(1)
	stack.Pop()

	min := stack.GetMin()

	if min != 2 {
		t.Fatalf("Expected minimum value to be 2 and not %d", min)
	}
}

func TestGetMinWhenTwoElementStackPopsCurrentMaximum(t *testing.T) {
	stack := Constructor()

	stack.Push(1)
	stack.Push(2)
	stack.Pop()

	min := stack.GetMin()

	if min != 1 {
		t.Fatalf("Expected minimum value to be 1 and not %d", min)
	}
}

func TestGetMinWhenThreeElementStackArePushedOutOfOrder(t *testing.T) {
	stack := Constructor()

	stack.Push(-2)
	stack.Push(0)
	stack.Push(-1)

	min := stack.GetMin()

	if min != -2 {
		t.Fatalf("Expected min to be -2 instead of %d", min)
	}
}
