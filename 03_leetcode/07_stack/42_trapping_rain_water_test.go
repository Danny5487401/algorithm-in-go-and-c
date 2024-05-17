package _7_stack

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 接雨水 https://leetcode.cn/problems/trapping-rain-water/description/

func TestTrap(t *testing.T) {
	convey.Convey("dailyTemperatures ", t, func() {
		testCase := []struct {
			input    []int
			expected int
		}{
			{
				[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6,
			},
			{
				[]int{4, 2, 0, 3, 2, 5}, 9,
			},
		}

		for _, tst := range testCase {
			rsp := trap1(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

// 方式一：1* (min(preMax,suffixMax)-height)
func trap1(height []int) int {
	ans := 0
	length := len(height)
	preMax := make([]int, len(height)) // preMax[i] 表示从 height[0] 到 height[i] 的最大值
	preMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		preMax[i] = max(preMax[i-1], height[i])
	}

	suffixMax := make([]int, len(height)) //  suffixMax[i] 表示从 height[i] 到 height[n-1] 的最大值
	suffixMax[length-1] = height[length-1]
	for i := len(height) - 2; i >= 0; i-- {
		suffixMax[i] = max(suffixMax[i+1], height[i])
	}

	for i := 0; i < len(height); i++ {
		ans += min(preMax[i], suffixMax[i]) - height[i]
	}
	return ans

}

// 方式二：在方式一的前提下优化空间 前缀最大值比右缀最大值小，那么左边木桶容量就是前缀最大值，同理右缀最大值比前缀最大值小，向右扩展，那么右边边木桶容量就是后缀最大值，向左扩展
func trap2(height []int) int {
	left := 0
	right := len(height) - 1
	var ans = 0
	preMax := 0
	suffixMax := 0
	for left <= right {
		preMax = max(preMax, height[left])
		suffixMax = max(suffixMax, height[right])
		if preMax < suffixMax {
			ans += preMax - height[left]
			left++
		} else {
			ans += suffixMax - height[right]
			right--
		}
	}
	return ans
}

func trap3(height []int) int {
	ans := 0
	stack := make([]int, 0)
	for i, curHeight := range height {
		// 假设宽度为1，面积取左右两边最小值
		for len(stack) > 0 && curHeight >= height[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			bottomHeight := height[j] // 取出第一个高度
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1] // 需要第二个高度
			dh := min(height[left], curHeight) - bottomHeight
			ans += dh * (i - left - 1)
		}
		stack = append(stack, i) // 不大于栈顶都放入
	}
	return ans
}
