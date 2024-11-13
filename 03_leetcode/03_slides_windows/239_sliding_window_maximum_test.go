package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 滑动窗口最大值 https://leetcode.cn/problems/sliding-window-maximum/description/
func TestMaxSlidingWindow(t *testing.T) {
	convey.Convey("滑动窗口最大值：窗口为k内的最大值 ", t, func() {
		testCase := []struct {
			nums     []int
			k        int
			expected []int
		}{
			{
				[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7},
			},
			{
				[]int{1}, 1, []int{1},
			},
		}

		for _, tst := range testCase {
			rsp := maxSlidingWindow(tst.nums, tst.k)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0, len(nums)-k+1) // 预先分配好空间

	var q []int // 记录的是索引
	for i, num := range nums {
		// 1 进入窗口
		for len(q) > 0 && nums[q[len(q)-1]] <= num {
			// 维护 q 的单调性: 当元素大于单调队列的最后一个元素
			q = q[:len(q)-1]
		}
		q = append(q, i)

		// 2 离开窗口
		if i-q[0]+1 > k {
			q = q[1:]
		}

		// 3 记录答案
		if i >= k-1 { // 当 k =3 , i=2 开始记录答案
			// 单调递减，队首就是最大值
			ans = append(ans, nums[q[0]])
		}
	}

	return ans
}
