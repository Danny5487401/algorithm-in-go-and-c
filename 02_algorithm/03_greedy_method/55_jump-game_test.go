package _3_greedy_method

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 跳跃游戏 https://leetcode.cn/problems/jump-game/description/?envType=study-plan-v2&envId=top-100-liked
func TestCanJump(t *testing.T) {
	convey.Convey("跳跃游戏:最初位于数组的第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。判断你是否能够到达最后一个下标，", t, func() {
		testCase := []struct {
			input  []int
			target bool
		}{
			{

				[]int{2, 3, 1, 1, 4}, true,
			},
			{

				[]int{3, 2, 1, 0, 4}, false,
			},
		}

		for _, tst := range testCase {
			rsp := canJump(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

// 超时写法
func canJumpWithTimeout(nums []int) bool {
	length := len(nums)
	var dfs func(i int) bool
	dfs = func(i int) bool {
		if i == length-1 {
			return true
		}
		if i >= length {
			return false
		}
		if nums[i] == 0 {
			return false
		}
		for step := 1; step <= nums[i]; step++ {
			if dfs(i + step) {
				return true
			}
		}
		return false
	}
	return dfs(0)

}

func canJump(nums []int) bool {
	mx := 0 // 维护最右可以到达的位置 mx
	for i, jumpStep := range nums {
		if i > mx { // 无法到达 i
			return false
		}
		mx = max(mx, i+jumpStep) // 从 i 最右可以跳到 i + jump
		if mx >= len(nums)-1 {   // 可以跳到 n-1
			break
		}
	}
	return true
}
