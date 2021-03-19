package fucker

import (
	"math/rand"
	"testing"
	"time"

	"github.com/bestgopher/fucker"
)

type sort struct{}

func newSort() *sort {
	return &sort{}
}

func (s *sort) sort(f func([]interface{}, fucker.CompareFunc), compare fucker.CompareFunc) bool {
	data := s.generate()
	for _, v := range data {
		f(v, compare)
		if !s.check(v, compare) {
			return false
		}
	}
	return true
}

// 随机生成乱序的一些切片
func (s *sort) generate() [][]interface{} {
	rand.Seed(time.Now().UnixNano())
	var l = rand.Intn(50) + 50
	var result = make([][]interface{}, 0, l)
	for i := 0; i < l; i++ {
		length := rand.Intn(100) + 50
		s := make([]interface{}, 0, length)
		for j := 0; j < length; j++ {
			s = append(s, rand.Intn(10000))
		}
		result = append(result, s)
	}

	// 添加三个特殊情况
	result = append(result, []interface{}{}, nil, []interface{}{1})

	return result
}

// 复制一份生成的数据
func (s *sort) cloneData(datas [][]int) [][]int {
	v := make([][]int, 0, len(datas))

	for _, data := range datas {
		d := make([]int, len(data))
		copy(d, data)
		v = append(v, d)
	}
	return v
}

// 检查切片是否是升序
// 是升序返回true，否则返回false
func (s *sort) check(data []interface{}, compare fucker.CompareFunc) bool {
	if len(data) < 2 {
		return true
	}

	for i := 1; i < len(data); i++ {
		if compare(data[i], data[i-1]) == fucker.Less {
			return false
		}
	}

	return true
}

var functions = map[string]func([]interface{}, fucker.CompareFunc){
	"Selection Sort": SelectionSort,
	"Merge Sort":     MergeSort,
	"Quick Sort":     QuickSort,
	"Bubble Sort":    BubbleSort,
	"Heap Sort":      HeapSort,
	"Insertion Sort": InsertionSort,
	"Shell Sort":     Shell,
}

func TestSort(t *testing.T) {
	for name, f := range functions {
		s := newSort()
		if s.sort(f, compareFunc) {
			t.Logf("%s success", name)
		} else {
			t.Fatalf("%s failed", name)
		}
	}
}

func compareFunc(i, j interface{}) fucker.Compare {
	a := i.(int)
	b := j.(int)

	if a > b {
		return fucker.Greater
	} else if a == b {
		return fucker.Equal
	} else {
		return fucker.Less
	}
}
