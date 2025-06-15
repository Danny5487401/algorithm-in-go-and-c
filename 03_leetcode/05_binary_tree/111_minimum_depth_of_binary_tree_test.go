package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 二叉树的最小深度 https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/
func TestMinDepth(t *testing.T) {
	convey.Convey("二叉树的最小深度: 最短路径,左右儿子为空", t, func() {
		testCase := []struct {
			input  *TreeNode
			target int
		}{
			{
				Ints2TreeNode([]int{3, 9, 20, NULL, NULL, 15, 7}),
				2,
			},
			{
				Ints2TreeNode([]int{2, NULL, 3, NULL, 4, NULL, 5, NULL, 6}),
				5,
			},
		}

		for _, tst := range testCase {
			rsp := minDepth(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

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

// 自底向上「归」
func minDepth2(root *TreeNode) int {
	if root == nil {
		// 如果 node 是空节点，由于没有节点，返回 0。
		return 0
	}
	// 如果 node 没有右儿子，那么深度就是左子树的深度加一，即 dfs(node)=dfs(node.left)+1
	if root.Right == nil {
		return minDepth2(root.Left) + 1
	}

	// 如果 node 没有左儿子，那么深度就是右子树的深度加一，即 dfs(node)=dfs(node.right)+1
	if root.Left == nil {
		return minDepth2(root.Right) + 1
	}
	// node 左右儿子都有，那么分别递归计算左子树的深 dfs(node)=min(dfs(node.left),dfs(node.right))+1
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
		// 如果递归中发现 depth≥ans，由于继续向下递归也不会让 depth 变小，直接返回。
		if depth >= ans {
			return // 最优性剪枝
		}

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
