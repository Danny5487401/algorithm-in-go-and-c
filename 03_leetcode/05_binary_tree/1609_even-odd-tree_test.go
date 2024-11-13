package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func TestIsEvenOddTree(t *testing.T) {
	convey.Convey("奇偶树", t, func() {
		testCase := []struct {
			root   *TreeNode
			target bool
		}{
			{
				Ints2TreeNode([]int{2, 12, 8, 5, 9, NULL, NULL, 18, 16}),
				false,
			},
			{
				Ints2TreeNode([]int{1, 10, 4, 3, NULL, 7, 9, 12, 8, 6, NULL, NULL, 2}),
				true,
			},
		}

		for _, tst := range testCase {
			rsp := isEvenOddTree(tst.root)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func isEvenOddTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	var even = true // 判断偶数层
	for len(queue) > 0 {
		var pre int
		if !even {
			pre = math.MaxInt32 // 初始值不一样
		}
		value := make([]int, len(queue))
		for _ = range value {
			node := queue[0]
			queue = queue[1:]
			if even { // 偶数下标 层上的所有节点的值都是 奇 整数，从左到右按顺序 严格递增
				if node.Val <= pre || node.Val%2 == 0 {
					return false
				}
			} else { // 奇数下标 层上的所有节点的值都是 偶 整数，从左到右按顺序 严格递减
				if node.Val >= pre || node.Val%2 == 1 {
					return false
				}
			}
			pre = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		even = !even
	}
	return true
}
