package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 使结果不超过阈值的最小除数 https://leetcode.cn/problems/find-the-smallest-divisor-given-a-threshold/description/
func TestSmallestDivisor(t *testing.T) {
	convey.Convey("使结果不超过阈值的最小除数", t, func() {
		testCase := []struct {
			input    []int
			target   int
			expected int
		}{
			{
				[]int{1, 2, 5, 9}, 6, 5,
			},

			{
				[]int{44, 22, 33, 11, 1}, 5, 44,
			},
			{
				[]int{2, 3, 5, 7, 11}, 11, 3,
			},
		}

		for _, tst := range testCase {
			rsp := smallestDivisor(tst.input, tst.target)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

func smallestDivisor(nums []int, threshold int) int {
	// 二分查找的下限是 1，这是最小的正整数。而二分查找的上限可以设置为数组 nums 中的最大值 M，
	// 这是因为当除数 d >= M 时，数组 nums 中的每个数除以 d 的结果均为 1，total 的值恒等于数组 nums

	left, right := 1, getMax(nums) // 范围 1 <= nums[i] <= 10^6 , nums.length <= threshold <= 106
	var ans int
	for left <= right {
		mid := (left + right) / 2
		var total float64
		for _, num := range nums {
			total += math.Ceil(float64(num) / float64(mid))
		}
		if total <= float64(threshold) { // 小于threshold 才能符合答案
			right = mid - 1
			ans = mid
		} else { // 说明要增大除数
			left = mid + 1

		}
	}
	return ans
}

func getMax(nums []int) int {
	var maxNum int
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}
