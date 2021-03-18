package tree

import (
	"fmt"
	"testing"
)

func TestAVLTreeSearch(t *testing.T) {
	binaryTree := NewAVLTree(CompareBstTreeInt, 3, 4, 1, 5, 6)
	fmt.Println(binaryTree.root,  )
	fmt.Println(binaryTree.root.left)
	fmt.Println(binaryTree.root.right)
	fmt.Println(binaryTree.root.right.left)
	fmt.Println(binaryTree.root.right.right)
	searchNode(binaryTree, t)
}

func TestAVLTreeDelete(t *testing.T) {
	//binaryTree := NewAVLTree(CompareBstTreeInt, 3, 4, 1, 5, 6, 7)
	//deleteNode(binaryTree, t)
}
