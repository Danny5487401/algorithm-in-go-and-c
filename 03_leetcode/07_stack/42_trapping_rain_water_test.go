package _7_stack

// 接雨水

func trap(height []int) int {
	ans := 0
	stack := make([]int, 0)
	for i, curHeight := range height {
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
