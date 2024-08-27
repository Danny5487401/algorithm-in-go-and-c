package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
func TestSearch(t *testing.T) {
	convey.Convey("搜索旋转排序数组：升序 ", t, func() {
		testCase := []struct {
			input    []int
			target   int
			expected int
		}{
			{

				[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4,
			},
			{
				[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1,
			},
			{
				[]int{1}, 0, -1,
			},
		}

		for _, tst := range testCase {
			rsp := search(tst.input, tst.target)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

// 方式一 两次二分：第一次找最小值，第二次找 target

// 方式二：
func search(nums []int, target int) int {
	/*
		设 x=nums[mid] 是我们现在二分取到的数
		如果 x 和 target 在不同的递增段：
			- 如果 target 在第一段（左），x 在第二段（右），说明 x 在 target 右边；
			- 如果 target 在第二段（右），x 在第一段（左），说明 x 在 target 左边。
		如果 x 和 target 在相同的递增段：
			- 比较 x 和 target 的大小即可


	*/
	// 在目标 target 及其右侧
	isBlue := func(i int) bool {
		end := nums[len(nums)-1]
		if nums[i] > end { // 有两段，说明i在前一段

			// 如果 x>nums[n−1]，说明 x 在第一段中，那么 target 也必须在第一段中（否则 x 一定在 target 左边）且 x 必须大于等于 target

			return target > end && nums[i] >= target
		} else {
			// 如果 x≤nums[n−1]，说明 x 在第二段中（或者 nums 只有一段），那么 target 可以在第一段，也可以在第二段。
			// - 如果 target 在第一段，那么 x 一定在 target 右边。
			// - 如果 target 在第二段，那么 x 必须大于等于 target。

			return target > end || nums[i] >= target
		}
	}

	left, right := -1, len(nums)-1 // 开区间 (-1, n-1)
	for left+1 < right {
		mid := (left + right) / 2
		if isBlue(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	if right == len(nums) || nums[right] != target {
		return -1
	}
	return right
}