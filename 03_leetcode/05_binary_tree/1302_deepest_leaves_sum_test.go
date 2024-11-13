package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 层数最深叶子节点的和 https://leetcode.cn/problems/deepest-leaves-sum/

func TestDeepestLeavesSum(t *testing.T) {
	convey.Convey("层数最深叶子节点的和", t, func() {
		testCase := []struct {
			input  *TreeNode
			target int
		}{
			{
				Ints2TreeNode([]int{1, 2, 3, 4, 5, NULL, 6, 7, NULL, NULL, NULL, NULL, 8}),
				15,
			},

			{
				Ints2TreeNode([]int{6, 7, 8, 2, 7, 1, 3, 9, NULL, 1, 4, NULL, NULL, NULL, 5}),
				19,
			},
		}

		for _, tst := range testCase {
			rsp := deepestLeavesSum(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}
func deepestLeavesSum(root *TreeNode) (sum int) {
	if root == nil {
		return 0
	}
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		sum = 0
		value := make([]*TreeNode, len(queue))
		for _ = range value {
			cur := queue[0]
			queue = queue[1:]
			sum += cur.Val
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

	}
	return
}
