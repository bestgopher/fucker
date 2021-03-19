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
// v1 小于 v2 返回值 小于0
// v1 等于 v2 返回值 等于0
// v1 大于 v2 返回值 大于0
type CompareFunc func(v1 interface{}, v2 interface{}) Compare
