package _2_link_list

import (
	"container/list"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// LRU 缓存 https://leetcode.cn/problems/lru-cache/description/?envType=study-plan-v2&envId=top-100-liked
func TestLRUCache(t *testing.T) {
	convey.Convey("LRUCache", t, func() {
		lRUCache := Constructor(2)
		lRUCache.Put(1, 1)                                 // 缓存是 {1=1}
		lRUCache.Put(2, 2)                                 // 缓存是 {1=1, 2=2}
		convey.So(lRUCache.Get(1), convey.ShouldEqual, 1)  //  返回 1
		lRUCache.Put(3, 3)                                 // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
		convey.So(lRUCache.Get(2), convey.ShouldEqual, -1) // 返回 -1 (未找到)
		lRUCache.Put(4, 4)                                 // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
		convey.So(lRUCache.Get(1), convey.ShouldEqual, -1) // 返回 -1 (未找到)
		convey.So(lRUCache.Get(3), convey.ShouldEqual, 3)  // 返回 3
		convey.So(lRUCache.Get(4), convey.ShouldEqual, 4)  // 返回 4
	})

}

type LRUCache struct {
	capacity    int
	array       *list.List
	entryToNode map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:    capacity,
		entryToNode: make(map[int]*list.Element),
		array:       list.New(),
	}
}

type entry struct {
	key, value int
}

func (l *LRUCache) Get(key int) int {
	var defaultVal = -1
	if node, ok := l.entryToNode[key]; ok {
		defaultVal = node.Value.(entry).value
		l.array.MoveToFront(node)
		return defaultVal
	}
	return defaultVal
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.entryToNode[key]; ok {
		node.Value = entry{key, value} // 更新
		l.array.MoveToFront(node)
		return
	}
	l.entryToNode[key] = l.array.PushFront(entry{key, value})
	if len(l.entryToNode) > l.capacity {
		delete(l.entryToNode, l.array.Remove(l.array.Back()).(entry).key)
	}

}
