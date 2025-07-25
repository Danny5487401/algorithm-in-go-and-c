package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 求根节点到叶节点数字之和 https://leetcode.cn/problems/sum-root-to-leaf-numbers/description/
func TestSumNumbers(t *testing.T) {
	convey.Convey("求根节点到叶节点数字之和 :计算从根节点到叶节点生成的 所有数字之和 ", t, func() {
		testCase := []struct {
			input  *TreeNode
			target int
		}{
			{
				/*
					从根到叶子节点路径 1->2 代表数字 12
					从根到叶子节点路径 1->3 代表数字 13
					因此，数字总和 = 12 + 13 = 25
				*/
				Ints2TreeNode([]int{1, 2, 3}),
				25,
			},
			{
				/*
					从根到叶子节点路径 4->9->5 代表数字 495
					从根到叶子节点路径 4->9->1 代表数字 491
					从根到叶子节点路径 4->0 代表数字 40
					因此，数字总和 = 495 + 491 + 40 = 1026
				*/
				Ints2TreeNode([]int{4, 9, 0, 5, 1}),
				1026,
			},
		}

		for _, tst := range testCase {
			rsp := sumNumbers(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func sumNumbers(root *TreeNode) int {

	var sum int
	var dfs func(node *TreeNode, tmpSum int)
	dfs = func(node *TreeNode, tmpSum int) {
		if node == nil {
			return
		}
		tmpSum = tmpSum*10 + node.Val
		if node.Left == nil && node.Right == nil { // node 是叶子节点,0 <= Node.val <= 9
			sum += tmpSum
			return
		}

		dfs(node.Left, tmpSum)
		dfs(node.Right, tmpSum)
	}

	dfs(root, 0)
	return sum

}
