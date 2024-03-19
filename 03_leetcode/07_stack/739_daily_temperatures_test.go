package _7_stack

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 每日温度 https://leetcode.cn/problems/daily-temperatures/?envType=study-plan-v2&envId=top-100-liked

func TestDailyTemperatures(t *testing.T) {
	convey.Convey("dailyTemperatures ", t, func() {
		testCase := []struct {
			input    []int
			expected []int
		}{
			{
				[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0},
			},
		}

		for _, tst := range testCase {
			rsp := dailyTemperatures(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func dailyTemperatures(temperatures []int) []int {
	//写法一从右到左 ：栈中记录下一个更大元素的「候选项」

	length := len(temperatures)
	ans := make([]int, length) // 记录的索引
	stack := make([]int, 0)    // 记录的索引
	for i := length - 1; i >= 0; i-- {
		cur := temperatures[i]
		for len(stack) > 0 && cur >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			// 只要栈内有结果，代表栈顶比当前数要大
			ans[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i)

	}
	return ans
}
