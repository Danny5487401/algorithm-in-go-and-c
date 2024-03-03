package _5_binary_tree

// 二叉树的层序遍历 https://leetcode.cn/problems/binary-tree-level-order-traversal/description/?envType=study-plan-v2&envId=top-100-liked

// 一个队列
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		// 队列的长度及是下方循环的次数
		n := len(queue)
		value := make([]int, n)
		for i := range value {
			node := queue[0]
			queue = queue[1:]
			value[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, value)

	}
	return ans
}

// 缺点：cur 和 next 两个数组
func levelOrderDepreciated(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	cur := []*TreeNode{root}
	for cur != nil {
		next := make([]*TreeNode, 0)   // 树的节点队列
		value := make([]int, len(cur)) // 用来存储每一层的值，到新的一层时需要进行清空
		for i, node := range cur {
			value[i] = node.Val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		cur = next
		ans = append(ans, value)
	}
	return ans
}
