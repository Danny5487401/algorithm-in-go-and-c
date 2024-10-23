package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 寻找旋转排序数组中的最小值 https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/
func TestFindMinimumInRotatedSortedAarrayII(t *testing.T) {
	convey.Convey("寻找旋转排序数组中的最小值 II ", t, func() {
		testCase := []struct {
			input    []int
			expected int
		}{
			{

				[]int{1, 5, 3}, 1,
			},
			{
				[]int{2, 2, 2, 0, 1}, 0,
			},
		}

		for _, tst := range testCase {
			rsp := findMinII(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

func findMinII(nums []int) int {
	return 0
}
