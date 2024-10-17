package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 将有序数组转换为二叉搜索树 https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/description/?envType=study-plan-v2&envId=top-100-liked
func TestSortedArrayToBST(t *testing.T) {
	convey.Convey("将有序数组转换为二叉搜索树", t, func() {
		testCase := []struct {
			input  []int
			target *TreeNode
		}{
			{
				[]int{-10, -3, 0, 5, 9},
				Ints2TreeNode([]int{0, -3, 9, -10, NULL, 5}),
			},
			{
				[]int{1, 3},
				Ints2TreeNode([]int{3, 1}),
			},
		}

		for _, tst := range testCase {
			rsp := sortedArrayToBST(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})
}

func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}
	mid := length / 2
	root := &TreeNode{nums[mid], nil, nil}

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
