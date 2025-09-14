package _2_bits

import (
	"github.com/smartystreets/goconvey/convey"
	"math/bits"
	"testing"
)

// 使两个整数相等的位更改次数 https://leetcode.cn/problems/number-of-bit-changes-to-make-two-integers-equal/description/
func TestMinChanges(t *testing.T) {
	convey.Convey("可以选择 n 的 二进制表示 中任意一个值为 1 的位，并将其改为 0", t, func() {
		testCase := []struct {
			n      int
			k      int
			target int
		}{
			{
				// n = (1101)2
				// k = (0100)2
				// 改变 n 的第一位和第四位
				13, 4, 2,
			},

			{
				21, 21, 0,
			},

			{
				// n= 1110
				// k= 1101
				14, 13, -1,
			},
		}

		for _, tst := range testCase {
			rsp := minChanges(tst.n, tst.k)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func minChanges(n int, k int) int {

	// 是否可以转化判断:判断是否 k 是 n 字集,直接 n & k = k
	if n&k != k {
		return -1
	}

	// 判断需要改的 求异或
	num := n ^ k
	// 计算一个个数
	return bits.OnesCount64(uint64(num))

}
