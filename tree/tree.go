package tree

// 获取节点值的接口
type Value interface {
	Value() interface{}
}

// Any tree must be implemented the Tree interface
type Tree interface {
	Insert(value interface{})       // insert a value to the tree
	Delete(value interface{})       // delete a value from the tree
	Search(value interface{}) Value // search for a value from the tree
}
