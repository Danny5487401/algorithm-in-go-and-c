package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"slices"
	"sort"
	"testing"
)

// 咒语和药水的成功对数 https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/description/
func TestSuccessfulPairs(t *testing.T) {
	convey.Convey("二叉树的中序遍历", t, func() {
		/*
			- 中序遍历: [ [左子树的前序遍历结果], 根节点,[右子树的前序遍历结果] ]
		*/
		testCase := []struct {
			spells  []int
			potions []int
			success int64
			target  []int
		}{
			{
				/*
					- 第 0 个咒语：5 * [1,2,3,4,5] = [5,10,15,20,25] 。总共 4 个成功组合。
					- 第 1 个咒语：1 * [1,2,3,4,5] = [1,2,3,4,5] 。总共 0 个成功组合。
					- 第 2 个咒语：3 * [1,2,3,4,5] = [3,6,9,12,15] 。总共 3 个成功组合。
				*/
				[]int{5, 1, 3},
				[]int{1, 2, 3, 4, 5},
				7,
				[]int{4, 0, 3},
			},
			{
				/*
					- 第 0 个咒语：3 * [8,5,8] = [24,15,24] 。总共 2 个成功组合。
					- 第 1 个咒语：1 * [8,5,8] = [8,5,8] 。总共 0 个成功组合。
					- 第 2 个咒语：2 * [8,5,8] = [16,10,16] 。总共 2 个成功组合。
				*/
				[]int{3, 1, 2},
				[]int{8, 5, 8},
				16,
				[]int{2, 0, 2},
			},
		}

		for _, tst := range testCase {
			rsp := successfulPairs(tst.spells, tst.potions, tst.success)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 方法一：排序 + 二分查找
/*
对于正整数，xy≥success 等价于 y≥⌈success⌉/x

上取整下取整转换公式的证明:
	- https://zhuanlan.zhihu.com/p/1890356682149838951



y≥⌈success⌉/x = ⌊success−1/x⌋+1
*/
func successfulPairs(spells []int, potions []int, success int64) []int {
	// 求 >=
	slices.Sort(potions)
	for i, x := range spells {
		spells[i] = len(potions) - sort.SearchInts(potions, (int(success)-1)/x+1)
	}

	return spells

}
