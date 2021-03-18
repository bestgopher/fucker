package tree

import (
	"fmt"
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
	compare CompareFunc
}

func NewAVLTree(compare CompareFunc, values ...int) *AVLTree {
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
	fmt.Println(value, "insert")
	// 插入值
	a.insert(a.root, node)
	// 整理高度
	a.makeHeight(node)
}

// 插入节点到树中
func (a *AVLTree) insert(node *avlTreeNode, valueNode *avlTreeNode) {
	// 寻找要插入的位置

	for {
		// 循环比较节点，找到插入的地方
		if c := a.compare(valueNode, node); c == 0 {
			return
		} else if c < 0 {
			if node.left != nil {
				node = node.left
			} else {
				valueNode.parent = node
				node.left = valueNode
				break
			}
		} else {
			if node.right != nil {
				node = node.right
			} else {
				valueNode.parent = node
				node.right = valueNode
				break
			}
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

		if c := a.compare(node, parent); c > 0 {
			flag += "R"
		} else {
			flag += "L"
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
	panic("implement me")
}

// 查找元素
func (a *AVLTree) Search(value interface{}) bool {
	node := a.root
	val := &avlTreeNode{value: value}

	for node != nil {
		if c := a.compare(val, node); c == 0 {
			return true
		} else if c < 0 {
			node = node.left
		} else {
			node = node.right
		}
	}

	return false
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
		if c := a.compare(leftSon, parent); c > 0 {
			parent.right = leftSon
		} else {
			parent.left = leftSon
		}
		leftSon.parent = parent
	} else {
		a.root = leftSon
	}
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
		if c := a.compare(rightSon, parent); c > 0 {
			parent.right = rightSon
		} else {
			parent.left = rightSon
		}
		rightSon.parent = parent
	} else {
		a.root = rightSon
	}
}

func (a *AVLTree) lrRotate(node *avlTreeNode) {
	leftSon := node.left                       // 左子节点
	rightGrandson := leftSon.right             // 右孙子节点
	rightGrandsonLeftSon := rightGrandson.left // 右孙子节点的左节点

	node.left = rightGrandson
	rightGrandson.parent = node

	rightGrandson.left = leftSon
	leftSon.parent = rightGrandson

	leftSon.left = rightGrandsonLeftSon
	if rightGrandsonLeftSon != nil {
		rightGrandsonLeftSon.parent = leftSon
	}

	// 右旋
	a.rrRotate(node)
}

// RL
func (a *AVLTree) rlRotate(node *avlTreeNode) {
	rightSon := node.right                     // 右子节点
	leftGrandson := rightSon.left              // 左孙子节点
	leftGrandsonRightSon := leftGrandson.right // 左孙子节点的右节点

	node.right = leftGrandson
	leftGrandson.parent = node

	leftGrandson.right = rightSon
	rightSon.parent = leftGrandson

	rightSon.left = leftGrandsonRightSon
	if leftGrandsonRightSon != nil {
		leftGrandsonRightSon.parent = rightSon
	}

	// 左旋
	a.llRotate(node)
}

// 返回树子节点的最大高度
func (a *AVLTree) maxHeight(tree *avlTreeNode) int {
	tree1, tree2 := tree.left, tree.right
	if tree1 != nil && tree2 != nil {
		if tree1.height > tree2.height {
			return tree1.height
		} else {
			return tree2.height
		}
	} else if tree1 != nil && tree2 == nil {
		return tree1.height
	} else if tree1 == nil && tree2 != nil {
		return tree2.height
	} else {
		return 0
	}
}

// 整理高度
// 后序遍历
func (a *AVLTree) makeHeight(tree *avlTreeNode) {
	parent := tree

	for parent != nil {
		parent.height = a.maxHeight(parent) + 1
		fmt.Println(parent.value, "make height")
		parent = parent.parent
	}
}
