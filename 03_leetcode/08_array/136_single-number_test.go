package _8_array

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 只出现一次的数字 https://leetcode.cn/problems/single-number/?envType=study-plan-v2&envId=top-100-liked
func TestSingleNumber(t *testing.T) {
	convey.Convey("只出现一次的数字", t, func() {
		testCase := []struct {
			input  []int
			target int
		}{
			{

				[]int{2, 2, 1}, 1,
			},

			{

				[]int{4, 1, 2, 1, 2}, 4,
			},

			{

				[]int{1}, 1,
			},
		}

		for _, tst := range testCase {
			rsp := singleNumber(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func singleNumber(nums []int) int {
	/*
		1. 任何数和 0 做异或运算，结果仍然是原来的数，即 a⊕0=a。
		2. 任何数和其自身做异或运算，结果是 0，即 a⊕a=0。
		3. 异或运算满足交换律和结合律，即 a⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=b。

	*/
	// 题目: 除了某个元素只出现一次以外，其余每个元素均出现两次

	var ans int
	for _, num := range nums {
		ans ^= num
	}
	return ans
}
