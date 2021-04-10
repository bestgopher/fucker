### fucker
一些常见的数据结构与算法的golang实现。



### 排序(sort)

| 算法                                                         | 时间复杂度          | 空间复杂度 | 稳定性 |
| ------------------------------------------------------------ | ------------------- | ---------- | ------ |
| [Bubble](https://github.com/bestgopher/fucker/blob/master/sort/internal/bubble.go) | O( n^2 )            | O(1)       | yes    |
| [Selection](https://github.com/bestgopher/fucker/blob/master/sort/internal/selection.go) | O( n^2 )            | O(1)       | no     |
| [Heap](https://github.com/bestgopher/fucker/blob/master/sort/internal/heap.go) | O(nlogn)            | O(1)       | no     |
| [Insertion](https://github.com/bestgopher/fucker/blob/master/sort/internal/insertion.go) | O(n^2)              | O(1)       | yes    |
| [Merge](https://github.com/bestgopher/fucker/blob/master/sort/internal/merge.go) | O(nlogn)            | O(n)       | yes    |
| [Quick](https://github.com/bestgopher/fucker/blob/master/sort/internal/quick.go) | O(nlogn)            | O(1)       | no     |
| [Shell](https://github.com/bestgopher/fucker/blob/master/sort/internal/shell.go) | O(n^(4/3)) ~ O(n^2) | 取决于步长 | no     |

- 概念

  - 排序的稳定性(stability)

     - 如果两个相等的元素，在排序前后的相对位置保持不变，那么这是稳定的排序算法。

       ```
       排序前：5 1 3(1) 4 7 3(2)
       稳定的排序：1 3(1) 3(2) 4 5 7
       稳定的排序：1 3(2) 3(1) 4 5 7
       ```
       
       稳定的排序就是排序后第一个`3`还是在第二个`3`前面。如果两个`3`的相对次序交换了，则就不是稳定的排序。

  - 原地算法(In-place Algorithm)

     - 不依赖额外的资源或者依赖少数的额外资源，仅依靠输出来覆盖输入
     - 空间复杂度为O(1)的都可以认为是原地算法

### 树(tree)

-  [二叉查找树(Binary Search Tree，BST树)](https://github.com/bestgopher/fucker/blob/master/tree/binary_search_tree.go)
-  [二叉平衡查找树(AVL Tree)](https://github.com/bestgopher/fucker/blob/master/tree/avl_tree.go)
-  [红黑树(Red-Black Tree)](https://github.com/bestgopher/fucker/blob/master/tree/red_black_tree.go)

### 字符串查找算法(strs)
-  [KMP算法](https://github.com/bestgopher/fucker/blob/master/strs/kmp.go)
