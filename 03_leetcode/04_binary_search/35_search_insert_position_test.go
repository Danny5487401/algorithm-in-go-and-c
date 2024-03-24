package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

// 搜索插入位置 https://leetcode.cn/problems/search-insert-position/?envType=study-plan-v2&envId=top-100-liked

func TestSearchInsert(t *testing.T) {
	convey.Convey("Find First and Last Position of Element in Sorted Array ", t, func() {
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
			rsp := searchInsert(tst.input, tst.target)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

func searchInsert(nums []int, target int) int {
	index := sort.SearchInts(nums, target)
	return index
}
