package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"slices"
	"testing"
)

//  数组的最大美丽值 https://leetcode.cn/problems/maximum-beauty-of-an-array-after-applying-operation/

func TestMaximumBeauty(t *testing.T) {
	convey.Convey("数组的最大美丽值: 美丽值定义为数组中由相等元素组成的最长子序列的长度,可以 nums[i] 替换为范围 [nums[i] - k, nums[i] + k] 内的任一整数", t, func() {
		testCase := []struct {
			nums     []int
			k        int
			expected int
		}{
			{
				/*
					- 选择下标 1 ，将其替换为 4（从范围 [4,8] 中选出），此时 nums = [4,4,1,2] 。
					- 选择下标 3 ，将其替换为 4（从范围 [0,4] 中选出），此时 nums = [4,4,1,4] 。
					执行上述操作后，数组的美丽值是 3（子序列由下标 0 、1 、3 对应的元素组成）。
					可以证明 3 是我们可以得到的由相等元素组成的最长子序列长度。
				*/
				[]int{4, 6, 1, 2},
				2,
				3,
			},
			{
				// 数组 nums 的美丽值是 4（整个数组）
				[]int{1, 1, 1, 1},
				10,
				4,
			},
		}

		for _, tst := range testCase {
			rsp := maximumBeauty(tst.nums, tst.k)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func maximumBeauty(nums []int, k int) int {
	/*
		由于选的是子序列，且操作后子序列的元素都相等，所以元素顺序对答案没有影响，可以先对数组排序
		考虑最左边的区间 [x−k,x+k] 和最右边的区间 [y−k,y+k] , 需要 交集 x+k≥y−k 即 y−x≤2k
		原问题等价于:排序后，找最长的连续子数组，其最大值减最小值不超过 2k。
		用 滑动窗口 解决。枚举 nums[right] 作为子数组的最后一个数，一旦 nums[right]−nums[left]>2k，就移动左端点 left
	*/
	var ans int

	slices.Sort(nums)
	left := 0

	for right := 0; right < len(nums); right++ {
		for nums[right]-nums[left] > 2*k {
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
