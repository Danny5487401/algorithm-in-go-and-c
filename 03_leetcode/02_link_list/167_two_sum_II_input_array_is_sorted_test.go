package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 两数之和 II 输入有序数组 https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/
func TestTwoSumII(t *testing.T) {
	convey.Convey("two sums: 有序 ", t, func() {
		testCase := []struct {
			input    []int
			target   int
			expected []int
		}{
			{
				[]int{2, 7, 11, 15}, 9, []int{1, 2},
			},
			{
				[]int{2, 3, 4}, 6, []int{1, 3},
			},
			{
				[]int{-1, 0}, -1, []int{1, 2},
			},
		}

		for _, tst := range testCase {
			ret := twoSum(tst.input, tst.target)
			convey.So(ret, convey.ShouldEqual, tst.expected)
		}
	})

}

func twoSum(numbers []int, target int) []int {
	length := len(numbers)
	first := 0
	var ans = make([]int, 0) // 记录索引下标+1
	last := length - 1
	for first < last {
		sum := numbers[first] + numbers[last]
		if sum == target {
			ans = append(ans, first+1, last+1)
			break
		}
		// 利用数组有序
		if sum < target {
			first++
		} else {
			last--
		}
	}
	return ans

}
