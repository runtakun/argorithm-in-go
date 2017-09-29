package main

import (
	"os"
	"github.com/runtakun/argorithm-in-go/list-node"
)

type Stack struct {
	Head *list_node.ListNode
	Capacity int
}

func (stack *Stack) Push(data int) {
	if stack.IsFullStack() {
		os.Stderr.WriteString("Stack is full\n")
		return
	}

	newNode := new(list_node.ListNode)
	newNode.Data = data

	if stack.Head == nil {
		stack.Head = newNode
	} else {
		p := stack.Head
		newNode.Next = p
		stack.Head = newNode
	}
}

func (stack *Stack) Pop() int {
		if stack.Head == nil {
			os.Stderr.WriteString("Stack is empty\n")
			return -1
		}

		data := stack.Head.Data
		stack.Head = stack.Head.Next
		return data
}

func (stack *Stack) Top() int {
	if stack.Head == nil {
		os.Stderr.WriteString("Stack is empty\n")
		return -1
	}

	return stack.Head.Data
}

func (stack *Stack) Size() int {
	if stack.Head == nil {
		return 0
	}

	count := 0
	current := stack.Head
	for current != nil {
		count++
		current = current.Next
	}

	return count
}

func (stack *Stack) IsEmptyStack() bool {
	return stack.Size() == 0
}

func (stack *Stack) IsFullStack() bool {
	return stack.Size() == stack.Capacity
}
