package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMaxSubarrayLength(t *testing.T) {
	convey.Convey("最多 K 个重复元素的最长子数组", t, func() {
		testCase := []struct {
			input    []int
			k        int
			expected int
		}{
			{
				[]int{1, 2, 3, 1, 2, 3, 1, 2}, 2, 6,
			},
			{
				[]int{1, 2, 1, 2, 1, 2, 1, 2}, 1, 2,
			},
		}

		for _, tst := range testCase {
			rsp := maxSubarrayLength(tst.input, tst.k)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func maxSubarrayLength(nums []int, k int) int {
	var count = make(map[int]int)

	left := 0
	length := len(nums)
	if length <= 1 {
		return length
	}

	var ans int

	for right, v := range nums {
		count[v]++
		for count[v] > k {
			count[nums[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans

}
