package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestSearch(t *testing.T) {

	binaryTree := NewBinarySearchTree(CompareBstTreeInt, 3, 4, 1, 5, 6, 7)
	assert.Equal(t, binaryTree.Search(3), true)
	assert.Equal(t, binaryTree.Search(4), true)
	assert.Equal(t, binaryTree.Search(7), true)
	assert.Equal(t, binaryTree.Search(2), false)
	assert.Equal(t, binaryTree.Search(8), false)
	assert.Equal(t, binaryTree.Search(12), false)
}

func TestDelete(t *testing.T) {
	binaryTree := NewBinarySearchTree(CompareBstTreeInt, 3, 4, 1, 5, 6, 7)
	assert.Equal(t, binaryTree.Search(3), true)
	assert.Equal(t, binaryTree.Search(4), true)
	assert.Equal(t, binaryTree.Search(7), true)
	assert.Equal(t, binaryTree.Search(2), false)
	assert.Equal(t, binaryTree.Search(8), false)
	assert.Equal(t, binaryTree.Search(12), false)
	binaryTree.Delete(3)
	assert.Equal(t, binaryTree.Search(3), false)
	assert.Equal(t, binaryTree.Search(4), true)
	assert.Equal(t, binaryTree.Search(7), true)
	assert.Equal(t, binaryTree.Search(2), false)
	assert.Equal(t, binaryTree.Search(8), false)
	assert.Equal(t, binaryTree.Search(12), false)
	binaryTree.Delete(6)
	assert.Equal(t, binaryTree.Search(3), false)
	assert.Equal(t, binaryTree.Search(4), true)
	assert.Equal(t, binaryTree.Search(7), true)
	assert.Equal(t, binaryTree.Search(6), false)
	assert.Equal(t, binaryTree.Search(2), false)
	assert.Equal(t, binaryTree.Search(8), false)
	assert.Equal(t, binaryTree.Search(12), false)
}
