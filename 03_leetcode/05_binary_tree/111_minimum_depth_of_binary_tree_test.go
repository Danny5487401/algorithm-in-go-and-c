package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 二叉树的最小深度 https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/
func TestMinDepth(t *testing.T) {
	convey.Convey("二叉树的最小深度: 最短路径", t, func() {
		testCase := []struct {
			input  *TreeNode
			target int
		}{
			{
				CreateTreeByArray([]int{3, 9, 20, 0, 0, 15, 7}),
				2,
			},
			{
				CreateTreeByArray([]int{2, 0, 3, 0, 4, 0, 5, 0, 6}),
				5,
			},
		}

		for _, tst := range testCase {
			rsp := minDepth2(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 自底向上「归」
func minDepth2(root *TreeNode) int {
	if root == nil {
		// 如果 node 是空节点，由于没有节点，返回 0。
		return 0
	}
	// 如果 node 没有右儿子，那么深度就是左子树的深度加一，即 dfs(node)=dfs(node.left)+1
	if root.Left != nil {
		return minDepth2(root.Left) + 1
	}

	// 如果 node 没有左儿子，那么深度就是右子树的深度加一，即 dfs(node)=dfs(node.right)+1
	if root.Right != nil {
		return minDepth2(root.Right) + 1
	}
	// dfs(node)=min(dfs(node.left),dfs(node.right))+1
	return min(minDepth2(root.Left), minDepth2(root.Right)) + 1
}

// 自顶向下「递」
func minDepth3(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var ans = math.MaxInt
	var dfs func(node *TreeNode, depth int)

	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		depth++
		if node.Left == nil && node.Right == nil { // 要求左右都为叶子节点
			ans = min(depth, ans)
			return
		}
		dfs(node.Left, depth)
		dfs(node.Right, depth)
	}

	dfs(root, 0)
	return ans
}

// 使用深度优先搜索的方法，遍历整棵树，记录最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}
	return minD + 1
}
