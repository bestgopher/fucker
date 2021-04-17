package cache

import (
	"errors"
)

// listNode 链表的节点
type listNode struct {
	key, value interface{}
	next, prev *listNode
}

// 双端链表的简单实现
type list struct {
	head, tail *listNode
}

// LRU, 最近最少使用，当数据满了，再插入时，淘汰掉最近最少使用的元素
// 基于链表和哈希表实现
// 使用双链表存储元素，链表尾部的元素是最近使用的元素
// 当元素个数已满时，淘汰掉链表头部的元素
type lruCache struct {
	// 长度
	len int
	// 容量
	cap int
	// 链表
	linkList *list
	// 哈希表
	hash map[interface{}]*listNode
}

// NewLRUCache 返回一个指定容量的LRUCache表
func NewLRUCache(cap int) (Cache, error) {
	if cap <= 0 {
		return nil, errors.New("cap cannot be negative or zero")
	}
	return &lruCache{cap: cap, hash: make(map[interface{}]*listNode, cap), linkList: &list{}}, nil
}

// Get 获取元素，并且把元素要放在链表的末尾
func (l *lruCache) Get(key interface{}) (interface{}, bool) {
	node, ok := l.hash[key]
	if !ok {
		return nil, false
	}

	// 最后返回元素
	return node.value, true
}

// Set 设置元素。把元素放到链表尾部
// 当元素满时，要插入新的元素时，需要淘汰掉链表头部的元素
func (l *lruCache) Set(key, value interface{}) {
	l.Delete(key)
	node := &listNode{key: key, value: value}
	if l.cap == l.len { // 当元素满了的时候，先把头部节点删除，然后把当前节点插入尾部
		l.Delete(l.linkList.head.key)
	}

	l.insert(node)
	l.hash[key] = node
	l.len++
}

// Delete 删除元素
func (l *lruCache) Delete(key interface{}) {
	node, ok := l.hash[key]
	if !ok {
		return
	}
	// 先删除hash表的数据
	delete(l.hash, key)
	// 再把节点从链表中删除
	prev := node.prev
	next := node.next
	if prev != nil {
		prev.next = next
	} else {
		l.linkList.head = next
	}

	if next != nil {
		next.prev = prev
	} else {
		l.linkList.tail = next
	}
	l.len--
}

// Len 返回LRUCache的元素个数
func (l *lruCache) Len() int {
	return l.len
}

// 改变元素的生命周期
func (l *lruCache) changeTTL(node *listNode) {
	// 1.首先去掉此元素
	prev, next := node.prev, node.next
	if prev != nil {
		prev.next = next
	} else {
		l.linkList.head = next
	}

	if next != nil {
		next.prev = prev
	}

	// 2.把此元素放在链表尾部
	l.linkList.tail.next = node
	node.prev = l.linkList.tail.next
	l.linkList.tail = node
}

// 向链表中插入元素
func (l *lruCache) insert(node *listNode) {
	if l.linkList.head == nil {
		l.linkList.tail = node
		l.linkList.head = node
	} else {
		l.linkList.tail.next = node
		node.prev = l.linkList.tail
		l.linkList.tail = node
	}
}
