package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 搜索旋转排序数组 https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
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
			rsp := searchMethod1(tst.input, tst.target)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

/*
核心思想
把某个数 x 与最后一个数 nums[n−1] 比大小：

1 如果 x>nums[n−1]，那么可以推出以下结论：
	* nums 一定被分成左右两个递增段；
	* 第一段的所有元素均大于第二段的所有元素；
	* x 在第一段。
2 如果 x≤nums[n−1]，那么 x 一定在第二段。（或者 nums 就是递增数组，此时只有一段。）


*/

// 方式一： 两次二分,第一次找最小值，第二次找 target
func searchMethod1(nums []int, target int) int {
	// 1 找到 nums 的最小值的下标 i
	// 2 然后分类讨论：
	//	- 如果 target>nums[n−1]，那么 target 一定在第一段 [0,i−1] 中，在 [0,i−1] 中二分查找 target。
	//	- 如果 target≤nums[n−1]，那么：
	//      如果 i=0，说明 nums 是递增的，直接在 [0,n−1] 中二分查找 target。
	//      如果 i>0，那么 target 一定在第二段 [i,n−1] 中，在 [i,n−1] 中二分查找 target。
	//      这两种情况可以合并成：在 [i,n−1] 中二分查找 target。

	i := findMinIndex(nums)
	if target > nums[len(nums)-1] {
		// target 一定在第一段
		return lowBoundWithBoundary(nums, 0, i, target) // 开区间 (-1, i)
	} else {
		// 合并后target 一定在第二段
		return lowBoundWithBoundary(nums, i-1, len(nums), target) // 开区间 (i-1, n)
	}
}

// 有序数组中找 target 的下标
func lowBoundWithBoundary(nums []int, left, right, target int) int {
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	if nums[right] != target {
		return -1
	}

	return right
}

// 方式二：一次二分
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

// 方式三：
func search3(nums []int, target int) int {
	var left, right = 0, len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// 判断mid是在左边递增区间还是右边递增区间
		if nums[left] <= nums[mid] { // 左边区间  [left,mid] 单调递增
			// 判断 target 是否在这区间
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else { // 左边区间  [mid,right] 单调递增
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
