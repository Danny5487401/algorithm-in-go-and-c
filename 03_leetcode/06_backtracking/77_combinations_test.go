package _6_backtracking

// 组合 https://leetcode.cn/problems/combinations/description/ 是 子集的特殊情况 https://leetcode.cn/problems/subsets/?envType=study-plan-v2&envId=top-100-liked

func combine(n int, k int) [][]int {
	var ans = make([][]int, 0)

	var dfs func(i int)
	var path = make([]int, 0)
	dfs = func(i int) {

		d := k - len(path) // 还要选 d 个数
		if d == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		// 从后往前容易点
		for j := i; j >= d; j-- {
			path = append(path, j)
			dfs(j - 1)
			// 返回恢复现场
			path = path[:len(path)-1]
		}
	}
	dfs(n)
	return ans
}
