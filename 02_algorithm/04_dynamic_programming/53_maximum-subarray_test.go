package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 最大子数组和 https://leetcode.cn/problems/maximum-subarray/description/?envType=study-plan-v2&envId=top-100-liked
func TestMaxSubArray(t *testing.T) {
	convey.Convey("最大子数组和:子数组的元素和等于两个前缀和的差", t, func() {
		testCase := []struct {
			input  []int
			target int
		}{
			{

				[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6,
			},
			{

				[]int{1}, 1,
			},
			{

				[]int{5, 4, -1, 7, 8}, 23,
			},
		}

		for _, tst := range testCase {
			rsp := maxSubArray(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

/*
s[i+1]=s[i]+nums[i]

子数组的元素和等于两个前缀和的差，所以求出 nums 的前缀和，问题就变成 121. 买卖股票的最佳时机

一边遍历数组计算前缀和，一边维护前缀和的最小值（相当于股票最低价格），用当前的前缀和（卖出价格）减去前缀和的最小值（买入价格），就得到了以当前元素结尾的子数组和的最大值（利润），用它来更新答案的最大值（最大利润）。
*/
func maxSubArray(nums []int) int {
	var minSum int
	var sum int

	// 背景: 1 <= nums.length <= 105
	// 要求:子数组最少包含一个元素,
	var ans = nums[0]
	for _, num := range nums {
		sum += num                // 当前的前缀和
		curSum := sum - minSum    // 减去前缀和的最小值
		minSum = min(sum, minSum) // 维护前缀和的最小值
		ans = max(ans, curSum)
	}
	return ans
}
