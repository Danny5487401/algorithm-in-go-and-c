package _8_array

// 区域和检索 - 数组不可变 https://leetcode.cn/problems/range-sum-query-immutable/description/
type NumArray []int

func Constructor(nums []int) NumArray {
	s := make(NumArray, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	return s
}

func (s NumArray) SumRange(left, right int) int {
	return s[right+1] - s[left]
}
