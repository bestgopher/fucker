package tree

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	fucker "gopkg.in/bestgopher/fucker/v1"
)

// 插入情况为rr旋转
func TestAVLTreeInsertRR(t *testing.T) {
	ass := assert.New(t)
	// 10
	//  \
	//  20
	//   \
	//   30
	ass.Nil(testAvlInsert([]interface{}{10, 20, 30}))
	//   10
	//   /\
	//  8  20
	//      \
	//      30
	//       \
	//       40
	ass.Nil(testAvlInsert([]interface{}{10, 8, 20, 30, 40}))
	//   10
	//   /\
	//  8  20
	//     /\
	//    15 30
	//       /
	//      25
	ass.Nil(testAvlInsert([]interface{}{10, 8, 20, 15, 30, 25}))
	//   10
	//   /\
	//  8  20
	//     /\
	//    15 30
	//        \
	//        40
	ass.Nil(testAvlInsert([]interface{}{10, 8, 20, 15, 30, 40}))
}

// 插入情况为ll旋转
func TestAVLTreeInsertLL(t *testing.T) {
	ass := assert.New(t)
	//    30
	//    /
	//   20
	//   /
	//  10
	ass.Nil(testAvlInsert([]interface{}{30, 20, 10}))
	//        30
	//        / \
	//       20 40
	//       /
	//      10
	//      /
	//     5
	ass.Nil(testAvlInsert([]interface{}{30, 20, 40, 10, 5}))
	//        30
	//        / \
	//       20 40
	//       / \
	//      10 25
	//      /
	//     5
	ass.Nil(testAvlInsert([]interface{}{30, 20, 40, 10, 25, 5}))
	//        30
	//        / \
	//       20 40
	//       / \
	//      10 25
	//       \
	//       15
	ass.Nil(testAvlInsert([]interface{}{30, 20, 40, 10, 25, 15}))
}

func TestAVLTreeInsertLR(t *testing.T) {
	ass := assert.New(t)
	//  30
	//   \
	//   40
	//   /
	//  35
	ass.Nil(testAvlInsert([]interface{}{30, 40, 35}))
	//     30
	//    / \
	//   20 40
	//     / \
	//    35 50
	//    /
	//   32
	ass.Nil(testAvlInsert([]interface{}{30, 20, 40, 35, 50, 32}))
	//     30
	//    / \
	//   20 40
	//     / \
	//    35 50
	//     \
	//     36
	ass.Nil(testAvlInsert([]interface{}{30, 20, 40, 35, 50, 36}))
}

// 插入情况为rl旋转
func TestAVLTreeInsertRL(t *testing.T) {
	ass := assert.New(t)
	//    30
	//    /
	//   20
	//    \
	//    25
	ass.Nil(testAvlInsert([]interface{}{30, 20, 25}))
	//       30
	//       / \
	//      20 40
	//     / \
	//    10 25
	//        \
	//        28
	ass.Nil(testAvlInsert([]interface{}{30, 20, 40, 10, 25, 18}))
	//        30
	//       / \
	//      20 40
	//     / \
	//    10 25
	//       /
	//      18
	ass.Nil(testAvlInsert([]interface{}{30, 20, 40, 10, 25, 15}))
}

func TestAVLTreeSearch(t *testing.T) {
	avlTree := NewAVLTree(CompareBstTreeInt, 3, 4, 1, 5, 6, 7)
	searchNode(avlTree, t)
}

func TestAVLTreeDelete(t *testing.T) {

	v := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(v); i++ {
		avlTree1 := NewAVLTree(CompareBstTreeInt, v...)
		assert.Nil(t, testCheckAvl(avlTree1))
		assert.NotNil(t, avlTree1.Search(v[i]))
		avlTree1.Delete(v[i])
		assert.Nil(t, avlTree1.Search(v[i]))
		assert.Nilf(t, testCheckAvl(avlTree1), "删除节点 %d", v[i])
	}
}

// 测试插入的数据
func testAvlInsert(values []interface{}) error {
	// 先找出最大值列表的最大值
	tree := NewAVLTree(CompareBstTreeInt, values...)
	return testCheckAvl(tree)
}

func testCheckAvl(tree *AVLTree) error {
	if !testCheckValue(tree.root) {
		return errors.New("检查值失败")
	}

	if !testCheckRelationship(tree.root, true) {
		return errors.New("检查关系错误")
	}

	if !testCheckHeight(tree.root) {
		return errors.New("检查深度错误")
	}

	return nil
}

// 检查值
// 检查左子节点的值是否小于当前节点，右子节点的值是否大于当前节点
func testCheckValue(node *avlTreeNode) bool {
	if node == nil {
		return true
	}

	if node.left != nil {
		if CompareBstTreeInt(node.left, node) == fucker.Greater {
			return false
		}
	}

	if node.right != nil {
		if CompareBstTreeInt(node.right, node) == fucker.Less {
			return false
		}
	}

	return testCheckValue(node.left) && testCheckValue(node.right)
}

// 检查父子关系
func testCheckRelationship(node *avlTreeNode, isRoot bool) bool {
	if node == nil {
		return true
	}

	if isRoot {
		if node.parent != nil {
			return false
		}
	} else {
		if !((node.parent.right != nil && CompareBstTreeInt(node, node.parent.right) == fucker.Equal) ||
			(node.parent.left != nil && CompareBstTreeInt(node, node.parent.left) == fucker.Equal)) {
			return false
		}
	}

	return testCheckRelationship(node.left, false) && testCheckRelationship(node.right, false)
}

// 检查深度
func testCheckHeight(node *avlTreeNode) bool {
	if node == nil {
		return true
	}

	if !(testCheckHeight(node.left) && testCheckHeight(node.right)) {
		return false
	}

	// 检查左右子树高度差是否为小于1
	check := func(left, right *avlTreeNode) bool {
		leftHeight := 0
		if left != nil {
			leftHeight = left.height
		}

		rightHeight := 0
		if right != nil {
			rightHeight = right.height
		}

		return (leftHeight-rightHeight) >= -1 && (leftHeight-rightHeight) <= 1
	}

	maxHeight, leftNode, rightNode := 0, node.left, node.right
	if leftNode != nil && rightNode != nil {
		if leftNode.height > rightNode.height {
			maxHeight = leftNode.height
		} else {
			maxHeight = rightNode.height
		}
	} else if leftNode != nil && rightNode == nil {
		maxHeight = leftNode.height
	} else if leftNode == nil && rightNode != nil {
		maxHeight = rightNode.height
	} else {
		maxHeight = 0
	}

	return node.height == maxHeight+1 && check(leftNode, rightNode)
}
