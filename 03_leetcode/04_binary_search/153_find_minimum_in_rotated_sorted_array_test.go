package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 寻找旋转排序数组中的最小值 https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/
func TestFindMin(t *testing.T) {
	convey.Convey("寻找旋转排序数组中的最小值：原先升序排列 ", t, func() {
		testCase := []struct {
			input    []int
			expected int
		}{
			{

				[]int{3, 4, 5, 1, 2}, 1,
			},
			{
				[]int{4, 5, 6, 7, 0, 1, 2}, 0,
			},
			{
				[]int{11, 13, 15, 17}, 11,
			},
		}

		for _, tst := range testCase {
			rsp := findMin(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

func findMin(nums []int) int {
	// 二分的范围 [0,n-2] --> 开区间 (-1,n-1)
	var left = -1
	var right = len(nums) - 1

	for left+1 < right {
		mid := left + (right-left)/2
		x := nums[mid]
		if x < nums[len(nums)-1] {
			/*
				如果 x>nums[n−1]，那么可以推出以下结论：
				- nums 一定被分成左右两个递增段；
				- 第一段的所有元素均大于第二段的所有元素；
				- x 在第一段。
				- 最小值在第二段。
				- 所以 x 一定在最小值的左边
				如果 x≤nums[n−1]，那么 x 一定在第二段。（或者 nums 就是递增数组，此时只有一段。）
				- x 要么是最小值，要么在最小值右边
			*/
			// 若小于右边界数，则表示当前中间位置属于旋转之前的前半部分，即表明最小值在中间位置左侧，去掉右侧一半数字，更新右边界的值为当前的中间位置
			right = mid
		} else {
			//若中间数大于右边界数，表示当前中间位置的数旋转后索引变小，也就是说最小值应该在右侧，去掉左侧一半数字，更新左边界的值为当前的中间位置 + 1；
			left = mid
		}

	}

	return nums[right]
}
