package _8_array

import (
	"slices"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// 加一 https://leetcode.cn/problems/plus-one/description/
func TestPlusOne(t *testing.T) {
	convey.Convey("合并区间", t, func() {
		testCase := []struct {
			input  []int
			target []int
		}{
			{
				[]int{1, 2, 3},
				[]int{1, 2, 4},
			},
			{
				[]int{4, 3, 2, 1},
				[]int{4, 3, 2, 2},
			},

			{
				[]int{9},
				[]int{1, 0},
			},
		}

		for _, tst := range testCase {
			rsp := plusOne(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func plusOne(digits []int) []int {

	/*
		算法：

		1 从右往左找第一个不等于 9 的数，记作 digits[i]。
		2 进位，把 digits[i] 加一，把下标在 [i+1,n−1] 中的数全变成 0。
		3. 特别地，如果所有数都等于 9，那么答案为 [1,0,0,…,0]，其中有 n 个 0。

	*/
	var index = -1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			index = i
			break
		}
	}

	for i := index + 1; i < len(digits); i++ {
		digits[i] = 0
	}
	if index == -1 {
		digits = append([]int{1}, digits...)
	} else {
		digits[index] += 1
	}
	return digits
}

func plusOne2(digits []int) []int {
	// go 1.23 引入slices.Backward
	for i, d := range slices.Backward(digits) {
		if d < 9 {
			digits[i]++ // 进位
			return digits
		}
		digits[i] = 0 // 进位数字的右边数字都变成 0
	}

	// digits 全是 9，加一后变成 100...00
	digits = append(digits, 0)
	digits[0] = 1
	return digits
}
