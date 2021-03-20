package fucker

// 两个元素比较的结果
type Compare uint8

const (
	// 相等
	Equal Compare = iota + 1
	// 小于
	Less
	// 大于
	Greater
)

// 用于比较节点的方法
type CompareFunc func(v1 interface{}, v2 interface{}) Compare
