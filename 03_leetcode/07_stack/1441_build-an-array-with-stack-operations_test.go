package _7_stack

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 用栈操作构建数组 https://leetcode.cn/problems/build-an-array-with-stack-operations/
func TestBuildArray(t *testing.T) {
	convey.Convey("用栈操作构建数组 ", t, func() {
		testCase := []struct {
			input    []int
			n        int
			expected []string
		}{
			{
				[]int{1, 3}, 3, []string{"Push", "Push", "Pop", "Push"},
			},
			{
				[]int{1, 2, 3}, 3, []string{"Push", "Push", "Push"},
			},
			{
				[]int{1, 2}, 4, []string{"Push", "Push"},
			},
		}

		for _, tst := range testCase {
			rsp := buildArray(tst.input, tst.n)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func buildArray(target []int, n int) []string {
	var push, pop = "Push", "Pop"
	var ans []string
	var pre = 1
	for j := 0; j < len(target); j++ {
		num := target[j]
		for i := pre; i < num; i++ {
			ans = append(ans, push, pop) // 如果不在 target 中，可以先将其 Push 入栈，紧接着 Pop 出栈。
		}
		ans = append(ans, push) // 如果它在 target 中，则只需要将它 Push 入栈即可。
		pre = num + 1
	}

	return ans

}
