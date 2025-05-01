package _6_backtracking

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 24 点游戏 https://leetcode.cn/problems/24-game/description/
func TestJudgePoint24(t *testing.T) {
	convey.Convey("字母大小写全排列", t, func() {
		testCase := []struct {
			input  []int
			target bool
		}{
			{
				[]int{4, 1, 8, 7}, true,
			},
			{

				[]int{1, 2, 1, 2}, false,
			},
		}

		for _, tst := range testCase {
			rsp := judgePoint24(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func judgePoint24(nums []int) bool {
	/*
		游戏的第一步是挑出两个数，算出一个新数替代这两个数。

		然后，在三个数中玩 24 点，再挑出两个数，算出一个数替代它们。

		然后，在两个数中玩 24 点…
	*/
	floatNums := make([]float64, len(nums))
	for i := range floatNums {
		floatNums[i] = float64(nums[i])
	}
	return dfs(floatNums)

}

func dfs(nums []float64) bool {
	if len(nums) == 1 {
		return math.Abs(nums[0]-24) < 1e-9
	}
	flag := false
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			n1, n2 := nums[i], nums[j]
			newNums := make([]float64, 0, len(nums))
			for k := 0; k < len(nums); k++ {
				if k != i && k != j {
					newNums = append(newNums, nums[k]) // 传入长度变小的新数组继续递归
				}
			}
			// 逐个尝试每一种运算
			// 一旦某个递归返回真，flag 就变为真，由于||的短路特性，后面的递归不会执行。
			flag = flag || dfs(append(newNums, n1+n2))
			flag = flag || dfs(append(newNums, n1-n2))
			flag = flag || dfs(append(newNums, n2-n1))
			flag = flag || dfs(append(newNums, n1*n2))
			if n1 != 0 {
				flag = flag || dfs(append(newNums, n2/n1))
			}
			if n2 != 0 {
				flag = flag || dfs(append(newNums, n1/n2))
			}
			if flag {
				return true
			}
		}
	}
	return false
}
