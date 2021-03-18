package tree

import (
	"testing"
)

func TestBinarySearchTreeSearch(t *testing.T) {
	binaryTree := NewBinarySearchTree(CompareBstTreeInt, 3, 4, 1, 5, 6, 7)
	searchNode(binaryTree, t)
}

func TestBinarySearchTreeDelete(t *testing.T) {
	binaryTree := NewBinarySearchTree(CompareBstTreeInt, 3, 4, 1, 5, 6, 7)
	deleteNode(binaryTree, t)
}
