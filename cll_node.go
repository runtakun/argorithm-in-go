package main

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
