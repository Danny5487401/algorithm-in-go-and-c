package _8_array

import (
	"github.com/smartystreets/goconvey/convey"
	"slices"
	"testing"
)

// 轮转数组 https://leetcode.cn/problems/rotate-array/description/?envType=study-plan-v2&envId=top-100-liked
func TestRotate(t *testing.T) {
	convey.Convey("轮转数组:将数组中的元素向右轮转 k 个位置", t, func() {
		testCase := []struct {
			input  []int
			k      int
			target []int
		}{
			{

				[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4},
			},

			{

				[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100},
			},
		}

		for _, tst := range testCase {
			rotate(tst.input, tst.k)
			convey.So(tst.input, convey.ShouldResemble, tst.target)
		}
	})

}

func rotate(nums []int, k int) {
	k %= len(nums) // 轮转 k 次等于轮转 k%n 次
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}
