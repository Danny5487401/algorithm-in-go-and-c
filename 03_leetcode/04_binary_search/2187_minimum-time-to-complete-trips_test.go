package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"slices"
	"testing"
)

// 完成旅途的最少时间 https://leetcode.cn/problems/minimum-time-to-complete-trips/description/
func TestMinimumTime(t *testing.T) {
	convey.Convey("time[i] 表示第 i 辆公交车完成 一趟旅途 所需要花费的时间。给你一个整数 totalTrips ，表示所有公交车 总共 需要完成的旅途数目。"+
		"请你返回完成 至少 totalTrips 趟旅途需要花费的 最少 时间。", t, func() {
		testCase := []struct {
			input  []int
			ans    int
			target int
		}{
			//
			{
				/*
					- 时刻 t = 1 ，每辆公交车完成的旅途数分别为 [1,0,0] 。
					  已完成的总旅途数为 1 + 0 + 0 = 1 。
					- 时刻 t = 2 ，每辆公交车完成的旅途数分别为 [2,1,0] 。
					  已完成的总旅途数为 2 + 1 + 0 = 3 。
					- 时刻 t = 3 ，每辆公交车完成的旅途数分别为 [3,1,1] 。
					  已完成的总旅途数为 3 + 1 + 1 = 5 。
					所以总共完成至少 5 趟旅途的最少时间为 3
				*/
				input:  []int{1, 2, 3},
				ans:    5,
				target: 3,
			},
			{
				input:  []int{2},
				ans:    1,
				target: 2,
			},
			{
				input:  []int{5, 10, 10},
				ans:    9,
				target: 25,
			},
		}

		for _, tst := range testCase {
			rsp := minimumTime2(tst.input, tst.ans)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func minimumTime2(time []int, totalTrips int) int64 {

	/*
		每辆车都用时 x，总共能完成多少趟旅途？能否达到 totalTrips？
		1 <= time[i], totalTrips <= 107
		时间越多，可以完成的旅途也就越多，有单调性.
		将它与 totalTrips 比较，如果比 totalTrips 小，说明二分的答案小了，更新二分区间左端点 left，否则更新二分区间右端点 right
	*/

	left := slices.Min(time)               // 这个时间任何车都没法完成一趟旅途
	right := slices.Min(time) * totalTrips // 让最快的车完成 totalTrips 趟旅途

	/*
		假设所有 time[i] 都等于 min(time)，那么每辆车可以完成的旅途为 k=x/min(time)

		假设所有 time[i] 都等于 max(time)，那么每辆车可以完成的旅途为 k=x/max(time)
	*/

	for left <= right {
		mid := left + (right-left)/2
		sum := 0 // 可以完成的 trip
		for _, t := range time {
			sum += mid / t
		}
		if sum >= totalTrips { // 小于threshold 才能符合答案
			right = mid - 1
		} else {
			left = mid + 1 // 增大时间
		}
	}
	// 这里需要 right + 1
	return int64(right) + 1
}

// 解题思路: https://leetcode.cn/problems/minimum-time-to-complete-trips/solutions/1295955/er-fen-da-an-python-yi-xing-gao-ding-by-xwvs8/
func minimumTime(time []int, totalTrips int) int64 {
	minT := slices.Min(time)
	left := minT - 1           // 循环不变量：check(left) 恒为 false
	right := minT * totalTrips // 循环不变量：check(right) 恒为 true
	for left+1 < right {       // 开区间 (left, right) 不为空
		mid := left + (right-left)/2
		sum := 0
		for _, t := range time {
			sum += mid / t
		}
		if sum >= totalTrips {
			right = mid // 缩小二分区间为 (left, mid)
		} else {
			left = mid // 缩小二分区间为 (mid, right)
		}
	}
	// 此时 left 等于 right-1
	// check(left) = false 且 check(right) = true，所以答案是 right
	return int64(right) // 最小的 true
}
