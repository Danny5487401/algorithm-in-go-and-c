package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"slices"
	"testing"
)

// 二叉树中的第 K 大层和 https://leetcode.cn/problems/kth-largest-sum-in-a-binary-tree/
func TestKthLargestLevelSum(t *testing.T) {
	convey.Convey("二叉树中的第 K 大层和 ", t, func() {
		testCase := []struct {
			input  *TreeNode
			k      int
			target int
		}{
			{
				Ints2TreeNode([]int{5, 8, 9, 2, 1, 3, 7, 4, 6}),
				2,
				13,
			},

			{
				Ints2TreeNode([]int{1, 2, NULL, 3}),
				1,
				3,
			},
		}

		for _, tst := range testCase {
			rsp := kthLargestLevelSum(tst.input, tst.k)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func kthLargestLevelSum(root *TreeNode, k int) int64 {

	var depthAns []int64
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		tempQueue := make([]int, len(queue))
		var sum int
		for _ = range tempQueue {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depthAns = append(depthAns, int64(sum))
	}
	n := len(depthAns)
	if k > n {
		return -1
	}
	slices.Sort(depthAns) // 注意这里是升序
	return depthAns[n-k]
}
