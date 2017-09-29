package cll_node

import (
	"fmt"
)

type CLLNode struct {
	Data int
	Next *CLLNode
}

func CircularListLength(head *CLLNode) int {
	current := head
	count := 0
	if head == nil {
		return 0
	}

	for {
		current = current.Next
		count++

		if current == head {
			break
		}
	}

	return count
}


func PrintCircularListData(head *CLLNode) {
	current := head

	if head == nil {
		return
	}

	for {
		fmt.Print("%d ", current.Data)

		current = current.Next

		if current == head {
			break
		}
	}
}

func InsertAtEndInCLL(head **CLLNode, data int) {
	current := *head

	newNode := new(CLLNode)
	newNode.Data = data

	for current.Next != *head {
		current = current.Next
	}
	newNode.Next = newNode

	if *head == nil {
		*head = newNode
	} else {
		newNode.Next = *head
		current.Next = newNode
	}
}

func InsertAtBeginInCLL(head **CLLNode, data int) {
	current := *head

	newNode := new(CLLNode)
	newNode.Data = data

	for current.Next != *head {
		current = current.Next
	}
	newNode.Next = newNode

	if *head == nil {
		*head = newNode
	} else {
		newNode.Next = *head
		current.Next = newNode
		*head = newNode
	}
}

func DeleteLastNodeFromCLL(head **CLLNode) {
	temp := *head
	current := *head

	if *head == nil {
		fmt.Println("List Empty")
		return
	}

	for current.Next != *head {
		temp = current
		current = current.Next
	}

	temp.Next = current.Next
	current = nil
}


func DeleteFrontNodeFromCLL(head **CLLNode) {
	current := *head

	if *head == nil {
		fmt.Println("List Empty")
		return
	}

	for current.Next != *head {
		current = current.Next
	}

	current.Next = (*head).Next
	*head = (*head).Next
}
