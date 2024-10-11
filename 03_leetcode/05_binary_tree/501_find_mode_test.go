package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二叉搜索树中的众数 https://leetcode.cn/problems/find-mode-in-binary-search-tree/description/
func TestFindMode(t *testing.T) {
	convey.Convey("二叉搜索树中的众数", t, func() {
		testCase := []struct {
			input  *TreeNode
			target []int
		}{
			{
				Ints2TreeNode([]int{1, NULL, 2, 2}),
				[]int{2},
			},
			{
				Ints2TreeNode([]int{0}),
				[]int{0},
			},
		}

		for _, tst := range testCase {
			rsp := findMode(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func findMode(root *TreeNode) (answer []int) {
	// -105 <= Node.val <= 105
	var base, count, maxCount int

	update := func(x int) {
		// 首先更新 base 和 count
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		// 然后更新 maxCount
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return
}
