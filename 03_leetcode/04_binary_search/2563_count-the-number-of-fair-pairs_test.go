package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

// 统计公平数对的数目 https://leetcode.cn/problems/count-the-number-of-fair-pairs/description/
func TestCountFairPairs(t *testing.T) {
	convey.Convey("lower <= nums[i] + nums[j] <= upper", t, func() {
		/*
			- 中序遍历: [ [左子树的前序遍历结果], 根节点,[右子树的前序遍历结果] ]
		*/
		testCase := []struct {
			nums   []int
			lower  int
			uppper int
			target int
		}{
			{
				/*
					共计 6 个公平数对：(0,3)、(0,4)、(0,5)、(1,3)、(1,4) 和 (1,5)
				*/
				[]int{0, 1, 7, 4, 4, 5},
				3,
				6,
				6,
			},
			{
				/*
					- 第 0 个咒语：3 * [8,5,8] = [24,15,24] 。总共 2 个成功组合。
					- 第 1 个咒语：1 * [8,5,8] = [8,5,8] 。总共 0 个成功组合。
					- 第 2 个咒语：2 * [8,5,8] = [16,10,16] 。总共 2 个成功组合。
				*/
				[]int{1, 7, 9, 2, 5},
				11,
				11,
				1,
			},
		}

		for _, tst := range testCase {
			rsp := countFairPairs(tst.nums, tst.lower, tst.uppper)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 方法一：二分查找
func countFairPairs(nums []int, lower int, upper int) int64 {
	// 排序不影响答案
	sort.Ints(nums)

	var ans int64

	// 枚举右边的 nums[j]， lower−nums[j]≤nums[i]≤upper−nums[j]

	for j, x := range nums {
		// 注意要在 [0, j-1] 中二分，因为题目要求两个下标 i < j

		// 找到 >upper−nums[j] 的第一个数,设其下标为 r，那么下标在 [0,r−1] 中的数都是 ≤upper−nums[j]
		right := sort.SearchInts(nums[:j], upper-x+1) // SearchInts 求的是 >=,  大于-->大于等于x+1

		// 找到 ≥lower−nums[j] 的第一个数，设其下标为 l，那么下标在 [0,l−1] 中的数都是 <lower−nums[j] 的，这有 l 个。
		left := sort.SearchInts(nums[:j], lower-x)

		ans += int64(right - left)
	}

	return ans
}
