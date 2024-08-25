package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

// 在排序数组中查找元素的第一个和最后一个位置 https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/?envType=study-plan-v2&envId=top-100-liked

func TestSearchRange(t *testing.T) {
	convey.Convey("在排序数组中查找元素的第一个和最后一个位置 ", t, func() {
		testCase := []struct {
			input    []int
			target   int
			expected []int
		}{
			{
				[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4},
			},
			{
				[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1},
			},
			{
				[]int{}, 0, []int{-1, -1},
			},
		}

		for _, tst := range testCase {
			rsp := searchRange(tst.input, tst.target)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

// 方式一： 使用内置库 sort.SearchInts
func searchRange1(nums []int, target int) []int {
	// 返回： 不存在是返回可以插入的位置, 存在返回元素索引
	leftmost := sort.SearchInts(nums, target) // 这里是 >=
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

// 方式二：自己实现 二分法
func searchRange(nums []int, target int) []int {
	// 求大于等于 和 小于等于
	start := lowerBound(nums, target)
	if start == len(nums) || nums[start] != target { // 边界判断 , 注意  nums[start] != target 判断，比如 [5,7,7,8,8,10] 找 6
		return []int{-1, -1}
	}
	// 小于等于转换 >=
	end := lowerBound3(nums, target+1) - 1 // 注意最后一个位置查找是 target - 1
	return []int{start, end}

}

// 左端点 方式一 [left,right]
func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1 // 闭区间 [left,right]
	for left <= right {
		// 循环不变量：
		// nums[left-1] < target
		// nums[right+1] >= target
		mid := left + (right-left)/2 // 溢出处理，注意这里是下取整
		if nums[mid] < target {      // L 左边为 小于8的
			left = mid + 1 // 闭区间 [mid+1,right]
		} else {
			right = mid - 1 // 闭区间 [left,mid-1]
		}
	}
	return left // right 右边是 》=target，即 right+1 = left
}

// 左端点 方式二 [left,right)
func lowerBound2(nums []int, target int) int {
	left, right := 0, len(nums) // 左闭右开 [left,right)
	for left < right {
		mid := left + (right-left)/2 // 溢出处理
		if nums[mid] < target {
			left = mid + 1 // 闭区间 [mid+1,right)
		} else {
			right = mid // 闭区间 [left,mid)
		}
	}
	return left
}

// 左端点 方式三 (left,right)
func lowerBound3(nums []int, target int) int {
	left, right := -1, len(nums) // 左开右开 (left,right)
	for left+1 < right {
		mid := (right-left)/2 + left // 溢出处理
		if nums[mid] < target {
			left = mid // 闭区间 (mid,right)
		} else {
			right = mid // 闭区间 (left,mid)
		}
	}
	return right
}
