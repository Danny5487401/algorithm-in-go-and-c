package _8_array

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 与车相交的点 https://leetcode.cn/problems/points-that-intersect-with-cars/description/

func TestNumberOfPoints(t *testing.T) {
	convey.Convey("与车相交的点", t, func() {
		testCase := []struct {
			input  [][]int
			target int
		}{
			{

				// 把 a 中下标在 [3,6] 的元素都加一，下标在 [1,5] 的元素都加一，下标在 [4,7] 的元素都加一
				[][]int{
					{3, 6}, {1, 5}, {4, 7},
				}, 7,
			},
			{
				[][]int{
					{1, 3}, {5, 8},
				}, 7,
			},
		}

		for _, tst := range testCase {
			rsp := numberOfPoints(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func numberOfPoints(nums [][]int) int {
	//// 最大结束位置
	//maxEnd := 0
	//for _, num := range nums {
	//	maxEnd = max(maxEnd, num[1])
	//}

	var differenceArray [10001]int
	for _, num := range nums {
		differenceArray[num[0]]++
		differenceArray[num[1]+1]--
	}
	var ans int
	var sum int
	for _, d := range differenceArray {
		sum += d
		if sum > 0 {
			ans++
		}
	}
	return ans
}
