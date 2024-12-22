package _7_stack

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 下一个更大元素 I  https://leetcode.cn/problems/next-greater-element-i/description/
func TestNextGreaterElement(t *testing.T) {
	convey.Convey("下一个更大元素 I:"+
		"nums1 是 nums2 的子集,找出满足 nums1[i] == nums2[j] 的下标 j，并且在 nums2 确定 nums2[j] 的 下一个更大元素", t, func() {
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

// 方式一: 从右到左
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// 把 nums1 的每个元素及其下标记录到一个哈希表 num1Index 中
	num1Index := make(map[int]int, len(nums1))
	for i, n := range nums1 {
		num1Index[n] = i
	}

	// 初始化答案
	var ans = make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}

	stack := make([]int, 0)                // 栈内放的是元素,不是索引
	for i := len(nums2) - 1; i >= 0; i-- { // 从右往左倒着遍历，栈中记录下一个更大元素的「候选项」
		cur := nums2[i]
		for len(stack) > 0 && cur >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 { // 有最大值
			if index, ok := num1Index[cur]; ok {
				ans[index] = stack[len(stack)-1]
			}
		}
		stack = append(stack, cur)
	}
	return ans
}

// 方法二：从左到右
func nextGreaterElement2(nums1, nums2 []int) []int {
	idx := make(map[int]int, len(nums1)) // 预分配空间
	for i, x := range nums1 {
		idx[x] = i
	}

	// 初始化答案
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}

	st := []int{}
	for _, x := range nums2 {
		for len(st) > 0 && x > st[len(st)-1] { // 只要遍历到比栈顶元素更大的数，就意味着栈顶元素找到了答案
			// x 是栈顶的下一个更大元素
			// 既然栈顶已经算出答案，弹出
			ans[idx[st[len(st)-1]]] = x // 记录答案
			st = st[:len(st)-1]
		}
		if _, ok := idx[x]; ok { // x 在 nums1 中
			st = append(st, x) // 只需把在 nums1 中的元素入栈
		}
	}
	return ans
}
