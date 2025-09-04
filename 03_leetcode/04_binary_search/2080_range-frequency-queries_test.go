package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

// 区间内查询数字的频率 https://leetcode.cn/problems/range-frequency-queries/description/

func TestConstructorRangeFreqQuery(t *testing.T) {
	convey.Convey("子数组中一个值的 频率 指的是这个子数组中这个值的出现次数。", t, func() {
		/*
			- 中序遍历: [ [左子树的前序遍历结果], 根节点,[右子树的前序遍历结果] ]
		*/
		testCase := []struct {
			RangeFreqQuery []int
			query1         []int
			query2         []int
			target         []int
		}{
			{
				[]int{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56},
				[]int{1, 2, 4},   // 返回 1 。4 在子数组 [33, 4] 中出现 1 次
				[]int{0, 11, 33}, // 返回 2 。33 在整个子数组中出现 2 次。
				[]int{1, 2},
			},
		}

		for _, tst := range testCase {
			c := Constructor(tst.RangeFreqQuery)
			ans1 := c.Query(tst.query1[0], tst.query1[1], tst.query1[2])
			ans2 := c.Query(tst.query2[0], tst.query2[1], tst.query2[2])
			ans := []int{ans1, ans2}
			convey.So(ans, convey.ShouldEqual, tst.target)
		}
	})

}

/*
满足 left≤i≤right 的下标 i 的个数。
例如 query(3,5,2)，由于数字 2 的下标列表 [1,4,5] 中的下标 4 和 5 都在区间 [3,5] 中，所以返回 2。

*/

type RangeFreqQuery map[int][]int

func Constructor(arr []int) RangeFreqQuery {
	pos := map[int][]int{}
	for i, x := range arr {
		pos[x] = append(pos[x], i)
	}
	return pos

}

func (r *RangeFreqQuery) Query(left int, right int, value int) int {

	a := (*r)[value]

	// > right 等价于 >= right+1
	// a 中的下标在 [p,q−1] 内的数都是满足要求的，
	return sort.SearchInts(a, right+1) - sort.SearchInts(a, left)
}
