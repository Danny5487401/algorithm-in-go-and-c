package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 组合总和 Ⅳ https://leetcode.cn/problems/combination-sum-iv/description/
func TestCombinationSum4(t *testing.T) {
	convey.Convey("组合总和 Ⅳ", t, func() {
		testCase := []struct {
			input  []int
			sum    int
			target int
		}{
			{

				[]int{1, 2, 3}, 4, 7,
			},
			{

				[]int{9}, 3, 0,
			},
		}

		for _, tst := range testCase {
			rsp := combinationSum4(tst.input, tst.sum)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func combinationSum4(nums []int, target int) int {
	// 本质是爬楼梯，相当于每次往上爬  nums[i] 步

	// 范围1 <= nums.length <= 200

	cache := make([]int, target+1)
	for i := 0; i <= target; i++ {
		cache[i] = -1
	}

	var dfs func(i int) int
	dfs = func(i int) (res int) {

		if i == 0 {
			return 1
		}
		if i < 0 {
			return 0
		}
		if cache[i] != -1 {
			return cache[i]
		}
		defer func() { cache[i] = res }()
		for n := range nums {
			res += dfs(i - nums[n])
		}
		return

	}
	return dfs(target)
}
