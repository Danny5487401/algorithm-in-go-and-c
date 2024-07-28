package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 乘积小于 K 的子数组 https://leetcode.cn/problems/subarray-product-less-than-k/
func TestNumSubarrayProductLessThanK(t *testing.T) {
	convey.Convey("乘积小于 K 的子数组:求个数", t, func() {
		testCase := []struct {
			nums     []int
			target   int
			expected int
		}{
			{
				[]int{10, 5, 2, 6}, 100, 8,
			},
			{
				[]int{1, 2, 3}, 0, 0,
			},
		}

		for _, tst := range testCase {
			rsp := numSubarrayProductLessThanK(tst.nums, tst.target)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func numSubarrayProductLessThanK(nums []int, k int) int {
	// 方法与 209长度最小的子数组 相同

	if k <= 1 {
		// 都是正数，要求小于k
		return 0
	}

	var ans int
	var product = 1 // 乘积
	var left int
	for right, num := range nums {
		product *= num

		for product >= k {
			product /= nums[left]
			// 不断缩小左端点
			left++
		}
		ans += right - left + 1 // 一个元素也算，所以 +1. 如果 [left,right] 乘积小于 k, 那么 [left+1,right]也是小于 k ，所以元素个数 right- left +1
	}
	return ans
}
