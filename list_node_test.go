package main

import (
	"testing"
	// "fmt"
)

func generateTestNode() *ListNode {
	node5 := &ListNode{4, nil}
	node4 := &ListNode{3, node5}
	node3 := &ListNode{2, node4}
	node2 := &ListNode{1, node3}
	node1 := &ListNode{0, node2}
	return node1
}

func TestListLength(t *testing.T) {
	node := generateTestNode()
	actual := ListLength(node)
	expected := 5
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}


func TestInsertInLinkedList1(t *testing.T) {
	node := generateTestNode()
	InsertInLinkedList(&node, -1, 1)

	actual := ListLength(node)
	expected := 6
	if actual != expected {
		t.Errorf("length of node: got %v, want %v", actual, expected)
	}

	actual = node.Data
	expected = -1
	if actual != expected {
		t.Errorf("node's data is wrong: got %v, want %v", actual, expected)
	}

	actual = node.Next.Next.Next.Data
	expected = 2
	if actual != expected {
		t.Errorf("node's data is wrong: got %v, want %v", actual, expected)
	}
}

func TestInsertInLinkedList2(t *testing.T) {
	node := generateTestNode()
	InsertInLinkedList(&node, 10, 3)

	actual := ListLength(node)
	expected := 6
	if actual != expected {
		t.Errorf("length of node: got %v, want %v", actual, expected)
	}

	actual = node.Next.Next.Data
	expected = 10
	if actual != expected {
		t.Errorf("node's data is wrong: got %v, want %v", actual, expected)
	}
}

func TestDeleteNodeFromLinkedList(t *testing.T) {
	node := generateTestNode()
	DeleteNodeFromLinkedList(&node, 3)

	actual := ListLength(node)
	expected := 4
	if actual != expected {
		t.Errorf("length of node: got %v, want %v", actual, expected)
	}

	actual = node.Next.Next.Next.Data
	expected = 4
	if actual != expected {
		t.Errorf("node's data is wrong: got %v, want %v", actual, expected)
	}
}


func TestDeleteinkedList(t *testing.T) {
	node := generateTestNode()
	DeleteLinkedList(&node)

	actual := ListLength(node)
	expected := 0
	if actual != expected {
		t.Errorf("length of node: got %v, want %v", actual, expected)
	}

	actual2 := node
	if actual2 != nil {
		t.Errorf("node is wrong: got %v, want %v", actual2, nil)
	}
}
