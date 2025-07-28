package main

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 好数对的数目 https://leetcode.cn/problems/number-of-good-pairs/description/
func TestNumIdenticalPairs(t *testing.T) {
	convey.Convey("如果一组数字 (i,j) 满足 nums[i] == nums[j] 且 i < j ，就可以认为这是一组 好数对  ", t, func() {
		testCase := []struct {
			input  []int
			target int
		}{
			{
				// (0,3), (0,4), (3,4), (2,5)
				[]int{1, 2, 3, 1, 1, 3}, 4,
			},
			{
				// 数组中的每组数字都是好数对
				[]int{1, 1, 1, 1}, 6,
			},
			{
				[]int{1, 2, 3}, 0,
			},
		}

		for _, tst := range testCase {
			ret := numIdenticalPairs(tst.input)
			convey.So(ret, convey.ShouldEqual, tst.target)
		}
	})

}

func numIdenticalPairs(nums []int) int {
	// 写个二重循环，枚举 j，统计左边（因为要求 i<j）有多少个数等于 nums[j]。
	var ans int
	cnt := map[int]int{}     //  用一个哈希表 cnt 统计 j 左边的每个数的出现次数。其中 cnt[x] 表示在 j 左边的值为 x 元素个数。
	for _, x := range nums { // x = nums[j]
		// cnt[x] 表示 [0,j-1] 中的 x 的个数
		ans += cnt[x]
		cnt[x]++
		// 现在 cnt[x] 表示 [0,j] 中的 x 的个数
	}
	return ans

}
