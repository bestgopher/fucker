package tree

import (
	"testing"
)

func CompareBstTreeInt(a Value, b Value) int {
	v1 := a.Value().(int)
	v2 := b.Value().(int)

	if v1 == v2 {
		return 0
	} else if v1 < v2 {
		return -1
	} else {
		return 1
	}
}

func TestBinarySearchTreeSearch(t *testing.T) {
	binaryTree := NewBinarySearchTree(CompareBstTreeInt, 3, 4, 1, 5, 6, 7)
	searchNode(binaryTree, t)
}

func TestBinarySearchTreeDelete(t *testing.T) {
	binaryTree := NewBinarySearchTree(CompareBstTreeInt, 3, 4, 1, 5, 6, 7)
	deleteNode(binaryTree, t)
}
