package list_node

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

func DeleteNodeFromLinkedList(head **ListNode, position int) {
	k := 1

	if *head == nil {
		fmt.Println("List Empty")
		return
	}

	q := new(ListNode)

	p := *head
	if position == 1 {
		*head = (*head).Next
		p = nil
		return
	} else {
		for p != nil && k < position {
			k++
			q = p
			p = p.Next
		}

		if p == nil {
			fmt.Println("Position does not exist.")
		} else {
			q.Next = p.Next
			p = nil
		}
	}
}


func DeleteLinkedList(head **ListNode) {
	auxilaryNode := new(ListNode)

	iterator := *head

	for iterator != nil {
		auxilaryNode = iterator.Next
		iterator = nil
		iterator = auxilaryNode
	}

	*head = nil
}
