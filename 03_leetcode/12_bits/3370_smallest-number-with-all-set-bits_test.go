package _2_bits

import (
	"github.com/smartystreets/goconvey/convey"
	"math/bits"
	"testing"
)

// 仅含 置位 位的最小整数 https://leetcode.cn/problems/smallest-number-with-all-set-bits/description/
func TestSmallestNumber(t *testing.T) {
	convey.Convey("返回 大于等于 n 且二进制表示仅包含 置位 位的 最小 整数 x.置位 位指的是二进制表示中值为 1 的位。", t, func() {
		testCase := []struct {
			input1 int
			target int
		}{
			{
				// 111
				5, 7,
			},

			{
				// 1111
				10, 15,
			},
			{

				// 11
				3, 3,
			},
		}

		for _, tst := range testCase {
			rsp := smallestNumber2(tst.input1)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 求长度
func smallestNumber(n int) int {
	// 101 --> 111 长度相同
	// 计算 n 的二进制长度 m，返回长为 m 的全为 1 的二进制数
	return 1<<bits.Len(uint(n)) - 1
}

// 枚举法
func smallestNumber2(n int) int {
	i := 1
	for {
		x := 1<<i - 1
		if x >= n {
			return x
		}
		i++
	}
}
