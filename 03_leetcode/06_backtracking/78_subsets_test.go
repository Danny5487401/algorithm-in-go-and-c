package _6_backtracking

// 子集 https://leetcode.cn/problems/subsets/?envType=study-plan-v2&envId=top-100-liked

func subsets(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return [][]int{}
	}
	var ans = make([][]int, 0)
	var path = make([]int, 0, length)
	var dfs func(i int)
	dfs = func(i int) {
		if i == length {
			ans = append(ans, append([]int(nil), path...)) // 固定答案
			return
		}
		// 不选
		dfs(i + 1)
		// 选
		path = append(path, nums[i])
		dfs(i + 1)

		// 恢复现场
		path = path[:len(path)-1]
	}
	dfs(0)
	return ans
}
