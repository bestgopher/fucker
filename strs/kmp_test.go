package strs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKMP(t *testing.T) {
	assert.Equal(t, 2, KMP("hello", "ll"))
	assert.Equal(t, -1, KMP("aaaaa", "bba"))
	assert.Equal(t, 6, KMP("ababcaababcaabc", "ababcaabc"))
	assert.Equal(t, 0, KMP("", ""))
	assert.Equal(t, 4, KMP("aabaaabaaac", "aabaaac"))
}
