package _4_binary_search

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// 寻找峰值 https://leetcode.cn/problems/find-peak-element/description/
func TestFindPeakElement(t *testing.T) {
	convey.Convey("寻找峰值", t, func() {
		testCase := []struct {
			input    []int
			expected int
		}{
			{
				[]int{1, 2, 3, 1}, 2,
			},
			{
				[]int{1, 2, 1, 3, 5, 6, 4}, 5,
			},
		}

		for _, tst := range testCase {
			rsp := findPeakElement(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

func findPeakElement(nums []int) int {
	// 二分的范围 [0,n-2] --> 开区间 (-1,n-1)
	/*
		- 红色：峰顶左侧元素
		- 蓝色：峰顶及其右侧元素
		- 白色：不确定
		根据此定义，因为 nums >0, 峰值一定存在，那么最右侧为蓝色
	*/
	left := -1
	right := len(nums) - 1
	for left+1 < right {
		mid := left + (right-left)/2 // 溢出处理
		if nums[mid] > nums[mid+1] {
			// 左大 ，说明M为峰顶,要么M在峰顶右侧, ,更新右侧为蓝色
			right = mid
		} else {
			// 右大, 说明左侧位红色
			left = mid
		}
	}
	return right
}
