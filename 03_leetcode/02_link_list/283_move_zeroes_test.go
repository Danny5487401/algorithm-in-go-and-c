package _2_link_list

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// 移动零 https://leetcode.cn/problems/move-zeroes/description/?envType=study-plan-v2&envId=top-100-liked
func TestMoveZeros(t *testing.T) {
	convey.Convey("移动零：0 在末尾，其他元素顺序不变 ", t, func() {
		testCase := []struct {
			input    []int
			expected []int
		}{
			{
				[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0},
			},
			{
				[]int{0}, []int{0},
			},
		}

		for _, tst := range testCase {
			moveZeroes(tst.input)
			convey.So(tst.input, convey.ShouldEqual, tst.expected)
		}
	})

}

func moveZeroes(nums []int) {
	// 左指针左边均为非零数
	// 右指针左边直到左指针处均为零
	left, right, n := 0, 0, len(nums)
	for right < n { //
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left] // 左右指针交换
			left++
		}
		right++
	}
}
