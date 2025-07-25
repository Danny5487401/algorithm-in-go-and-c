package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 大小为 K 且平均值大于等于阈值的子数组数目 https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/description/
func TestNumOfSubarrays(t *testing.T) {
	convey.Convey("无重复字符的最长子串", t, func() {
		testCase := []struct {
			input     []int
			k         int
			threshold int
			expected  int
		}{
			{
				[]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4, 3,
			},
			{
				[]int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, 3, 5, 6,
			},
		}

		for _, tst := range testCase {
			rsp := numOfSubarrays(tst.input, tst.k, tst.threshold)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func numOfSubarrays(arr []int, k int, threshold int) int {
	var ans int
	var num int
	for i, v := range arr {
		num += v
		if i < k-1 {
			continue
		}
		if num/k >= threshold {
			ans++
		}
		num -= arr[i-k+1]
	}
	return ans
}
