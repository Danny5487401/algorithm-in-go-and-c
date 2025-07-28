package main

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 与对应负数同时存在的最大正整数 https://leetcode.cn/problems/largest-positive-integer-that-exists-with-its-negative/description/
func TestFindMaxK(t *testing.T) {
	convey.Convey("两数之和", t, func() {
		testCase := []struct {
			input  []int
			target int
		}{
			{
				[]int{-1, 2, -3, 3}, 3,
			},
			{
				[]int{-1, 10, 6, 7, -7, 1}, 7,
			},
			{
				[]int{-10, 8, 6, 7, -2, -3}, -1,
			},
		}

		for _, tst := range testCase {
			ret := findMaxK(tst.input)
			convey.So(ret, convey.ShouldEqual, tst.target)
		}
	})

}
func findMaxK(nums []int) int {
	var cache = map[int]int{}

	var maxNum = -1
	for _, num := range nums {
		if _, ok := cache[0-num]; ok {
			maxNum = max(maxNum, int(math.Abs(float64(num))))
			continue
		}
		cache[num] = num
	}
	return maxNum
}
