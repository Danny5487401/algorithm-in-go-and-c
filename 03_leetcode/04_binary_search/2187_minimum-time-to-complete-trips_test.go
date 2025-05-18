package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 完成旅途的最少时间
func TestMinimumTime(t *testing.T) {
	convey.Convey("time[i] 表示第 i 辆公交车完成 一趟旅途 所需要花费的时间。给你一个整数 totalTrips ，表示所有公交车 总共 需要完成的旅途数目。请你返回完成 至少 totalTrips 趟旅途需要花费的 最少 时间。", t, func() {
		testCase := []struct {
			input  []int
			ans    int
			target int
		}{
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
		}

		for _, tst := range testCase {
			rsp := minimumTime(tst.input, tst.ans)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}
func minimumTime(time []int, totalTrips int) int64 {
	return 0
	//// 1 <= time[i], totalTrips <= 107
	//// 时间越多，可以完成的旅途也就越多，有单调性
	//
	//left := slices.Min(time) - 1           // 这个时间任何车都没法完成一趟旅途
	//right := slices.Max(time) * totalTrips // 让最快的车完成 totalTrips 趟旅途
	//
	///*
	//	假设所有 time[i] 都等于 min(time)，那么每辆车可以完成的旅途为 k=x/min(time)
	//
	//	假设所有 time[i] 都等于 max(time)，那么每辆车可以完成的旅途为 k=x/max(time)
	//*/
	//
	//for left <= right {
	//	mid := left + (right-left)/2
	//	sum := 0
	//	for _, t := range time {
	//		sum += mid / t
	//	}
	//	if sum >= totalTrips { // 小于threshold 才能符合答案
	//		right = mid - 1
	//	} else {
	//		left = mid + 1
	//	}
	//}
	//return int64(right)
}
