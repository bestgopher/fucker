package tree

import (
	"gopkg.in/bestgopher/fucker"
)

// BST节点
type bstTreeNode struct {
	value interface{}
	left  *bstTreeNode
	right *bstTreeNode
}

func (b *bstTreeNode) Value() interface{} { return b.value }

// 二叉查找树
type BinarySearchTree struct {
	root    *bstTreeNode
	compare fucker.CompareFunc
}

func NewBinarySearchTree(compare fucker.CompareFunc, values ...interface{}) *BinarySearchTree {
	t := &BinarySearchTree{compare: compare}
	for _, v := range values {
		t.Insert(v)
	}

	return t
}

// 插入节点
func (b *BinarySearchTree) Insert(value interface{}) {
	if b.root == nil {
		b.root = &bstTreeNode{value: value}
		return
	}

	node := b.root
	r := &bstTreeNode{value: value}

LOOP:
	for {
		switch b.compare(r, node) {
		case fucker.Equal:
			node.value = r.value
			return
		case fucker.Less:
			if node.left == nil {
				node.left = r
				break LOOP
			} else {
				node = node.left
			}
		case fucker.Greater:
			if node.right == nil {
				node.right = r
				break LOOP
			} else {
				node = node.right
			}
		default:
			break LOOP
		}
	}
}

// 搜索节点
func (b *BinarySearchTree) Search(value interface{}) Value {
	node := b.root
	r := &bstTreeNode{value: value}

	for node != nil {
		switch b.compare(r, node) {
		case fucker.Equal:
			return node
		case fucker.Less:
			node = node.left
		case fucker.Greater:
			node = node.right
		}
	}

	return nil
}

// 删除节点
// 当被删除节点为叶子节点时(即没有子节点)，直接删除此节点
// 当被删除节点只有一个子节点时，删除此节点，然后子节点替换到此节点的位置
// 当被删除节点有两个子节点时，删除此节点，使用此节点的后继(前驱)节点替换此节点位置(后继:此节点右子树中的最小节点，前驱：此节点左子树中的最大节点)
func (b *BinarySearchTree) Delete(value interface{}) {
	b.root = b.delete(b.root, value)
}

func (b *BinarySearchTree) delete(node *bstTreeNode, value interface{}) *bstTreeNode {
	if node == nil {
		return nil
	}

	r := &bstTreeNode{value: value}
	// 比较当前节点与待删除节点的值
	switch b.compare(r, node) {
	case fucker.Equal:
		if node.left == nil && node.right == nil { // 左右子节点都为空时
			node = nil
		} else if node.left == nil && node.right != nil { // 左子节点为空，右子节点不为空
			node = node.right
		} else if node.right == nil && node.left != nil { // 右子节点为空，左子节点不为空
			node = node.left
		} else {
			// 左右子节点都不为空时，获取右子树的最小子节点与当前节点交换
			n1, n2 := node, node.right
			for n2.left != nil {
				n1, n2 = n2, n2.left
			}
			node.value, n1.left = n2.value, n2.right
		}
	case fucker.Less:
		node.left = b.delete(node.left, value)
	case fucker.Greater:
		node.right = b.delete(node.right, value)
	}

	return node
}
