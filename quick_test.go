package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	s := newSort(QuickSort, t)
	s.test()
}
