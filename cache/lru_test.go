package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLruCache_Set(t *testing.T) {
	var l, _ = NewLRUCache(10)
	for i := 1; i < 11; i++ {
		l.Set(i, i)
	}

	for i := 1; i < 11; i++ {
		_, ok := l.Get(i)
		assert.Equal(t, true, ok)
	}

	assert.Equal(t, 10, l.Len())
}

func TestLruCache_Get(t *testing.T) {
	var l, _ = NewLRUCache(10)
	for i := 1; i < 11; i++ {
		l.Set(i, i)
	}

	l.Set(20, 20)
	_, ok := l.Get(1)
	assert.Equal(t, false, ok)

	l.Set(21, 21)
	_, ok = l.Get(2)
	assert.Equal(t, false, ok)

	l.Get(3)
	l.Set(22, 22)
	_, ok = l.Get(3)
	assert.Equal(t, true, ok)

	_, ok = l.Get(4)
	assert.Equal(t, false, ok)
}

func TestLruCache_Delete(t *testing.T) {
	var l, _ = NewLRUCache(10)
	for i := 1; i < 11; i++ {
		l.Set(i, i)
	}

	l.Delete(5)
	_, ok := l.Get(5)
	assert.Equal(t, false, ok)
	assert.Equal(t, 9, l.Len())
}
