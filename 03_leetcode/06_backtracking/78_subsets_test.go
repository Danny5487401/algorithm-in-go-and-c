package _6_backtracking

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 子集 https://leetcode.cn/problems/subsets/?envType=study-plan-v2&envId=top-100-liked

func TestSubsets(t *testing.T) {
	convey.Convey("subsets ", t, func() {
		testCase := []struct {
			input  []int
			target [][]int
		}{
			{

				[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
			},
		}

		for _, tst := range testCase {
			rsp := subsets2(tst.input)
			compareRsp := intSliceSliceEqual(rsp, tst.target)
			convey.So(compareRsp, convey.ShouldBeTrue)
		}
	})

}

// 方法一：输入的视角（选或不选）
func subsets(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return [][]int{}
	}
	var ans = make([][]int, 0)
	var path = make([]int, 0, length)
	var dfs func(i int)
	dfs = func(i int) {
		if i == length {
			ans = append(ans, append([]int{}, path...)) // 固定答案
			return
		}
		// 下面两则顺序可更改（注意还原）
		// 1 不选
		dfs(i + 1)
		// 2 选
		path = append(path, nums[i])
		dfs(i + 1)

		// 选对应需要返回恢复现场
		path = path[:len(path)-1]
	}
	dfs(0)
	return ans
}

// 方案二： 从答案的视角，必须选一个
func subsets2(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return [][]int{}
	}

	var ans = make([][]int, 0)
	var path = make([]int, 0, length)
	var dfs func(i int)
	dfs = func(i int) {
		// 由于子集长度没有约束，所以递归的每个节点都是答案
		ans = append(ans, append([]int(nil), path...))
		if i == length {
			return
		}
		for j := i; j < length; j++ {
			path = append(path, nums[j])
			dfs(j + 1)

			// 选对应需要返回恢复现场
			path = path[:len(path)-1]
		}

	}
	dfs(0)
	return ans
}

func intSliceSliceEqual(v1, v2 [][]int) bool {
	if len(v1) != len(v2) {
		return false
	}

loop:
	for _, vv1 := range v1 {
		for _, vv2 := range v2 {
			if len(vv1) != len(vv2) {
				continue
			}
			if sliceSameValue(vv1, vv2) {
				continue loop
			}
		}

		return false
	}

	return true
}

func sliceSameValue(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	return cmp.Equal(s1, s2, cmpopts.SortSlices(func(i, j int) bool { return i < j }))
}
