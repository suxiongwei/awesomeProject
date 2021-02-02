package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return tree_max(maxDepth(root.left), maxDepth(root.right)) + 1
}

func tree_max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func levelOrderSimple(root *TreeNode) [][]int {
	return dfs(root, 0, [][]int{})
}

func dfs(root *TreeNode, level int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	// 这一层如果不存在 需要追加一层（当前的层数为3 len为2 所以需要追加一层）
	if len(res) == level {
		res = append(res, []int{root.value})
	}else {
		res[level] = append(res[level], root.value)
	}
	res = dfs(root.left, level + 1, res)
	res = dfs(root.right, level + 1, res)
	return res
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	// 定义一个双向队列
	queue := list.New()
	// 头部插入根节点
	queue.PushFront(root)
	// 进行广度搜索
	for queue.Len() > 0 {
		var current []int
		listLength := queue.Len()
		// 将本层的放入current（也就是从队列的尾部移出去），将下一层的放入队列的头部
		for i := 0; i < listLength; i++ {
			// 消耗尾部
			// queue.Remove(queue.Back()).(*TreeNode)：移除最后一个元素并将其转化为TreeNode类型
			node := queue.Remove(queue.Back()).(*TreeNode)
			current = append(current, node.value)
			if node.left != nil {
				//插入头部
				queue.PushFront(node.left)
			}
			if node.right != nil {
				queue.PushFront(node.right)
			}
		}
		result = append(result, current)
	}
	return result
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.value {
		root.left = deleteNode( root.left, key )
		return root
	}
	if key > root.value {
		root.right = deleteNode( root.right, key )
		return root
	}
	//到这里意味已经查找到目标
	if root.right == nil {
		//右子树为空
		return root.left
	}
	if root.left == nil {
		//左子树为空
		return root.right
	}
	minNode := root.right
	for minNode.left != nil {
		//查找后继
		minNode = minNode.left
	}
	root.value = minNode.value
	// 最终return的还是root.right
	root.right = deleteMinNode( root.right )
	return root
}

/**
最终返回的还是root（递归） 本方法作用是递归删除最小的节点，每一层递归root一直在变，直到找到左子树为空，此时找到了最小值，
然后在递归往上更新root的left，因为deleteMinNode 传的就是root->left,所以返回的root.left = tmpNode
 */
func deleteMinNode( root *TreeNode ) *TreeNode {
	if root.left == nil {
		pRight := root.right
		root.right = nil
		return pRight
	}
	tmpNode := deleteMinNode(root.left)
	// 有这一行的作用是删除了节点后会，如果最小节点有右子树，返回右子树，否则返回左子树的nil，当递归到最小节点6时，return nil，
	//此时return到上一层 root=8 root8.left = nil,然后一直往上层走root.left = tmpNode 10.left = 8, 最终递归结束返回root10
	root.left = tmpNode
	return root
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.left) || !isBalanced(root.right) {
		return false
	}
	leftH := maxDepth(root.left) + 1;
	rightH := maxDepth(root.right) + 1;
	if abs(leftH-rightH) > 1 {
		return false
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main13()  {
	//node15 := TreeNode{15, nil, nil}
	//node7 := TreeNode{7, nil, nil}
	//node9 := TreeNode{9, nil, nil}
	//node20 := TreeNode{20, &node15, &node7}
	//node3 := TreeNode{3, &node9, &node20}
	//
	//order := levelOrder(&node3)
	//fmt.Println(order)

	//node2 := TreeNode{2, nil, nil}
	//node4 := TreeNode{4, nil, nil}
	//node6 := TreeNode{6, nil, nil}
	//node9 := TreeNode{9, nil, nil}
	//node8 := TreeNode{8, &node6, &node9}
	//node3 := TreeNode{3, &node2, &node4}
	//node11 := TreeNode{11, nil, nil}
	//node10 := TreeNode{10, &node8, &node11}
	//node5 := TreeNode{5, &node3, &node10}
	//
	//node := deleteNode(&node5, 5)
	//fmt.Println(node.value)
	//fmt.Println(node.left.value)
	//fmt.Println(node.right.value)


	//node2 := TreeNode{2, nil, nil}
	//node4 := TreeNode{4, nil, nil}
	node6 := TreeNode{6, nil, nil}
	node9 := TreeNode{9, nil, nil}
	node8 := TreeNode{8, &node6, &node9}
	//node3 := TreeNode{3, &node2, &node4}
	node11 := TreeNode{11, nil, nil}
	node10 := TreeNode{10, &node8, &node11}
	//node5 := TreeNode{5, &node3, &node10}
	node5 := TreeNode{5, nil, &node10}


	fmt.Println(isBalanced(&node5))

}
