package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二叉树的垂序遍历 https://leetcode.cn/problems/vertical-order-traversal-of-a-binary-tree/description/

func TestVerticalTraversal(t *testing.T) {
	convey.Convey(" 从最左边的列开始直到最右边的列结束，按列索引每一列上的所有结点，形成一个按出现位置从上到下排序的有序列表。如果同行同列上有多个结点，则按结点的值从小到大进行排序", t, func() {
		testCase := []struct {
			input  *TreeNode
			target [][]int
		}{
			{
				/*
					列 -1 ：只有结点 9 在此列中。
					列  0 ：只有结点 3 和 15 在此列中，按从上到下顺序。
					列  1 ：只有结点 20 在此列中。
					列  2 ：只有结点 7 在此列中。
				*/
				Ints2TreeNode([]int{3, 9, 20, NULL, NULL, 15, 7}),
				[][]int{
					{9}, {3, 15}, {20}, {7},
				},
			},
			{
				Ints2TreeNode([]int{1, 2, 3, 4, 5, 6, 7}),
				[][]int{
					{4}, {2}, {1, 5, 6}, {3}, {7},
				},
			},
			{
				Ints2TreeNode([]int{1, 2, 3, 4, 6, 5, 7}),
				[][]int{
					{4}, {2}, {1, 5, 6}, {3}, {7},
				},
			},
		}

		for _, tst := range testCase {
			rsp := verticalTraversal(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func verticalTraversal(root *TreeNode) [][]int {
	res := [][]int{}

	return res
}
