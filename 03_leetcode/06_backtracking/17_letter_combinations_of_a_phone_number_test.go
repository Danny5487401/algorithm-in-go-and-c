package _6_backtracking

// 电话号码的字母组合 https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/?envType=study-plan-v2&envId=top-100-liked
// 数字字母映射
var mapping = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	length := len(digits)
	if length == 0 {
		return []string{}
	}
	ans := make([]string, 0)
	path := make([]byte, length) // 本题 path 长度固定为 length
	var dfs func(i int)
	dfs = func(i int) {
		if i == length {
			ans = append(ans, string(path))
			return
		}
		// 枚举第i个字母
		for _, c := range mapping[digits[i]-'0'] {
			path[i] = byte(c)
			dfs(i + 1)
		}
	}
	dfs(0)
	return ans
}
