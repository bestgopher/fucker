package sort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type sort struct {
	handler func([]int)
	t       *testing.T
}

func newSort(handler func([]int)) *sort {
	return &sort{handler: handler}
}

func (s *sort) test() bool {
	data := s.generate()

	for _, v := range data {
		s.handler(v)
		if !s.check(v) {
			return false
		}
	}

	return true
}

// 随机生成乱序的一些切片
func (s *sort) generate() [][]int {
	rand.Seed(time.Now().UnixNano())
	var l = rand.Intn(50) + 50
	var result = make([][]int, 0, l)
	for i := 0; i < l; i++ {
		length := rand.Intn(100) + 50
		s := make([]int, 0, length)
		for j := 0; j < length; j++ {
			s = append(s, rand.Intn(10000))
		}
		result = append(result, s)
	}

	// 添加三个特殊情况
	result = append(result, []int{}, nil, []int{1})

	return result
}

// 检查切片是否是升序
// 是升序返回true，否则返回false
func (s *sort) check(data []int) bool {
	if len(data) < 2 {
		return true
	}

	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}

	return true
}

func TestSort(t *testing.T) {
	funcs := map[string]func([]int){
		"Selection Sort": SelectionSort,
		"Merge Sort":     MergeSort,
		"Quick Sort":     QuickSort,
		"Bubble Sort":    BubbleSort,
	}

	for name, f := range funcs {
		s := newSort(f)
		if assert.Equalf(t, s.test(), true, "%s failed", name) {
			t.Logf("%s success", name)
		}
	}
}
