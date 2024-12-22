package _7_stack

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 每日温度 https://leetcode.cn/problems/daily-temperatures/?envType=study-plan-v2&envId=top-100-liked

func TestDailyTemperatures(t *testing.T) {
	convey.Convey("每日温度:下一个更高温度", t, func() {
		testCase := []struct {
			input    []int
			expected []int
		}{
			{
				[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0},
			},
			{
				[]int{30, 40, 50, 60}, []int{1, 1, 1, 0},
			},
			{
				[]int{30, 60, 90}, []int{1, 1, 0},
			},
		}

		for _, tst := range testCase {
			rsp := dailyTemperatures2(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

// 写法一从右到左 ：栈中记录下一个更大元素的「候选项」
func dailyTemperatures(temperatures []int) []int {

	length := len(temperatures)
	ans := make([]int, length) // 记录索引
	stack := make([]int, 0)    // 初始化单调栈记录索引,栈中记录下一个更大元素的「候选项」
	for i := length - 1; i >= 0; i-- {
		cur := temperatures[i]
		for len(stack) > 0 && cur >= temperatures[stack[len(stack)-1]] {
			// 没有找到答案，取出栈顶元素，相同的时候保留左边的数字
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			// 只要栈内有结果，代表栈顶比当前数要大，取栈顶元素
			ans[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i)

	}
	return ans
}

// 写法二从左往右：栈内是还没有找到更大的数
func dailyTemperatures2(temperatures []int) []int {

	length := len(temperatures)
	ans := make([]int, length) // 记录索引
	stack := make([]int, 0)    // 初始化单调栈记录索引, 栈内是还没有找到更大的数
	for i, cur := range temperatures {
		for len(stack) > 0 && cur > temperatures[stack[len(stack)-1]] { // 注意是循环处理
			j := stack[len(stack)-1]     // 取出栈顶元素
			ans[j] = i - j               // 注意 更新的是第j个数据并且是 i-j
			stack = stack[:len(stack)-1] // 答案算出来了,不用再保留元素
		}
		stack = append(stack, i) // 写入还没找到的数据

	}
	return ans
}
