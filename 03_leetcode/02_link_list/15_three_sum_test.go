package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

// 三数之和 https://leetcode.cn/problems/3sum/?envType=study-plan-v2&envId=top-100-liked

func TestThreeSum(t *testing.T) {
	convey.Convey("三数之和为0：未排序，不可以包含重复的三元组", t, func() {
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
			{
				[]int{0, 1, 1}, [][]int{},
			},
			{
				[]int{0, 0, 0}, [][]int{
					[]int{0, 0, 0},
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

	// 可以转成 两数之和 II - 输入有序数组 https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/
	// 也就是 num[i]+num[j] = -num[k]
	// 排序满足 a<b<c 不重复,如 -4 -1 -1 0 1 2
	sort.Ints(nums)

	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			// first 需要从第二个开始对比，不可以有重复的三元组
			continue
		}
		// 优化处1：最小的三个数大于0，那么剩下右边的也大于0
		if nums[first]+nums[first+1]+nums[first+2] > 0 {
			break
		}
		// 优化处2 : 最大的三个数小于0，那么剩下左边的也小于0
		if nums[first]+nums[n-1]+nums[n-2] < 0 {
			continue
		}

		second := first + 1
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		firstNum := nums[first]

		for second < third {
			sum := nums[second] + nums[third] + firstNum
			// 需要保证 b 的指针在 c 的指针的左侧
			if sum > 0 {
				// 如果太大就移动右指针
				third--
			} else if sum < 0 {
				// 如果太小就移动左指针
				second++
			} else {
				// 三元组顺序不重要
				ans = append(ans, []int{nums[first], nums[second], nums[third]})

				// 去重操作： second 和 third 重复也需要跳过
				second++
				for second < third && nums[second] == nums[second-1] {
					second++
				}

				third--
				for second < third && nums[third] == nums[third+1] {
					third--
				}

			}

		}

	}
	return ans
}
