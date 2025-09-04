package _8_array

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 变长子数组求和 https://leetcode.cn/problems/sum-of-variable-length-subarrays/description/
func TestSubarraySumVariable(t *testing.T) {
	convey.Convey("返回为数组中每个下标定义的子数组中所有元素的总和。求 nums[start ... i], start = max(0, i - nums[i])", t, func() {
		testCase := []struct {
			input  []int
			target int
		}{
			{
				// [0,0] 2
				// [0,1] 2 + 3 = 5
				// [1.2] 3 + 1 = 4
				[]int{2, 3, 1}, 11,
			},
			{

				[]int{3, 1, 1, 2}, 13,
			},
		}

		for _, tst := range testCase {
			rsp := subarraySumVariable(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func subarraySumVariable(nums []int) int {
	var ans int
	s := make([]int, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	for i, num := range nums {
		ans += s[i+1] - s[max(i-num, 0)]
	}
	return ans

}
