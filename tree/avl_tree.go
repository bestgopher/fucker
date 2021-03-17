package tree

/*
平衡二叉查找树

	平衡(balance): 左深和右深的差距小于2(即左深==右深，或者左深-右深=1，或者右深-左深=1)
	左(右)深: 左(右)子树的深度

	旋转：RR右旋，LL左旋，RL先右后左旋，LR先左后右旋

	LL：加入节点在左孩子的左孩子上
	RR：加入节点在右孩子的右孩子上

*/
type AVLTree struct {
	root    *avlTreeNode
	compare CompareFunc
}

type avlTreeNode struct {
	value  interface{}  // value
	left   *avlTreeNode // left node
	right  *avlTreeNode // right node
	height int          // height
}

func (a *avlTreeNode) Value() interface{} { return a.value }

func NewAVLTree(compare CompareFunc, values ...int) *AVLTree {
	tree := &AVLTree{compare: compare}

	for _, v := range values {
		tree.Insert(v)
	}

	return tree
}

func (a *AVLTree) Insert(value int) {
	if a.root == nil {
		a.root = &avlTreeNode{value: value, height: 1}
	}

}
