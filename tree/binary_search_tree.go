package tree

// 二叉查找树
type BinarySearchTree struct {
	Root *Node
}

// 节点
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewBinarySearchTree(values ...int) *BinarySearchTree {
	t := new(BinarySearchTree)
	for _, v := range values {
		t.Insert(v)
	}

	return t
}

func (b *BinarySearchTree) Insert(value int) {
	if b.Root == nil {
		b.Root = &Node{Value: value}
		return
	}

	node := b.Root

	for {
		if value < node.Value {
			if node.Left == nil {
				node.Left = &Node{Value: value}
				break
			} else {
				node = node.Left
			}
		} else {
			if node.Right == nil {
				node.Right = &Node{Value: value}
				break
			} else {
				node = node.Right
			}
		}
	}
}

func (b *BinarySearchTree) Search(value int) bool {
	node := b.Root

	for node != nil {

		if value == node.Value {
			return true
		} else if value < node.Value {
			node = node.Left
		} else {
			node = node.Right

		}
	}

	return false
}

func (b *BinarySearchTree) Delete(value int) {
	b.Root = b.delete(b.Root, value)
}

func (b *BinarySearchTree) delete(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if node.Value == value {
		if node.Left == nil && node.Right == nil { // 左右子节点都为空时
			node = nil
		} else if node.Left == nil && node.Right != nil { // 左子节点为空，右子节点不为空
			node = node.Right
		} else if node.Right == nil && node.Left != nil { // 右子节点为空，左子节点不为空
			node = node.Left
		} else {
			// 左右子节点都不为空时，获取右子树的最小子节点与当前节点交换
			n1, n2 := node, node.Right
			for n2.Left != nil {
				n1 = n2
				n2 = n2.Left
			}
			node.Value = n2.Value
			n1.Left = n2.Right
		}
	} else if value < node.Value {
		node.Left = b.delete(node.Left, value)
	} else {
		node.Right = b.delete(node.Right, value)
	}

	return node
}
