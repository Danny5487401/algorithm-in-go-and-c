package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

// 统计和小于目标的下标对数目 https://leetcode.cn/problems/count-pairs-whose-sum-is-less-than-target/
func TestCountPairs(t *testing.T) {
	convey.Convey("nums[i] + nums[j] < target 的下标对 (i, j) 的数目", t, func() {
		testCase := []struct {
			input1 []int
			target int
			output int
		}{
			{
				[]int{-1, 1, 2, 3, 1},
				2,
				3,
			},
			{
				[]int{-6, 2, 5, -2, -7, -1, 3},
				-2,
				10,
			},
		}

		for _, tst := range testCase {
			rsp := countPairs(tst.input1, tst.target)
			convey.So(rsp, convey.ShouldEqual, tst.output)
		}
	})

}

func countPairs(nums []int, target int) int {
	/*
		排序+双指针
	*/
	sort.Ints(nums)

	left := 0
	right := len(nums) - 1
	var ans int

	for left < right {
		if nums[left]+nums[right] < target {
			// 在 [left+1,right] 中的任何 nums[i] 相加，都是 <target 的
			ans += right - left
			left++
		} else {
			right--
		}
	}
	return ans
}
