package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 子数组最大平均数 I: https://leetcode.cn/problems/maximum-average-subarray-i/description/
func TestFindMaxAverage(t *testing.T) {
	convey.Convey("子数组最大平均数 I", t, func() {
		testCase := []struct {
			nums     []int
			k        int
			expected float64
		}{
			{
				[]int{1, 12, -5, -6, 50, 3}, 4, 12.75,
			},
			{[]int{5}, 1, 5.00000},
		}

		for _, tst := range testCase {
			rsp := findMaxAverage(tst.nums, tst.k)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func findMaxAverage(nums []int, k int) float64 {

	var maxAvgNum float64
	var sum int
	for i, num := range nums {
		// 1. 进入窗口
		sum += num
		if i < k-1 { // 窗口大小不足 k
			continue
		}
		// 2. 更新答案
		maxAvgNum = max(maxAvgNum, float64(sum)/float64(k))
		// 3. 离开窗口
		sum -= nums[i-k+1]
	}
	return maxAvgNum

}
