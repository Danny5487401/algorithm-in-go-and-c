package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"slices"
	"testing"
)

func TestClosestNodes(t *testing.T) {
	convey.Convey("二叉搜索树最近节点查询", t, func() {
		testCase := []struct {
			input  *TreeNode
			query  []int
			target [][]int
		}{
			{
				Ints2TreeNode([]int{6, 2, 13, 1, 4, 9, 15, NULL, NULL, NULL, NULL, NULL, NULL, 14}),
				[]int{2, 5, 16},
				[][]int{
					{2, 2},
					{4, 6},
					{15, -1},
				},
			},

			{
				Ints2TreeNode([]int{4, NULL, 9}),
				[]int{3},
				[][]int{
					{-1, 4},
				},
			},
		}

		for _, tst := range testCase {
			rsp := closestNodes(tst.input, tst.query)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func closestNodes(root *TreeNode, queries []int) [][]int {
	// n == queries.length
	// 1 <= n <= 105

	// 二叉搜索中找到 min,max
	// 94. 二叉树的中序遍历 得到有一个严格递增数组 a，再在 a 上做二分查找

	var array = make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		array = append(array, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	ans := make([][]int, len(queries))
	for i, q := range queries {
		minVal, maxVal := -1, -1
		// BinarySearch 升序的 slices: 如果找到返回索引,没有找到返回应该插入的位置
		index, ok := slices.BinarySearch(array, q)

		// 最大值
		if index < len(array) {
			// 右边界,代表插入位置
			maxVal = array[index]

		}
		/*
			对于min:
				1. 如果 j<n 且 a[j]=q, 那么 min =a[j]
				2. 否则找不到,如果 j>0，那么 min=a[j−1]
				3. 否则 -1
		*/

		if !ok {
			// 没有找到, 代表 // a[j]>q, a[j-1]<q
			index--
		}
		if index >= 0 {
			// 代表有左边界
			minVal = array[index]
		}
		ans[i] = []int{minVal, maxVal}
	}

	return ans

}
