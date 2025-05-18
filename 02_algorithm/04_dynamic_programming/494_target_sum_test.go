package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 目标和 https://leetcode.cn/problems/target-sum/description/
func TestFindTargetSumWays(t *testing.T) {
	convey.Convey("目标和：加减实现 ", t, func() {
		testCase := []struct {
			input struct {
				nums   []int
				target int
			}
			target int
		}{
			{

				struct {
					nums   []int
					target int
				}{
					nums:   []int{1, 1, 1, 1, 1},
					target: 3},
				5,
			},
			{

				struct {
					nums   []int
					target int
				}{
					nums:   []int{1},
					target: 1},
				1,
			},
		}

		for _, tst := range testCase {
			rsp := findTargetSumWays2(tst.input.nums, tst.input.target)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 递归搜索 + 保存计算结果 = 记忆化搜索
func findTargetSumWays(nums []int, target int) int {
	// p 添加所有正数的和
	// s 所有元素的和
	// s-p 所有负数的和
	// 目标和 t = p - (s-p)= 2p-s -->p=(s+t)/2 含义 从nums 中选择一些数字==一个数（方案数量），并且和是偶数,并且非负数

	// 计算 s+t
	for _, num := range nums {
		target += num
	}
	if target < 0 || target%2 == 1 {
		return 0
	}

	// 计算 (s+t)/2
	target = target / 2
	length := len(nums)

	// 缓存: 不为负数， i，c 两个变量，所以是两层数组
	cache := make([][]int, length) // 外层是 i
	for i := range cache {
		cache[i] = make([]int, target+1) // target+0所以要加一
		for j := range cache[i] {
			cache[i][j] = -1 // -1 表示没用访问过
		}
	}

	var dfs func(i int, c int) int
	dfs = func(i int, c int) (res int) {
		// 从后往前
		if i < 0 {
			if c == 0 {
				// 代表找到了一个方案
				return 1
			}
			return 0
		}
		if cache[i][c] != -1 {
			return cache[i][c]
		}
		defer func() {
			cache[i][c] = res
		}()
		if c < nums[i] { // 代表不能选
			return dfs(i-1, c)
		}
		return dfs(i-1, c) + dfs(i-1, c-nums[i]) // 不选和选

	}
	return dfs(length-1, target)

}

// 记忆化搜索转成递推
func findTargetSumWays2(nums []int, target int) int {
	// 计算 s+t
	for _, num := range nums {
		target += num
	}
	if target < 0 || target%2 == 1 {
		return 0
	}
	target /= 2
	length := len(nums)

	// 二维数组  f[i+1][c] = f[i][c] + f[i][c-x]

	f := make([][]int, length+1) // 防止负数
	for i := range f {
		f[i] = make([]int, target+1) // 防止负数
	}
	f[0][0] = 1
	for i, x := range nums {
		for c := 0; c <= target; c++ {
			if c < x {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = f[i][c] + f[i][c-x]
			}
		}
	}
	return f[length][target]
}

// 推荐实现：https://leetcode.cn/problems/target-sum/solutions/2119041/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-s1cx/
