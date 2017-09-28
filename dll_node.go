package main

import (
	"fmt"
)

type DLLNode struct {
	Data int
	Next *DLLNode
	Prev *DLLNode
}

func DLLLength(head *DLLNode) int {
	current := head
	count := 0
	for current != nil {
		count++
		current = current.Next
	}

	return count
}

func DLLInsert(head **DLLNode, data, position int) {
	k := 1

	newNode := new(DLLNode)

	newNode.Data = data

	if position == 1 {
		newNode.Next = *head
		newNode.Prev = nil

		if (*head != nil) {
				(*head).Prev = newNode
		}
		*head = newNode
		return
	}

	temp := *head
	for k < position - 1 && temp.Next != nil {
		temp = temp.Next
		k++
	}

	if k != position {
		fmt.Println("Desired position does not exist")
	}

	newNode.Next = temp.Next
	newNode.Prev = temp

	if (temp.Next != nil) {
		temp.Next.Prev = newNode
		temp.Next = newNode
	}
}
