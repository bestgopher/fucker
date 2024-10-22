package sort

import (
	"gopkg.in/bestgopher/fucker.v1/sort/internal"
)

var (
	BubbleSort    = internal.BubbleSort
	SelectionSort = internal.SelectionSort
	MergeSort     = internal.MergeSort
	QuickSort     = internal.QuickSort
	InsertionSort = internal.InsertionSort
	HeapSort      = internal.HeapSort
	ShellSort     = internal.Shell
	CountingSort  = internal.CountingSort
)
