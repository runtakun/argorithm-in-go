package cll_node

import (
	"testing"
)

func generateTestCLLNode() *CLLNode {
	node5 := &CLLNode{5, nil}
	node4 := &CLLNode{4, node5}
	node3 := &CLLNode{3, node4}
	node2 := &CLLNode{2, node3}
	node1 := &CLLNode{1, node2}
	node5.Next = node1

	return node1
}

func TestCircularListLength(t *testing.T) {
	node := generateTestCLLNode()

	actual := CircularListLength(node)
	expected := 5
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}


func TestPrintCircularListData(t *testing.T) {
	node := generateTestCLLNode()

	PrintCircularListData(node)
}


func TestInsertAtEndInCLL(t *testing.T) {
	node := generateTestCLLNode()
	InsertAtEndInCLL(&node, 0)
	
	actual := CircularListLength(node)
	expected := 6
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}	

	actual = node.Next.Next.Next.Next.Next.Data
	expected = 0
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}

	actual = node.Next.Next.Next.Next.Next.Next.Data
	expected = 1
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}


func TestInsertAtBeginInCLL(t *testing.T) {
	node := generateTestCLLNode()
	InsertAtBeginInCLL(&node, 0)
	
	actual := CircularListLength(node)
	expected := 6
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}	

	actual = node.Data
	expected = 0
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}

	actual = node.Next.Next.Next.Next.Next.Next.Data
	expected = 0
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestDeleteLastNodeFromCLL(t *testing.T) {
	node := generateTestCLLNode()
	DeleteLastNodeFromCLL(&node)
	
	actual := CircularListLength(node)
	expected := 4
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}	

	actual = node.Next.Next.Next.Data
	expected = 4
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}

	actual = node.Next.Next.Next.Next.Data
	expected = 1
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestDeleteFrontNodeFromCLL(t *testing.T) {
	node := generateTestCLLNode()
	DeleteFrontNodeFromCLL(&node)
	
	actual := CircularListLength(node)
	expected := 4
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}	

	actual = node.Data
	expected = 2
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}

	actual = node.Next.Next.Next.Next.Data
	expected = 2
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}
