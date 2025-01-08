package _0_heap

import (
	"container/heap"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 前 K 个高频元素 https://leetcode.cn/problems/top-k-frequent-elements/description/?envType=study-plan-v2&envId=top-100-liked
func TestTopKFrequent(t *testing.T) {
	convey.Convey("前 K 个高频元素 ", t, func() {
		testCase := []struct {
			input  []int
			k      int
			target []int
		}{
			{

				[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2},
			},

			{

				[]int{1}, 1, []int{1},
			},
		}

		for _, tst := range testCase {
			rsp := topKFrequent(tst.input, tst.k)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func topKFrequent(nums []int, k int) []int {

	// 首先遍历整个数组，并使用哈希表记录每个数字出现的次数，并形成一个「出现次数数组」

	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	/*
		- 如果堆的元素个数小于 k，就可以直接插入堆中。
		- 如果堆的元素个数等于 k，则检查堆顶与当前出现次数的大小。如果堆顶更大，说明至少有 k 个数字的出现次数比当前值大，故舍弃当前值；否则，就弹出堆顶，并将当前值插入堆中

	*/
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k { //保留 k 个
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int // 二元数组为 key, key出现的次数

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
