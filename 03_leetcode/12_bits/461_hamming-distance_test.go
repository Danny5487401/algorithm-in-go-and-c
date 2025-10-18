package _2_bits

import (
	"github.com/smartystreets/goconvey/convey"
	"math/bits"
	"testing"
)

// 汉明距离 https://leetcode.cn/problems/hamming-distance/description/

func TestHammingDistance(t *testing.T) {
	convey.Convey("汉明距离 指的是这两个数字对应二进制位不同的位置的数目", t, func() {
		testCase := []struct {
			x      int
			y      int
			target int
		}{
			{
				/*
					1   (0 0 0 1)
					4   (0 1 0 0)
				*/
				1, 4, 2,
			},
			{
				3, 1, 1,
			},
		}

		for _, tst := range testCase {
			rsp := hammingDistance(tst.x, tst.y)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func hammingDistance(x int, y int) int {
	num := x ^ y
	return bits.OnesCount(uint(num))
}
