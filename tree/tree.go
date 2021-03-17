package tree

// 用于比较节点的方法
// v1 小于 v2 返回值 小于0
// v1 等于 v2 返回值 等于0
// v1 大于 v2 返回值 大于0
type CompareFunc func(v1 Value, v2 Value) int

// 获取节点值的接口
type Value interface {
	Value() interface{}
}
