package _8_array

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 和为 K 的子数组 https://leetcode.cn/problems/subarray-sum-equals-k/description/?envType=study-plan-v2&envId=top-100-liked
func TestSubarraySumy(t *testing.T) {
	convey.Convey("和为 K 的子数组:子数组是数组中元素的连续非空序列", t, func() {
		testCase := []struct {
			input  []int
			k      int
			target int
		}{
			{

				[]int{1, 1, 1}, 2, 2,
			},
			{

				[]int{1, 2, 3}, 3, 2,
			},
		}

		for _, tst := range testCase {
			rsp := subarraySum(tst.input, tst.k)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

/*
不能用滑动窗口:滑动窗口需要满足单调性，当右端点元素进入窗口时，窗口元素和是不能减少的.本题 nums 包含负数，当负数进入窗口时，窗口左端点反而要向左移动，导致算法复杂度不是线性的。

前缀和的定义：s[0]=0, s[i+1]=nums[0]+nums[1]+⋯+nums[i]

s[j]−s[i]=k  (i<j) --> s[i]=s[j]−k

在遍历 s[j] 的同时，用一个哈希表 cnt 统计 s[j] 的个数。那么枚举到 s[j] 时，从哈希表中就可以找到有 cnt[s[j]−k] 个 s[i]，即为元素和等于 k 的子数组个数，加入答案.
每次寻找 s[j] 前面有多少个前缀和等于 s[j]−k

以 nums=[1,1,−1,1,−1], k=1 为例，其前缀和 s=[0,1,2,1,2,1]
j	s[j]	s[j]−k	在 s[j] 左边的 s[j]−k 的个数	解释
0	0		−1		0							无
1	1		0		1							s[0]=0
2	2		1		1							s[1]=1
3	1		0		1							s[0]=0
4	2		1		2							s[1]=s[3]=1
5	1		0		1							s[0]=0

*/

// 两次遍历
func subarraySum(nums []int, k int) (ans int) {
	s := make([]int, len(nums)+1) // 前缀和,包括s[0]=0
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	cnt := map[int]int{} // key 代表目标值 value 代表数目
	for _, sj := range s {
		ans += cnt[sj-k]
		cnt[sj]++
	}
	return
}

// 一次遍历
func subarraySum2(nums []int, k int) (ans int) {
	sum := 0
	// 把 s[0]=0 也加到哈希表: 举个最简单的例子，nums=[1],k=1。如果不把 s[0]=0 加到哈希表中，按照我们的算法，没法算出这里有 1 个符合要求的子数组。
	// 也可以这样理解，要想把任意子数组都表示成两个前缀和的差，必须添加 s[0]=0，否则当子数组是前缀时，没法减去一个数
	cnt := map[int]int{0: 1} // s[0]=0 单独统计
	for _, x := range nums {
		sum += x
		ans += cnt[sum-k]
		cnt[sum]++
	}
	return
}
