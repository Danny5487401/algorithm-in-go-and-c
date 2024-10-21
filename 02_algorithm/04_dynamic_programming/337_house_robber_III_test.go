package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 打家劫舍III https://leetcode.cn/problems/house-robber-iii/description/
func TestHouseRobberIII(t *testing.T) {
	convey.Convey("打家劫舍III", t, func() {
		testCase := []struct {
			input  *TreeNode
			target int
		}{
			{

				Ints2TreeNode([]int{3, 2, 3, NULL, 3, NULL, 1}), 7,
			},
			{

				Ints2TreeNode([]int{3, 4, 5, 1, 3, NULL, 1}), 9,
			},
		}

		for _, tst := range testCase {
			rsp := robIII(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func robIII(root *TreeNode) int {

	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (rob int, unRob int) {
		if node == nil { // 递归边界
			return
		}
		robLeft, unRobLeft := dfs(node.Left)                        // 递归左子树
		robRight, unRobRight := dfs(node.Right)                     // 递归右子树
		rob = unRobLeft + unRobRight + node.Val                     // 选
		unRob = max(robLeft, unRobLeft) + max(robRight, unRobRight) // 不选
		return
	}
	return max(dfs(root))

}
