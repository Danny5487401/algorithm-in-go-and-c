package _3_greedy_method

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 跳跃游戏 II https://leetcode.cn/problems/jump-game-ii/?envType=study-plan-v2&envId=top-100-liked
func TestJump(t *testing.T) {
	convey.Convey("跳跃游戏II:最初位于数组的第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。判断你是否能够到达最后一个下标，返回到达 nums[n - 1] 的最小跳跃次数", t, func() {
		testCase := []struct {
			input  []int
			target int
		}{
			{

				[]int{2, 3, 1, 1, 4}, 2,
			},
			{

				[]int{2, 3, 0, 1, 4}, 2,
			},
		}

		for _, tst := range testCase {
			rsp := jump(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func jump(nums []int) (ans int) {
	curRight := 0  // 已建造的桥的右端点
	nextRight := 0 // 下一座桥的右端点的最大值
	for i, num := range nums[:len(nums)-1] {
		// 遍历的过程中，记录下一座桥的最远点
		nextRight = max(nextRight, i+num)
		if i == curRight { // 无路可走，必须建桥
			curRight = nextRight // 建桥后，最远可以到达 nextRight
			ans++
		}
	}
	return
}
