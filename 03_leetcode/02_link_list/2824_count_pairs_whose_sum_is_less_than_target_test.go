package _2_link_list

import "sort"

func countPairs(nums []int, target int) int {
	sort.Ints(nums)

	left := 0
	right := len(nums) - 1
	var ans int

	for left < right {
		if nums[left]+nums[right] < target {
			// 在 [left+1,right] 中的任何 nums[i] 相加，都是 <target 的
			ans += right - left
			left++
		} else {
			right--
		}
	}
	return ans
}
