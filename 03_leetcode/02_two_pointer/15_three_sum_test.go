package main

import (
	"github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	convey.Convey("two sums ", t, func() {
		testCase := []struct {
			input    []int
			expected [][]int
		}{
			{
				[]int{-1, 0, 1, 2, -1, -4}, [][]int{
					[]int{-1, -1, 2},
					[]int{-1, 0, 1},
				},
			},
		}

		for _, tst := range testCase {
			ret := threeSum(tst.input)
			convey.So(ret, convey.ShouldEqual, tst.expected)
		}
	})

}

func threeSum(nums []int) [][]int {

	n := len(nums)
	// 排序满足 a<b<c 不重复,如 -4 -1 -1 0 1 2
	sort.Ints(nums)

	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			// 需要和上一次枚举的数不相同
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]

		// 枚举 b
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				// 同样需要与上次不同
				continue
			}

			// 如果太大就移动右指针
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}

			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}

			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}

		}

	}
	return ans
}
