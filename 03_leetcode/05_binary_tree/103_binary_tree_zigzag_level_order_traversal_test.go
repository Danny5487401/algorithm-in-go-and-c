package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二叉树的锯齿形层序遍历 https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/
func TestZigzagLevelOrder(t *testing.T) {
	convey.Convey("二叉树的锯齿形层序遍历:先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行", t, func() {
		testCase := []struct {
			input  *TreeNode
			target [][]int
		}{
			{
				Ints2TreeNode([]int{3, 9, 20, NULL, NULL, 15, 7}),
				[][]int{
					{3},
					{20, 9},
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
			rsp := zigzagLevelOrder(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})
}
func zigzagLevelOrder(root *TreeNode) [][]int {
	// 反转偶数层数组即可
	var ans = make([][]int, 0)
	if root == nil {
		return ans
	}

	queue := []*TreeNode{root} // 辅助队列
	var even bool              // 奇数层
	for len(queue) > 0 {
		n := len(queue)         // 注意先求长度
		value := make([]int, n) // 每一层的数据
		for i := range value {
			node := queue[0]
			queue = queue[1:]

			if !even {
				value[i] = node.Val
			} else {
				value[n-1-i] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ans = append(ans, value)
		even = !even

	}

	return ans
}
