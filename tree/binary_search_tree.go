package tree

// 节点
type bstTreeNode struct {
	value interface{}
	left  *bstTreeNode
	right *bstTreeNode
}

func (b *bstTreeNode) Value() interface{} { return b.value }

// 二叉查找树
type BinarySearchTree struct {
	root    *bstTreeNode
	compare CompareFunc
}

func NewBinarySearchTree(compare CompareFunc, values ...interface{}) *BinarySearchTree {
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

	for {
		if c := b.compare(r, node); c < 0 {
			if node.left == nil {
				node.left = r
				break
			} else {
				node = node.left
			}
		} else {
			if node.right == nil {
				node.right = r
				break
			} else {
				node = node.right
			}
		}
	}
}

// 搜索节点
func (b *BinarySearchTree) Search(value interface{}) bool {
	node := b.root
	r := &bstTreeNode{value: value}

	for node != nil {
		if c := b.compare(r, node); c == 0 {
			return true
		} else if c < 0 {
			node = node.left
		} else {
			node = node.right
		}
	}

	return false
}

func (b *BinarySearchTree) Delete(value interface{}) {
	b.root = b.delete(b.root, value)
}

func (b *BinarySearchTree) delete(node *bstTreeNode, value interface{}) *bstTreeNode {
	if node == nil {
		return nil
	}

	r := &bstTreeNode{value: value}
	// 比较当前节点与待删除节点的值
	if c := b.compare(r, node); c == 0 {
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
	} else if c < 0 {
		node.left = b.delete(node.left, value)
	} else {
		node.right = b.delete(node.right, value)
	}

	return node
}
