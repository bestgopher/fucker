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

func searchNode(tree Tree, t *testing.T) {
	assert.NotNil(t, tree.Search(3))
	assert.Equal(t, tree.Search(3).Value(), 3)
	assert.NotNil(t, tree.Search(4))
	assert.Equal(t, tree.Search(4).Value(), 4)
	assert.NotNil(t, tree.Search(7))
	assert.Equal(t, tree.Search(7).Value(), 7)
	assert.Nil(t, tree.Search(2))
	assert.Nil(t, tree.Search(8))
	assert.Nil(t, tree.Search(12))
}

func deleteNode(tree Tree, t *testing.T) {
	assert.NotNil(t, tree.Search(3))
	assert.NotNil(t, tree.Search(4))
	assert.NotNil(t, tree.Search(7))
	assert.Nil(t, tree.Search(2))
	assert.Nil(t, tree.Search(8))
	assert.Nil(t, tree.Search(12))
	tree.Delete(3)
	assert.Nil(t, tree.Search(3))
	assert.NotNil(t, tree.Search(4))
	assert.NotNil(t, tree.Search(7))
	assert.Nil(t, tree.Search(2))
	assert.Nil(t, tree.Search(8))
	assert.Nil(t, tree.Search(12))
	tree.Delete(6)
	assert.Nil(t, tree.Search(3))
	assert.NotNil(t, tree.Search(4))
	assert.NotNil(t, tree.Search(7))
	assert.Nil(t, tree.Search(6))
	assert.Nil(t, tree.Search(2))
	assert.Nil(t, tree.Search(8))
	assert.Nil(t, tree.Search(12))
}
