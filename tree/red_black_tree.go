package tree

import (
	"github.com/bestgopher/fucker"
)

// 红黑树的颜色
type _color uint8

const (
	// 红色
	red _color = iota + 1
	// 黑色
	black
)

// 红黑树的节点类
type redBlackTreeNode struct {
	// 节点值
	value interface{}
	// 节点颜色
	color _color
	// 左子节点
	left *redBlackTreeNode
	// 右子节点
	right *redBlackTreeNode
	// 父节点
	parent *redBlackTreeNode
}

func (r *redBlackTreeNode) Value() interface{} {
	return r.value
}

// 获取节点的颜色，当节点为空时，默认为黑色
func (r *redBlackTreeNode) getColor() _color {
	if r == nil {
		return black
	}
	return r.color
}

/*`
红黑树：
	是平衡树的一种替换方案，它比平衡树效率高，红黑树在最坏情况下的操作时间是O(logN)。(平衡树维护平衡需要很多操作)

红黑树定义:
	1.每个节点都被标记为红色或者黑色
	2.根是黑色的
	3.如果一个节点是红色的，那么其儿子节点都是黑色的(两个红色节点不能相连)。
	4.从任一节点到空节点的路径必须具有相同数量的黑色节点，俗称 黑高
		- 如果一个节点存在黑子节点，那么该节点肯定有两个子节点
	5.每个叶子节点(NIL)是黑色

红黑树性质:
	1.根节点到空节点的任一路径上，设含黑节点个数为H
		- 全是黑节点 -> 路径最佳(H)
		- 黑节点总在红节点后(除根) -> 路径最长(2H)
	2.设红黑树高度为h，节点数为n -> h < 2H
		- 第1层到第H层不可能有空节点(否则与H个黑节点矛盾)
		- 第1层到第H层总节点数2^H-1 (n > 2^H-1)

插入:
	1.插入过程同普通的二叉查找树，只是插入后被插入的节点要被着色
	2.着成黑色是不可能的，会违反定义4，必须着成红色
	3.如父节点是黑色的，插入结束
	4.如父节点是红色的，则违反定义3，需调整节点颜色。

删除:
	1.首先使用普通的二叉查找树的删除算法删除节点，然后进行旋转及颜色的调整
	2.在二叉查找树的删除中，最坏可以归结到两种情况：
		- 删除叶节点
		- 删除只有一个儿子的节点
	3.只要解决了这两种情况下的着色问题，就解决了红黑树的删除
	4.删除叶节点：红色的直接删除，黑色删除后要调整
	  删除只有一个儿子的节点：将儿子的颜色改为黑色
	  删除有两个儿子的节点：找替身 -> 删一个儿子情况

自平衡: 左旋、右旋、变色
	变色：节点的颜色由红变黑或者由黑变红。
	左旋：以某个节点作为支点(旋转节点)，其右子节点变为旋转节点的父节点，右子节点的左子节点变为旋转节点的右子节点，左子节点保持不变。
	右旋：以某个节点作为支点(旋转节点)，其左子节点变为旋转节点的父节点，左子节点变为旋转节点的左子节点，右子节点保持不变。
*/
type RedBlackTree struct {
	// 红黑树的根节点
	root *redBlackTreeNode
	// 红黑树比较的函数
	compare fucker.CompareFunc
}

// 构建一颗红黑树
func NewRedBlackTree(compare fucker.CompareFunc, values ...interface{}) *RedBlackTree {
	t := &RedBlackTree{compare: compare}

	for _, v := range values {
		t.Insert(v)
	}

	return t
}

// 右旋：
//		旋转节点成为其左节点的右节点，其左节点代替旋转节点的位置，其左节点的右节点成为旋转节点的左节点
// 左旋：
//		旋转节点成为其右节点的左节点，其右节点代替旋转节点的位置，其右节点的左节点成为旋转节点的右节点
// 红黑树插入数据
// 1.如果红黑树为空，直接插入根节点，并且根节点为黑色
// 2.当插入的数据已存在，直接替换已存在的数据
// 3.插入节点的父节点是黑节点，由于插入的节点是红色的，当插入节点是黑色时，并不会影响红黑树的平衡，直接插入即可，无需做自平衡。
// 4.插入节点的父节点是红色的
//		4.1.叔叔节点存在并且为红节点
//			依据红黑树性质，红色节点不能相连 ==> 祖父节点肯定为黑节点
//			因为不可以同时存在两个相连的红节点，那么此时该插入子树的红黑层数的情况是：黑红红。显然最简单的处理方式是把其改为：红黑红
//			a.将父节点与叔叔节点改为黑色
//			b.将祖父节点改为红色
//			c.将祖父节点设置为当前节点，进行后续处理
//		4.2.叔叔节点不存在或者为黑节点，并且插入节点的父亲节点是祖父节点的左子节点
//			注意：单纯从插入前来看，叔叔节点非红即空(NIL节点)，否则的话破坏了红黑树性质5，此路径会比其它路径多一个黑色节点。
//			a.新插入节点，为其父节点的左子节点 LL双红。第一步先把父亲节点设置为黑色，祖父节点设置为红色；第二步祖父节点为旋转节点进行右旋。
//			b.新插入节点，为其父节点的右子节点 LR双红。第一步以父节点为旋转节点左旋，然后得到LL双红，进行a步骤。
//		4.3.叔叔节点不存在或者为黑节点，并且插入节点的父亲节点是祖父节点的右子节点
//			a.新插入节点，为其父节点的右子节点 RR双红。第一步先把父亲节点设置为黑色，祖父节点设置为红色；第二步祖父节点为旋转节点进行左旋。
//			b.新插入节点，为其父节点的右子节点 RL双红。第一步以父节点为旋转节点右旋，然后得到RR双红，进行a步骤。
func (r *RedBlackTree) Insert(value interface{}) {
	node := &redBlackTreeNode{value: value}
	if r.root == nil {
		// 当插入的节点是根节点时，根节点必须为黑色
		node.color = black
		r.root = node
		return
	}
	// 先着红色
	node.color = red
	parent := r.root

LOOP:
	for {
		switch r.compare(node, parent) {
		case fucker.Equal:
			parent.value = node.value
			return
		case fucker.Less:
			if parent.left == nil {
				node.parent = parent
				parent.left = node
				break LOOP
			} else {
				parent = parent.left
			}
		case fucker.Greater:
			if parent.right == nil {
				node.parent = parent
				parent.right = node
				break LOOP
			} else {
				parent = parent.right
			}
		default:
			break LOOP
		}
	}

	// 然后旋转
	r.rotate(node)
}

// 删除节点
func (r *RedBlackTree) Delete(value interface{}) {
	node := &redBlackTreeNode{value: value}

	for n := r.root; n != nil; {
		switch r.compare(node, n) {
		case fucker.Equal:
			parent := n.parent

			if parent == nil {
				if n.left == nil && n.right == nil {
					r.root = nil
				} else if n.left != nil && n.right == nil {
					n.parent = nil
					r.root = n.left
					n.color = black
				} else if n.left == nil && n.right != nil {
					n.parent = nil
					r.root = n.right
					n.color = black
				} else {
					n1 := n.right
					for n1.left != nil {
						n1 = n1.left
					}

					n1.parent.left = nil
					n.value = n1.value
				}
				return
			} else {
				isLeft := parent.left != nil && r.compare(node, parent.left) == fucker.Equal

				n.parent = nil
				if n.left == nil && n.right == nil {
					if isLeft {
						parent.left = nil
					} else {
						parent.right = nil
					}
				} else if n.left != nil && n.right == nil {
					n.left.color = n.color
					n.left.parent = n.parent
					if isLeft {
						parent.left = n.left
					} else {
						parent.right = n.left
					}
				} else if n.left == nil && n.right != nil {
					n.right.color = n.color
					n.right.parent = n.parent
					if isLeft {
						parent.left = n.right
					} else {
						parent.right = n.right
					}
				} else {
					n1 := n.right
					for n1.left != nil {
						n1 = n1.left
					}

					n1.parent.left = nil
					n.value = n1.value
				}
				return
			}

		case fucker.Greater:
			n = n.right
		case fucker.Less:
			n = n.left
		}
	}
}

// 寻找节点
func (r *RedBlackTree) Search(value interface{}) Value {
	node := r.root
	val := &redBlackTreeNode{value: value}

	for node != nil {
		switch r.compare(val, node) {
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

// 旋转
// 修复红黑树
func (r *RedBlackTree) rotate(node *redBlackTreeNode) {
	if node.parent == nil {
		node.color = black
		return
	}

	if node.parent != nil && node.parent.color == black {
		return
	}

	parent := node.parent
	grandfather := parent.parent
	// 情景4.插入节点的父节点是红色的
	if grandfather.left == nil || r.compare(parent, grandfather.left) == fucker.Equal {
		uncle := grandfather.right
		// 情景4.1叔叔存在并且为红色
		if uncle != nil && uncle.color == red {
			parent.color = black
			uncle.color = black
			grandfather.color = red
			r.rotate(grandfather)
			return
		}

		// 情景4.2 叔叔节点不存在或者为黑色
		// 情景4.2.1 插入节点为其父节点的左子节点(LL)，将父节点染为黑色，将祖父染色为红色，然后以祖父节点右旋，就完成了
		if r.compare(node, parent.left) == fucker.Equal {
			parent.color = black
			grandfather.color = red
			r.rightRotate(grandfather)
		} else {
			// 情景4.2.2：插入节点为其父节点的右子节点(LR)
			// 以父节点进行一次左旋，得到LL双红，然后以父节点为当前节点进行下一轮处理
			r.leftRotate(parent)
			r.rotate(parent)
		}
	} else {
		uncle := grandfather.left
		// 情景4.1叔叔存在并且为红色
		if uncle != nil && uncle.color == red {
			parent.color = black
			uncle.color = black
			grandfather.color = red
			r.rotate(grandfather)
			return
		}

		// 情景4.3叔叔不存在或者为黑色
		if uncle == nil || uncle.color == black {
			// 情景4.3.1：插入节点为其父节点的右子节点(RR)，将父节点染色为黑色，将祖父节点染色为红色，然后祖父节点右旋
			if r.compare(node, parent.right) == fucker.Equal {
				parent.color = black
				grandfather.color = red
				r.leftRotate(grandfather)
			} else {
				// 情景4.3.2 插入节点为其父节点的左子节点(RL)
				// 以父节点进行一次右旋，得到RR双红，然后指定父节点为当前节点进行下一轮处理
				r.rightRotate(parent)
				r.rotate(parent)
			}
		}
	}
}

/*
LL旋转
左旋x节点
        p                p
        |                |
        x                y
       / \        -->   / \
      lx  y            x   ry
         / \          / \
        ly ry        lx  ly

	1.将y的左子节点ly的父节点更新为x，并将x的右子节点指向y的左子节点ly
	2.当x的父节点不为空时，更新y的父节点为x的父节点，x的父节点的子树为y
	3.将y的左子节点更新为x
*/
func (r *RedBlackTree) leftRotate(x *redBlackTreeNode) {
	y := x.right
	p := x.parent

	// 	1.将y的左子节点ly的父节点更新为x，并将x的右子节点指向y的左子节点
	x.right = y.left
	if y.left != nil {
		y.left.parent = x.right
	}

	// 	2.当x的父节点不为空时，更新y的父节点为x的父节点，x的父节点的子树为y
	y.parent = p
	if p == nil { // 父节点为空时，说明x是根节点
		r.root = y
	} else {
		if r.compare(p.left, x) == fucker.Equal {
			p.left = y
		} else {
			p.right = y
		}
	}

	// 	3.将y的左子节点更新为x
	y.left = x
	x.parent = y
}

/*
右旋y节点
        p                p
        |                |
        y                x
       / \        -->   / \
      x  ry            x   y
     / \                  / \
    lx rx                lx  ry

	1.将x的右子节点rx的父节点更新为y，并将y的左子节点指向x的右子节点rx
	2.当y的父节点p不为空时，更新x的父节点为y的父节点p，y的父节点p的子树为x
	3.将x的右子节点更新为y
*/
func (r *RedBlackTree) rightRotate(y *redBlackTreeNode) {
	x := y.left
	p := y.parent

	// 1.将x的右子节点rx的父节点更新为y，并将y的左子节点指向x的右子节点rx
	y.left = x.right
	if x.right != nil {
		x.right.parent = y.left
	}

	// 2.当y的父节点p不为空时，更新x的父节点为y的父节点p，y的父节点p的子树为x
	x.parent = p
	if p == nil {
		r.root = x
	} else {
		if r.compare(p.left, y) == fucker.Equal {
			p.left = x
		} else {
			p.right = x
		}
	}

	// 3.将x的右子节点更新为y
	x.right = y
	y.parent = x
}
