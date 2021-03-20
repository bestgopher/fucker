package tree

import (
	"github.com/bestgopher/fucker"
)

// 平衡二叉搜索树的节点
type avlTreeNode struct {
	value  interface{}  // value
	parent *avlTreeNode // parent node
	left   *avlTreeNode // left node
	right  *avlTreeNode // right node
	height int          // height
}

func (a *avlTreeNode) Value() interface{} { return a.value }

/*
平衡二叉查找树
	树的高度：一片叶子的位置高度为1，null孩子的高度为0
	高度平衡属性：对于T中每一个位置p，p的孩子的高度最多相差1。满足高度平衡属性的二叉搜索树被称为AVL树。AVL树的子树也是一颗AVL树。

	为了维护一个位置的子树高度之差绝对值为1。因此每次添加或者删除后都应该进行调整。
	删除和插入的方式和二叉搜索树一致。

插入
	插入一个节点，会在叶子节点p的位置上产生一个新的节点。这个操作可能违反了高度平衡属性。且唯一可能变得不平衡的位置是p的祖先。
	对于节点A，四种方式插入会破坏A的平衡性
	LL：在A的左子树根节点的左子树上插入节点
	RR：在A的右子树根节点的右子树上插入节点
	LR：在A的左子树根节点的右子树上插入节点
	RL：在A的右子树根节点的左子树上插入节点
*/
type AVLTree struct {
	root    *avlTreeNode
	compare fucker.CompareFunc
}

func NewAVLTree(compare fucker.CompareFunc, values ...int) *AVLTree {
	tree := &AVLTree{compare: compare}

	for _, v := range values {
		tree.Insert(v)
	}

	return tree
}

// Insert 插入一个值到树中
func (a *AVLTree) Insert(value interface{}) {
	node := &avlTreeNode{value: value, height: 1}
	if a.root == nil {
		a.root = node
		return
	}
	// 插入值
	a.insert(a.root, node)
}

// 插入节点到树中
func (a *AVLTree) insert(node *avlTreeNode, valueNode *avlTreeNode) {
	// 寻找要插入的位置

LOOP:
	for {
		// 循环比较节点，找到插入的地方
		switch a.compare(valueNode, node) {
		case fucker.Less:
			if node.left != nil {
				node = node.left
			} else {
				valueNode.parent = node
				node.left = valueNode
				break LOOP
			}
		case fucker.Greater:
			if node.right != nil {
				node = node.right
			} else {
				valueNode.parent = node
				node.right = valueNode
				break LOOP
			}
		default:
			return
		}
	}

	a.rotate(valueNode)
}

// 旋转
// 从插入节点向上寻找，找到最近不平衡节点
func (a *AVLTree) rotate(node *avlTreeNode) {
	parent := node.parent
	flag := "" // 用于标示是LL\RR\LR\RL

	for {
		if parent == nil {
			return
		}
		switch a.compare(node, parent) {
		case fucker.Greater:

			flag += "R"
		case fucker.Less:
			flag += "L"
		default:
			return
		}

		parent.height = a.maxHeight(parent) + 1

		if !a.isBalance(parent) {
			break
		}

		node = parent
		parent = node.parent
	}

	switch flag[len(flag)-2:] {
	case "LL":
		a.llRotate(parent)
	case "RR":
		a.rrRotate(parent)
	case "LR":
		a.lrRotate(parent)
	case "RL":
		a.rlRotate(parent)
	}
}

// 删除元素
func (a *AVLTree) Delete(value interface{}) {
	n := &avlTreeNode{value: value}
	node := a.root
	isLeft := false // 是父节点的左节点的时候为true，反之为false

LOOP:
	for node != nil {
		switch a.compare(n, node) {
		case fucker.Equal:
			parent := node.parent

			if node.left != nil && node.right != nil {
				// 当node节点的左右节点都不为空时，获取node节点右节点的最小节点，与node交换
				smallestRightNode := node.right
				if smallestRightNode.left != nil {
					smallestRightNode = smallestRightNode.left
				}

				node.value = smallestRightNode.value
				if smallestRightNode.right != nil {
					smallestRightNode.value = smallestRightNode.right.value
					smallestRightNode.right.parent = nil
					smallestRightNode.right = nil
				}

			} else if node.left != nil && node.right == nil {
				// 当只有左节点不为空时，直接左节点替换当前节点
				if isLeft {
					node.parent.left = node.left
				} else {
					node.parent.right = node.left
				}
				node.left.parent = parent

			} else if node.left == nil && node.right != nil {
				// 当只有右节点不为空时，直接右节点替换当前节点
				if isLeft {
					node.parent.left = node.right
				} else {
					node.parent.right = node.right
				}
				node.right.parent = parent

			} else {
				// 当左右节点都为空时，直接移除当前节点
				if isLeft {
					node.parent.left = nil
				} else {
					node.parent.right = nil
				}
			}
			// 从父节点开始重新整理高度
			if parent == nil {
				a.makeHeight(a.root)
			} else {
				a.makeHeight(parent)
			}
			break LOOP

		case fucker.Less:
			node = node.left
			isLeft = true

		case fucker.Greater:
			node = node.right
			isLeft = false

		default:
			break LOOP
		}
	}
}

// 查找元素
func (a *AVLTree) Search(value interface{}) Value {
	node := a.root
	val := &avlTreeNode{value: value}

	for node != nil {
		switch a.compare(val, node) {
		case fucker.Less:
			node = node.left
		case fucker.Greater:
			node = node.right
		default:
			return node
		}
	}

	return nil
}

// isBalance 判断节点是否平衡
func (a *AVLTree) isBalance(node *avlTreeNode) bool {
	// 当节点为空时，肯定是平衡的
	if node == nil {
		return true
	}

	left := 0
	if node.left != nil {
		left = node.left.height
	}

	right := 0
	if node.right != nil {
		right = node.right.height
	}

	h := left - right
	return h >= -1 && h <= 1
}

// LL
func (a *AVLTree) llRotate(node *avlTreeNode) {
	parent := node.parent          // 父节点
	leftSon := node.left           // 左子节点
	rightGrandson := leftSon.right // 右孙子节点

	// node节点放在左儿子节点的右子节点上
	leftSon.right = node
	node.parent = leftSon

	// 原右孙子节点放在node的左节点上
	node.left = rightGrandson
	if rightGrandson != nil {
		rightGrandson.parent = node

	}

	// 原左子节点放在原父节点上
	if parent != nil {
		switch a.compare(leftSon, parent) {
		case fucker.Greater:
			parent.right = leftSon
		case fucker.Less:
			parent.left = leftSon
		}
		leftSon.parent = parent
	} else {
		a.root = leftSon
		leftSon.parent = nil
	}

	a.makeHeight(leftSon)
}

// rr
func (a *AVLTree) rrRotate(node *avlTreeNode) {
	parent := node.parent         // 父节点
	rightSon := node.right        // 右子节点
	leftGrandson := rightSon.left // 左孙子节点

	// node节点放在左儿子节点的右子节点上
	rightSon.left = node
	node.parent = rightSon

	// 原右孙子节点放在node的左节点上
	node.right = leftGrandson
	if leftGrandson != nil {
		leftGrandson.parent = node
	}

	// 原左子节点放在原父节点上
	if parent != nil {
		switch a.compare(rightSon, parent) {
		case fucker.Greater:
			parent.right = rightSon
		case fucker.Less:
			parent.left = rightSon
		}
		rightSon.parent = parent
	} else {
		a.root = rightSon
		rightSon.parent = nil
	}

	a.makeHeight(rightSon)
}

// lr
func (a *AVLTree) lrRotate(node *avlTreeNode) {
	rightSon := node.right                           // node节点的右子节点
	rightSonLeftSon := rightSon.left                 // node节点的右子节点的左子节点
	rightSonLeftSonRightSon := rightSonLeftSon.right // node节点的右子节点的左子节点的右子节点

	node.right = rightSonLeftSon
	rightSonLeftSon.parent = node

	rightSonLeftSon.right = rightSon
	rightSon.parent = rightSonLeftSon

	rightSon.left = rightSonLeftSonRightSon
	if rightSonLeftSonRightSon != nil {
		rightSonLeftSonRightSon.parent = rightSon
	}
	// 右旋
	a.rrRotate(node)
}

// RL
func (a *AVLTree) rlRotate(node *avlTreeNode) {
	leftSon := node.left                           // node的左子节点
	leftSonRightSon := leftSon.right               // node的左子节点的右子节点
	leftSonRightSonLeftSon := leftSonRightSon.left // node的左子节点的右子节点的左子节点

	node.left = leftSonRightSon
	leftSonRightSon.parent = node

	leftSonRightSon.left = leftSon
	leftSon.parent = leftSonRightSon

	leftSon.right = leftSonRightSonLeftSon
	if leftSonRightSonLeftSon != nil {
		leftSonRightSonLeftSon.parent = leftSon
	}
	// 左旋
	a.llRotate(node)
}

// 返回树子节点的最大高度
func (a *AVLTree) maxHeight(node *avlTreeNode) int {
	leftNode, rightNode := node.left, node.right
	if leftNode != nil && rightNode != nil {
		if leftNode.height > rightNode.height {
			return leftNode.height
		} else {
			return rightNode.height
		}
	} else if leftNode != nil && rightNode == nil {
		return leftNode.height
	} else if leftNode == nil && rightNode != nil {
		return rightNode.height
	} else {
		return 0
	}
}

// 整理高度
// 先向下整理，后向上整理
func (a *AVLTree) makeHeight(node *avlTreeNode) {
	a.downMakeHeight(node)
	a.upMakeHeight(node)
}

// 向下整理高度
func (a *AVLTree) downMakeHeight(node *avlTreeNode) {
	if node == nil {
		return
	}

	a.downMakeHeight(node.left)
	a.downMakeHeight(node.right)
	node.height = a.maxHeight(node) + 1
}

// 向上整理高度
func (a *AVLTree) upMakeHeight(node *avlTreeNode) {
	for node != nil {
		node.height = a.maxHeight(node) + 1
		node = node.parent
	}
}
