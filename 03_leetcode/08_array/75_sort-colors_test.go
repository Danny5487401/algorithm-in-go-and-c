package _8_array

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 颜色分类 https://leetcode.cn/problems/sort-colors/?envType=study-plan-v2&envId=top-100-liked
func TestSortColors(t *testing.T) {
	convey.Convey("颜色分类:0、 1 和 2 分别表示红色、白色和蓝色, 原地 对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列 ", t, func() {
		testCase := []struct {
			input  []int
			target []int
		}{
			{

				[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2},
			},

			{

				[]int{2, 0, 1}, []int{0, 1, 2},
			},
		}

		for _, tst := range testCase {
			sortColors(tst.input)
			convey.So(tst.input, convey.ShouldResemble, tst.target)
		}
	})

}

// 方法一：单指针
func sortColors(nums []int) {
	/*
		在第一次遍历中，我们将数组中所有的 0 交换到数组的头部。在第二次遍历中，我们将数组中所有的 1 交换到头部的 0 之后。
	*/
	head := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[head] = nums[head], nums[i]
			head++
		}
	}
	for i := head; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[i], nums[head] = nums[head], nums[i]
			head++
		}
	}
	return
}

// 方法二：双指针

func sortColors2(nums []int) {
	/*
		用指针 p0来交换 0，初始值为 0; p2 来交换 2，初始值为 n−1
	*/
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
