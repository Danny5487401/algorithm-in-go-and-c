package _0_heap

import (
	"container/heap"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 数组中的第K个最大元素 https://leetcode.cn/problems/kth-largest-element-in-an-array/description/?envType=study-plan-v2&envId=top-100-liked
func TestFindKthLargest(t *testing.T) {
	convey.Convey("数组排序后的第 k 个最大的元素", t, func() {
		testCase := []struct {
			input  []int
			k      int
			target int
		}{
			{

				[]int{3, 2, 1, 5, 6, 4}, 2, 5,
			},

			{

				[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4,
			},
		}

		for _, tst := range testCase {
			rsp := findKthLargest(tst.input, tst.k)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func findKthLargest(nums []int, k int) int {
	example := IntHeap(nums)
	heap.Init(&example)
	var ans int
	for i := 0; i < k; i++ {
		ans = heap.Pop(&example).(int)

	}
	return ans
}

// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
