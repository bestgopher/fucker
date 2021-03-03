package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	s := newSort(BubbleSort, t)
	s.test()
}
