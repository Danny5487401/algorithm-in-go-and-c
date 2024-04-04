package _4_dynamic_programming

// 最长递增子序列 https://leetcode.cn/problems/longest-increasing-subsequence/?envType=study-plan-v2&envId=top-100-liked

func lengthOfLIS(nums []int) int {
	// 枚举选哪个
	n := len(nums)
	var dfs func(i int) int

	var cache = make([]int, n)

	dfs = func(i int) int {
		if cache[i] != 0 {
			return cache[i]
		}
		res := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				res = max(res, dfs(j))
			}
		}

		res++
		defer func() {
			cache[i] = res
		}()
		return res

	} // dfs[i]= max{dfs[j}+1
	ans := 0
	for i := n - 1; i >= 0; i-- {
		ans = max(ans, dfs(i))
	}
	return ans
}
