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
