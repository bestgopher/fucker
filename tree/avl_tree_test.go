package tree

import (
	"testing"
)

func TestAVLTreeSearch(t *testing.T) {
	binaryTree := NewAVLTree(CompareBstTreeInt, 3, 4, 1, 5, 6, 7)
	searchNode(binaryTree, t)
}

func TestAVLTreeDelete(t *testing.T) {
	//binaryTree := NewAVLTree(CompareBstTreeInt, 3, 4, 1, 5, 6, 7)
	//deleteNode(binaryTree, t)
}
