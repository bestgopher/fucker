package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func searchNode(tree Tree, t *testing.T) {
	assert.Equal(t, tree.Search(3), true)
	assert.Equal(t, tree.Search(4), true)
	assert.Equal(t, tree.Search(7), true)
	assert.Equal(t, tree.Search(2), false)
	assert.Equal(t, tree.Search(8), false)
	assert.Equal(t, tree.Search(12), false)
}

func deleteNode(tree Tree, t *testing.T) {
	assert.Equal(t, tree.Search(3), true)
	assert.Equal(t, tree.Search(4), true)
	assert.Equal(t, tree.Search(7), true)
	assert.Equal(t, tree.Search(2), false)
	assert.Equal(t, tree.Search(8), false)
	assert.Equal(t, tree.Search(12), false)
	tree.Delete(3)
	assert.Equal(t, tree.Search(3), false)
	assert.Equal(t, tree.Search(4), true)
	assert.Equal(t, tree.Search(7), true)
	assert.Equal(t, tree.Search(2), false)
	assert.Equal(t, tree.Search(8), false)
	assert.Equal(t, tree.Search(12), false)
	tree.Delete(6)
	assert.Equal(t, tree.Search(3), false)
	assert.Equal(t, tree.Search(4), true)
	assert.Equal(t, tree.Search(7), true)
	assert.Equal(t, tree.Search(6), false)
	assert.Equal(t, tree.Search(2), false)
	assert.Equal(t, tree.Search(8), false)
	assert.Equal(t, tree.Search(12), false)
}
