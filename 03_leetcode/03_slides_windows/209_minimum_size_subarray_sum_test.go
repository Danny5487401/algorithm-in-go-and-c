package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 长度最小的子数组 https://leetcode.cn/problems/minimum-size-subarray-sum/
func TestMinSubArrayLen(t *testing.T) {
	convey.Convey("长度最小的子数组", t, func() {
		testCase := []struct {
			nums   []int
			target int

			expected int
		}{
			{
				[]int{2, 3, 1, 2, 4, 3}, 7, 2,
			},
			{
				[]int{1, 4, 4}, 4, 1,
			},
			{
				[]int{1, 1, 1, 1, 1, 1, 1, 1}, 11, 0,
			},
		}

		for _, tst := range testCase {
			rsp := minSubArrayLen(tst.target, tst.nums)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func minSubArrayLen(target int, nums []int) int {
	// 注意这里所有数都是正数
	length := len(nums)

	// 题目答案最多为长度，所以初始化答案大于他即可，因为比较更小的值
	ans := math.MaxInt64
	sum := 0
	left := 0

	for right, num := range nums {
		sum += num
		for sum-nums[left] >= target { // 这里没有判断左端点大于右端点，因为target >0
			sum -= nums[left]
			left++
		}
		if sum >= target {
			ans = min(ans, right-left+1)
		}

	}
	if ans <= length {
		return ans
	}

	return 0
}
