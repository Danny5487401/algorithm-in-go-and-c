package _2_bits

import (
	"github.com/smartystreets/goconvey/convey"
	"math/bits"
	"sort"
	"testing"
)

// 根据数字二进制下 1 的数目排序 https://leetcode.cn/problems/sort-integers-by-the-number-of-1-bits/description/
func TestSortByBits(t *testing.T) {
	convey.Convey("将数组中的元素按照其二进制表示中数字 1 的数目升序排序."+
		"如果存在多个数字二进制中 1 的数目相同，则必须将它们按照数值大小升序排列", t, func() {
		testCase := []struct {
			input  []int
			target []int
		}{
			{
				/*
					[0] 是唯一一个有 0 个 1 的数。
					[1,2,4,8] 都有 1 个 1 。
					[3,5,6] 有 2 个 1 。
					[7] 有 3 个 1
				*/
				[]int{0, 1, 2, 3, 4, 5, 6, 7, 8},
				[]int{0, 1, 2, 4, 8, 3, 5, 6, 7},
			},
			{
				// 数组中所有整数二进制下都只有 1 个 1 ，所以你需要按照数值大小将它们排序
				[]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
				[]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
			}, {
				[]int{10000},
				[]int{10000},
			},
			{
				[]int{2, 3, 5, 7, 11, 13, 17, 19},
				[]int{2, 3, 5, 17, 7, 11, 13, 19},
			},
		}

		for _, tst := range testCase {
			rsp := sortByBits(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func sortByBits(arr []int) []int {
	arr1 := OneBits(arr)
	sort.Sort(arr1)
	return arr1
}

// 暴力方式
func sortByBits1(arr []int) []int {
	arr1 := OneBits(arr)
	sort.Sort(arr1)
	return arr1
}

type OneBits []int

func (b OneBits) Len() int {
	return len(b)
}

func (b OneBits) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b OneBits) Less(i, j int) bool {
	if bits.OnesCount64(uint64(b[i])) == bits.OnesCount64(uint64(b[j])) {
		return b[i] < b[j] // 如果存在多个数字二进制中 1 的数目相同，则必须将它们按照数值大小升序排列"
	}
	return bits.OnesCount64(uint64(b[i])) < bits.OnesCount64(uint64(b[j]))
}
