package dll_node

import (
	"testing"
)

func generateTestDLLNode() *DLLNode {
	node5 := &DLLNode{Data: 5, Next: nil, Prev: nil}
	node4 := &DLLNode{Data: 4, Next: node5, Prev: nil}
	node3 := &DLLNode{Data: 3, Next: node4, Prev: nil}
	node2 := &DLLNode{Data: 2, Next: node3, Prev: nil}
	node1 := &DLLNode{Data: 1, Next: node2, Prev: nil}

	node2.Prev = node1
	node3.Prev = node2
	node4.Prev = node3
	node5.Prev = node4

	return node1
}

func TestDLLInsert1(t *testing.T) {
	node := generateTestDLLNode()
	DLLInsert(&node, 0, 3)

	actual := DLLLength(node)
	expected := 6
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}

	actual = node.Next.Next.Data
	expected = 0
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}

	actual = node.Next.Next.Next.Prev.Prev.Data
	expected = 2
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}

	actual = node.Next.Next.Next.Next.Next.Data
	expected = 5
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestDLLInsert2(t *testing.T) {
	node := generateTestDLLNode()
	DLLInsert(&node, 0, 1)

	actual := DLLLength(node)
	expected := 6
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}

	actual = node.Data
	expected = 0
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestDLLDelete1(t *testing.T) {
	node := generateTestDLLNode()
	DLLDelete(&node, 3)

	actual := DLLLength(node)
	expected := 4
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}

	actual = node.Next.Next.Data
	expected = 4
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}

	actual = node.Next.Next.Next.Prev.Prev.Data
	expected = 2
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}


func TestDLLDelete2(t *testing.T) {
	node := generateTestDLLNode()
	DLLDelete(&node, 1)

	actual := DLLLength(node)
	expected := 4
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}

	actual = node.Data
	expected = 2
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}

	actual = node.Next.Next.Next.Prev.Prev.Data
	expected = 3
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}
