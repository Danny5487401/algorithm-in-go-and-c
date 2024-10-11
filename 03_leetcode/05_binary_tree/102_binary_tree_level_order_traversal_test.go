package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二叉树的层序遍历 https://leetcode.cn/problems/binary-tree-level-order-traversal/description/?envType=study-plan-v2&envId=top-100-liked
func TestLevelOrder(t *testing.T) {
	convey.Convey("二叉树的层序遍历: 自顶向下的层序遍历", t, func() {
		testCase := []struct {
			input  *TreeNode
			target [][]int
		}{
			{
				Ints2TreeNode([]int{3, 9, 20, NULL, NULL, 15, 7}),
				[][]int{
					{3},
					{9, 20},
					{15, 7},
				},
			},
			{
				Ints2TreeNode([]int{1}),
				[][]int{
					{1},
				},
			},
			{
				Ints2TreeNode([]int{}),
				[][]int{},
			},
		}

		for _, tst := range testCase {
			rsp := levelOrder(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})
}

// 一个队列
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root} // 辅助队列
	for len(queue) != 0 {
		// 队列的长度及是下方循环的次数
		n := len(queue)
		value := make([]int, n) // 每一层的数据
		for i := range value {
			node := queue[0]
			queue = queue[1:]
			value[i] = node.Val
			// 将左右儿子入队
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
