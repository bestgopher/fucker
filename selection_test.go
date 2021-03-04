package sort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	s := newSort(SelectionSort, t)
	s.test()
}
