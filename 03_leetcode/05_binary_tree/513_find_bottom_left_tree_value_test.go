package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 找树左下角的值 https://leetcode.cn/problems/find-bottom-left-tree-value/
func TestFindBottomLeftValue(t *testing.T) {
	convey.Convey("找树左下角的值: 最底层 最左边 节点的值", t, func() {
		testCase := []struct {
			input  *TreeNode
			target int
		}{
			{
				Ints2TreeNode([]int{2, 1, 3}),
				1,
			},
			{
				Ints2TreeNode([]int{1, 2, 3, 4, NULL, 5, 6, NULL, NULL, 7}),
				7,
			},
		}

		for _, tst := range testCase {
			rsp := findBottomLeftValue(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})
}

// 方案一：层次遍历第一个值
func findBottomLeftValue(root *TreeNode) int {
	// 节点数目 [1, 104]
	var ans int

	var queue = []*TreeNode{root}

	for len(queue) > 0 {
		var val = make([]int, len(queue))
		for i := range val {
			node := queue[0]
			queue = queue[1:]
			val[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		ans = val[0]
	}

	return ans
}

// 方案二：从右到左，所以右节点先入队,最后出队就是答案
func findBottomLeftValue2(root *TreeNode) int {

	node := root // 最后出队
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		//从右到左
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)

		}
	}
	return node.Val

}
