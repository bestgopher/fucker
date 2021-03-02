package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	s := newSort(MergeSort, t)
	s.test()
}
