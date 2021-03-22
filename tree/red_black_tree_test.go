package tree

import (
	"testing"
)

func TestRedBlackTreeSearch(t *testing.T) {
	rbTree := NewRedBlackTree(CompareBstTreeInt, 3, 4, 1, 5, 6, 7)
	searchNode(rbTree, t)
}

func TestRedBlackTreeDelete(t *testing.T) {
	rbTree := NewRedBlackTree(CompareBstTreeInt, 3, 4, 1, 5, 6, 7)
	deleteNode(rbTree, t)
}
