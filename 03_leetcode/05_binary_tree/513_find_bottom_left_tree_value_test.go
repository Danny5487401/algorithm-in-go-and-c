package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 找树左下角的值 https://leetcode.cn/problems/find-bottom-left-tree-value/
func TestFindBottomLeftValue(t *testing.T) {
	convey.Convey("找树左下角的值", t, func() {
		testCase := []struct {
			input  *TreeNode
			target int
		}{
			{
				&TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				1,
			},
			{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 5,
							Left: &TreeNode{
								Val: 7,
							},
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
				7,
			},
		}

		for _, tst := range testCase {
			rsp := findBottomLeftValue2(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})
}

// 方案一：层次遍历第一个值
func findBottomLeftValue(root *TreeNode) int {
	return 0
}

// 方案二：从右到左，所以右节点先入队
func findBottomLeftValue2(root *TreeNode) int {

	node := root // 最后出队
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)

		}
	}
	return node.Val

}
