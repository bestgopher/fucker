package tree

type avlTreeNode struct {
	value  interface{}  // value
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

func (a *AVLTree) Insert(value interface{}) {
	if a.root == nil {
		a.root = &avlTreeNode{value: value, height: 1}
	}

}

func (a *AVLTree) Delete(value interface{}) {
	panic("implement me")
}

func (a *AVLTree) Search(value interface{}) bool {
	panic("implement me")
}

// isBalance checks whether the node is balanced
func (a *AVLTree) isBalance(node *avlTreeNode) bool {
	left := 0
	if node.left != nil {
		left = node.left.height
	}

	right := 0
	if node.right != nil {
		left = node.right.height
	}

	return (right-left) >= -1 || (right-left) <= 1
}
