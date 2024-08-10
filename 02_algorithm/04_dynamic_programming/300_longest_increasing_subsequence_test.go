package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 最长递增子序列 https://leetcode.cn/problems/longest-increasing-subsequence/?envType=study-plan-v2&envId=top-100-liked
func TestLengthOfLIS(t *testing.T) {
	convey.Convey("最长递增子序列：不运行相同", t, func() {
		testCase := []struct {
			input  []int
			target int
		}{
			{

				[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4,
			},
			{

				[]int{0, 1, 0, 3, 2, 3}, 4,
			},
			{

				[]int{7, 7, 7, 7, 7, 7, 7}, 1,
			},
		}

		for _, tst := range testCase {
			rsp := lengthOfLIS(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func lengthOfLIS(nums []int) int {
	// 枚举选哪个
	n := len(nums)
	var dfs func(i int) int

	var cache = make([]int, n)

	dfs = func(i int) int {
		if cache[i] != 0 {
			return cache[i]
		}

		res := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				// 这里是 不允许 相同序列，如果可以相同 使用 nums[j] <= nums[i]
				res = max(res, dfs(j))
			}
		}

		res++
		defer func() {
			cache[i] = res
		}()
		return res

	} // dfs[i]= max{dfs[j}+1
	ans := 0

	// 从后面开始
	for i := n - 1; i >= 0; i-- {
		ans = max(ans, dfs(i))
	}
	return ans
}

// 配合使用公共子序列
