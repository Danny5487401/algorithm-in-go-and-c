package _7_stack

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 下一个更大元素 I  https://leetcode.cn/problems/next-greater-element-i/description/

func TestNextGreaterElement(t *testing.T) {
	convey.Convey("下一个更大元素 I:"+
		"nums1 是 nums2 的子集,找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素 ", t, func() {
		testCase := []struct {
			nums1    []int
			nums2    []int
			expected []int
		}{
			{
				[]int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1},
			},
			{
				[]int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1},
			},
		}

		for _, tst := range testCase {
			rsp := nextGreaterElement(tst.nums1, tst.nums2)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// 把 nums1 的每个元素及其下标记录到一个哈希表 num1Index 中
	num1Index := make(map[int]int, len(nums1))
	for i, n := range nums1 {
		num1Index[n] = i
	}
	var ans = make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}

	stack := make([]int, 0) // 栈内放的是元素,不是索引
	for i := len(nums2) - 1; i >= 0; i-- {
		cur := nums2[i]
		for len(stack) > 0 && cur >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			if index, ok := num1Index[cur]; ok {
				ans[index] = stack[len(stack)-1]
			}
		}
		stack = append(stack, cur)
	}
	return ans
}
