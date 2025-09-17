package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 半径为 k 的子数组平均值 https://leetcode.cn/problems/k-radius-subarray-averages/
func TestGetAverages(t *testing.T) {
	convey.Convey("x 个元素的 平均值 是 x 个元素相加之和除以 x ，此时使用截断式 整数除法 ，即需要去掉结果的小数部分。", t, func() {
		testCase := []struct {
			input    []int
			k        int
			expected []int
		}{
			{
				[]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3, []int{-1, -1, -1, 5, 4, 4, -1, -1, -1},
			},
			{
				[]int{100000}, 0, []int{100000},
			},
			{
				[]int{8}, 100000, []int{-1},
			},
		}

		for _, tst := range testCase {
			rsp := getAverages(tst.input, tst.k)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func getAverages(nums []int, k int) []int {
	// 相当于一个长为 2k+1 的滑动窗口
	var sum int
	var ans = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = -1
	}
	for i, num := range nums {
		sum += num
		if i < 2*k {
			continue
		}
		ans[i-k] = sum / (2*k + 1)

		sum -= nums[i-2*k]

	}
	return ans
}
