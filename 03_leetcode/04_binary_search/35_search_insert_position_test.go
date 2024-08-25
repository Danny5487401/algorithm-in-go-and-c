package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

// 搜索插入位置 https://leetcode.cn/problems/search-insert-position/?envType=study-plan-v2&envId=top-100-liked

func TestSearchInsert(t *testing.T) {
	convey.Convey("搜索插入位置:排序数组，相同元素取该索引", t, func() {
		testCase := []struct {
			input    []int
			target   int
			expected int
		}{
			{

				[]int{1, 3, 5, 6}, 5, 2,
			},
			{
				[]int{1, 3, 5, 6}, 2, 1,
			},
			{
				[]int{1, 3, 5, 6}, 7, 4,
			},
		}

		for _, tst := range testCase {
			rsp := searchInsert2(tst.input, tst.target)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

func searchInsert(nums []int, target int) int {
	index := sort.SearchInts(nums, target)
	return index
}

func searchInsert2(nums []int, target int) int {
	index := lowerBound3(nums, target)
	return index
}
