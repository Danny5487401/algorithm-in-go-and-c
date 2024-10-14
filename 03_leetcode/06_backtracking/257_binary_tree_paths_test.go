package _6_backtracking

import (
	"github.com/smartystreets/goconvey/convey"
	"strconv"
	"strings"
	"testing"
)

// 二叉树的所有路径 https://leetcode.cn/problems/binary-tree-paths/
func TestBinaryTreePaths(t *testing.T) {
	convey.Convey("二叉树的所有路径:深度优先", t, func() {
		testCase := []struct {
			input  *TreeNode
			target []string
		}{

			{
				Ints2TreeNode([]int{1, 2, 3, NULL, 5}),
				[]string{"1->2->5", "1->3"},
			},

			{
				Ints2TreeNode([]int{1}),
				[]string{"1"},
			},
		}

		for _, tst := range testCase {
			rsp := binaryTreePaths(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func binaryTreePaths(root *TreeNode) []string {

	var ans []string

	var dfs func(node *TreeNode, path []string)

	dfs = func(node *TreeNode, path []string) {
		if node != nil {
			path = append(path, strconv.Itoa(node.Val))
			if node.Left == nil && node.Right == nil { // 当前节点是叶子节点
				ans = append(ans, strings.Join(path, "->"))
				return
			}

			dfs(node.Left, path)
			dfs(node.Right, path)
		}

	}
	dfs(root, nil)

	return ans
}
