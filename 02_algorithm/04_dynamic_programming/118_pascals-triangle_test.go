package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGenerate(t *testing.T) {
	convey.Convey("杨辉三角, ", t, func() {
		testCase := []struct {
			numRows int
			target  [][]int
		}{
			{
				5,
				[][]int{
					{1},
					{1, 1},
					{1, 2, 1},
					{1, 3, 3, 1},
					{1, 4, 6, 4, 1},
				},
			},
		}

		for _, tst := range testCase {
			rsp := generate(tst.numRows)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func generate(numRows int) [][]int {
	/*

		每一排左对齐：
		[1]
		[1,1]
		[1,2,1]
		[1,3,3,1]
		[1,4,6,4,1]

		- 每一排的第一个数和最后一个数都是 1，即 c[i][0]=c[i][i]=1
		- 其余数字等于左上方的数，加上正上方的数，即 c[i][j]=c[i−1][j−1]+c[i−1][j]
	*/
	var ans = make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0], ans[i][i] = 1, 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}

	}
	return ans
}
