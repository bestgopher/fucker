package cache

type Cache interface {
	// 获取对应的值
	Get(key interface{}) (interface{}, bool)
	// 设置对应的值
	Set(key, value interface{})
	// 删除对应的值
	Delete(key interface{})
	// 返回数据的长度
	Len() int
}
