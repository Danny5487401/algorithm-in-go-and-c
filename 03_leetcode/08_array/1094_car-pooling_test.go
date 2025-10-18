package _8_array

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 拼车 https://leetcode.cn/problems/car-pooling/description/
func TestCarPooling(t *testing.T) {
	convey.Convey("当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false", t, func() {
		// trips[i] = [numPassengers(i), from(i), to(i)] 表示第 i 次旅行有 numPassengers(i) 乘客，接他们和放他们的位置分别是 from(i) 和 to(i)
		testCase := []struct {
			input    [][]int
			capacity int
			target   bool
		}{
			{

				[][]int{
					{2, 1, 5},
					{3, 3, 7},
				},
				4,
				false,
			},

			{

				[][]int{
					{2, 1, 5},
					{3, 3, 7},
				},
				5,
				true,
			},
		}

		for _, tst := range testCase {
			rsp := carPooling(tst.input, tst.capacity)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func carPooling(trips [][]int, capacity int) bool {

	// 1 <= trips.length <= 1000
	var differenceArray = [1001]int{} // 保证差分数组不会下标越界
	for _, t := range trips {
		differenceArray[t[1]] += t[0]
		differenceArray[t[2]] -= t[0]
	}
	var sum = 0
	for _, pas := range differenceArray {
		sum += pas
		if pas > capacity {
			return false
		}
	}
	return true
}
