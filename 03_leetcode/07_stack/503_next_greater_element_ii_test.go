package _7_stack

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 下一个更大元素 II https://leetcode.cn/problems/next-greater-element-ii/description/

func TestNextGreaterElements(t *testing.T) {
	convey.Convey("下一个更大元素 II: 循环地搜索它的下一个更大的数,nums[nums.length - 1] 的下一个元素是 nums[0] ."+
		"可以把 nums 复制一份，拼在 nums 右边，这样就把环形数组变成一般数组了,代码实现时，无需真的把数组复制一份，而是用下标模 n 的方式取到对应的元素值 ", t, func() {
		testCase := []struct {
			nums     []int
			expected []int
		}{
			{
				[]int{1, 2, 1}, []int{2, -1, 2},
			},
			{
				[]int{1, 2, 3, 4, 3}, []int{2, 3, 4, -1, 4},
			},
		}

		for _, tst := range testCase {
			rsp := nextGreaterElements(tst.nums)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func nextGreaterElements(nums []int) []int {
	var ans = make([]int, len(nums))

	for i := range ans {
		ans[i] = -1
	}
	var stack = make([]int, 0)
	length := len(nums)
	for i := 2*length - 1; i >= 0; i-- {
		cur := nums[i%length]
		for len(stack) > 0 && cur >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if i < length && len(stack) > 0 {
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, cur)
	}

	return ans
}
