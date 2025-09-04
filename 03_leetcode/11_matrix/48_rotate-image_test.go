package _1_matrix

import (
	"github.com/smartystreets/goconvey/convey"
	"slices"
	"testing"
)

// 旋转图像 https://leetcode.cn/problems/rotate-image/description/
func TestRotate(t *testing.T) {
	convey.Convey("旋转图像", t, func() {
		testCase := []struct {
			input  [][]int
			output [][]int
		}{
			{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				[][]int{
					{7, 4, 1},
					{8, 5, 2},
					{9, 6, 3},
				},
			},
		}

		for _, tst := range testCase {
			rotate(tst.input)
			convey.So(tst.input, convey.ShouldEqual, tst.output)
		}
	})

}

/*
竖着看：

第一列的元素去到第一行。
第二列的元素去到第二行。
……
第 j 列的元素去到第 j 行。
横着看：

第一行的元素去到最后一列。
第二行的元素去到倒数第二列。
……
第 i 行的元素去到第 n−1−i 列。（i 从 0 开始）

结论 (i,j)→(j,n−1−i)

这两步操作都可以原地实现。
转置：把矩阵按照主对角线翻转，位于 (i,j) 的元素去到 (j,i)。
行翻转：把每一行翻转，位于 (j,i) 的元素去到 (j,n−1−i)。
*/
func rotate(matrix [][]int) {
	n := len(matrix)
	// 第一步：转置
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ { // 遍历对角线下方元素
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 第二步：行翻转
	for _, row := range matrix {
		slices.Reverse(row)
	}
}
