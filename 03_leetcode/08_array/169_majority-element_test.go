package _8_array

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 多数元素 https://leetcode.cn/problems/majority-element/?envType=study-plan-v2&envId=top-100-liked
func TestMajorityElement(t *testing.T) {
	convey.Convey("多数元素:多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。 ", t, func() {
		testCase := []struct {
			input  []int
			target int
		}{
			{

				[]int{3, 2, 3}, 3,
			},

			{

				[]int{2, 2, 1, 1, 1, 2, 2}, 2,
			},
		}

		for _, tst := range testCase {
			rsp := majorityElement(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

// 方法一：哈希表
func majorityElement(nums []int) int {

	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	maxVal := len(nums) / 2
	var ans int
	for k, v := range count {
		if v > maxVal {
			maxVal = v
			ans = k
		}
	}
	return ans
}

/*
// 方法二：摩尔投票法（Boyer–Moore majority vote algorithm），也被称作「多数投票法」，算法解决的问题是：如何在任意多的候选人中（选票无序），选出获得票数最多的那个

算法可以分为两个阶段：

对抗阶段：分属两个候选人的票数进行两两对抗抵消
计数阶段：计算对抗结果中最后留下的候选人票数是否有效
*/

func majorityElement2(nums []int) int {
	/*
		我们遍历投票数组，将当前票数最多的候选人与其获得的（抵消后）票数分别存储在 major 与 count 中。

		当我们遍历下一个选票时，判断当前 count 是否为零：

		若 count == 0，代表当前 major 空缺，直接将当前候选人赋值给 major，并令 count++
		若 count != 0，代表当前 major 的票数未被完全抵消，因此令 count--，即使用当前候选人的票数抵消 major 的票数

	*/

	major := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count += 1
		} else {
			count -= 1
		}
	}

	return major
}
