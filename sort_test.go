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

func newSort(handler func([]int), t *testing.T) *sort {
	return &sort{handler: handler, t: t}
}

func (s *sort) test() {
	data := s.generate()

	for _, v := range data {
		s.t.Logf("排序前： %v\n", v)
		s.handler(v)
		s.t.Logf("排序后： %v\n", v)
		assert.Equal(s.t, s.check(v), true)
	}
}

// 随机生成乱序的一些切片
func (s *sort) generate() [][]int {
	rand.Seed(time.Now().UnixNano())
	var l = rand.Intn(50) + 50
	var result = make([][]int, 0, l)
	for i := 0; i < l; i++ {
		length := rand.Intn(100)
		s := make([]int, 0, length)
		for j := 0; j < length; j++ {
			s = append(s, rand.Intn(10000))
		}
		result = append(result, s)
	}

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
