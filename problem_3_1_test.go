package main

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := &Stack{Capacity: 3}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	actual := stack.Size()
	expected := 3
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestPop(t *testing.T) {
	stack := &Stack{Capacity: 3}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	actual := stack.Pop()
	expected := 3
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}

	actual = stack.Size()
	expected = 2
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestTop(t *testing.T) {
	stack := &Stack{Capacity: 3}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	actual := stack.Top()
	expected := 3
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}

	actual = stack.Size()
	expected = 3
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestSize(t *testing.T) {
	stack := &Stack{Capacity: 3}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	actual := stack.Size()
	expected := 3
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestIsEmptyStack(t *testing.T) {
	stack1 := &Stack{Capacity: 3}
	
	actual := stack1.IsEmptyStack()
	expected := true
	if actual != expected {
		t.Errorf("stack1: got %v, want %v", actual, expected)
	}

	stack2 := &Stack{Capacity: 3}
	stack2.Push(1)
	stack2.Push(2)
	stack2.Push(3)
	
	actual = stack2.IsEmptyStack()
	expected = false
	if actual != expected {
		t.Errorf("stack2: got %v, want %v", actual, expected)
	}
}


func TestIsFullStack(t *testing.T) {
	stack1 := &Stack{Capacity: 3}
	stack1.Push(1)
	stack1.Push(2)
	
	actual := stack1.IsFullStack()
	expected := false
	if actual != expected {
		t.Errorf("stack1: got %v, want %v", actual, expected)
	}

	stack2 := &Stack{Capacity: 3}
	stack2.Push(1)
	stack2.Push(2)
	stack2.Push(3)
	
	actual = stack2.IsFullStack()
	expected = true
	if actual != expected {
		t.Errorf("stack2: got %v, want %v", actual, expected)
	}
}
