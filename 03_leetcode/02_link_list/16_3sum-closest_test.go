package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"sort"
	"testing"
)

// 最接近的三数之和 https://leetcode.cn/problems/3sum-closest/

func TestThreeSumClosest(t *testing.T) {
	convey.Convey("最接近的三数之和", t, func() {
		testCase := []struct {
			input    []int
			target   int
			expected int
		}{
			{
				[]int{-1, 2, 1, -4}, 1, 2,
			},
			{
				[]int{0, 0, 0}, 1, 0,
			},
		}

		for _, tst := range testCase {
			ret := threeSumClosest(tst.input, tst.target)
			convey.So(ret, convey.ShouldEqual, tst.expected)
		}
	})

}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var ans int
	var minDiff = math.MaxInt
	for first := 0; first < len(nums)-2; first++ {
		firstNum := nums[first]

		second := first + 1
		// c 对应的指针初始指向数组的最右端
		length := len(nums)
		third := length - 1
		var sum int

		// 优化一
		s := firstNum + nums[first+1] + nums[first+2]
		if s > target { // 后面无论怎么选，选出的三个数的和不会比 s 还小
			if s-target < minDiff {
				ans = s // 由于下面直接 break，这里无需更新 minDiff
			}
			break
		}

		// 优化二
		s = firstNum + nums[length-2] + nums[length-1]
		if s < target {
			// 如果 s<target，由于数组已经排序，nums[i] 加上后面任意两个数都不超过 s，所以下面的双指针就不需要跑了，无法找到比 s 更优的答案。但是后面还有更大的 nums[i]，可能找到一个离 target 更近的三数之和，所以还需要继续枚举，continue 外层循环。
			if target-s < minDiff {
				minDiff = target - s
				ans = s
			}
			continue
		}

		for second < third {
			sum = nums[second] + nums[third] + firstNum
			if sum == target { // 如果 s=target，那么答案就是 s，直接返回 s
				return sum
			} else if sum < target {
				if target-sum < minDiff {
					ans = sum
					minDiff = target - sum
				}
				second++
			} else if sum > target {
				if sum-target < minDiff {
					ans = sum
					minDiff = sum - target
				}
				third--
			}
		}

	}
	return ans
}
