package main

import (
	"fmt"
)

type ListNode struct {
	Data int
	Next *ListNode
}

func ListLength(head *ListNode) int {
	current := head
	count := 0
	for current != nil {
		count++
		current = current.Next
	}

	return count
}

func InsertInLinkedList(head **ListNode, data, position int) {
	k := 1

	if *head == nil {
		fmt.Println("List Empty")
		return
	}

	q := new(ListNode)

	newNode := new(ListNode)
	newNode.Data = data

	p := *head
	if position == 1 {
		newNode.Next = p
		*head = newNode
	} else {
		for p != nil && k < position {
			k++
			q = p
			p = p.Next
		}
		q.Next = newNode
		newNode.Next = p
	}
}
