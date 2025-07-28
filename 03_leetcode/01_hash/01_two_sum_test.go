package main

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 两数之和 https://leetcode.cn/problems/two-sum/description/?envType=study-plan-v2&envId=top-100-liked
func TestTwoSum(t *testing.T) {
	convey.Convey("两数之和 ", t, func() {
		testCase := []struct {
			input    []int
			target   int
			expected []int
		}{
			{
				[]int{2, 7, 11, 15}, 9, []int{0, 1},
			},
			{
				[]int{3, 2, 4}, 6, []int{1, 2},
			},
			{
				[]int{3, 3}, 6, []int{0, 1},
			},
		}

		for _, tst := range testCase {
			ret := twoSum(tst.input, tst.target)
			convey.So(ret, convey.ShouldEqual, tst.expected)
		}
	})

}

// 不推荐行为：两个循环暴力枚举
// 推荐行为：对于 双变量问题，例如两数之和 ai+aj=t，可以枚举右边的 aj，转换成 单变量问题，也就是在 aj左边查找是否有 ai=t−aj，这可以用哈希表维护
func twoSum(nums []int, target int) []int {
	// 存储出现过的结果
	hashTable := map[int]int{} // key为元素，value为索引
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i} // 小的索引在前面
		}
		hashTable[x] = i // 注意这在判断之后
	}
	return nil // 题目保证一定有解，代码不会执行到这里
}
