package _4_binary_search

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// 寻找峰值II https://leetcode.cn/problems/find-a-peak-element-ii/description/
func TestFindPeakElementII(t *testing.T) {
	convey.Convey("寻找峰值II:  严格大于其相邻格子(上、下、左、右)的元素,相邻数字不同 ", t, func() {
		testCase := []struct {
			input    [][]int
			expected []int
		}{
			{
				[][]int{
					{10, 20, 15},
					{21, 30, 14},
					{7, 16, 32},
				}, []int{1, 1},
			},
			{
				[][]int{
					{1, 4},
					{3, 2},
				},
				[]int{0, 1},
			},
		}

		for _, tst := range testCase {
			rsp := findPeakGrid(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

func findPeakGrid(mat [][]int) []int {
	/*
		二分包含峰顶的行号 i：

		如果 mat[i] 的最大值比它下面的相邻数字小，则存在一个峰顶，其行号大于 i。缩小二分范围，更新二分区间左端点 left。
		如果 mat[i] 的最大值比它下面的相邻数字大，则存在一个峰顶，其行号小于或等于 i。缩小二分范围，更新二分区间右端点 right。

	*/

	// 范围:1 <= m, n <= 500
	left, right := -1, len(mat)-1 // 注意这里是 right= len(mat)-1, 因为后面会进行 +1 判断
	for left+1 < right {
		mid := left + (right-left)/2
		j := indexOfMax(mat[mid])
		if mat[mid][j] > mat[mid+1][j] { // 相邻对比
			right = mid
		} else {
			left = mid
		}
	}
	return []int{right, indexOfMax(mat[right])}

}

func indexOfMax(a []int) (index int) {
	for i, v := range a {
		if v > a[index] {
			index = i
		}
	}
	return
}
