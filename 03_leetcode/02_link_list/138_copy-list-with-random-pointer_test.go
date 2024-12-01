package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 随机链表的复制 https://leetcode.cn/problems/copy-list-with-random-pointer/description/?envType=study-plan-v2&envId=top-100-liked
func TestCopyRandomList(t *testing.T) {
	convey.Convey("链表的 深拷贝", t, func() {

	})
}

func copyRandomList(head *Node) *Node {
	/*
		    依次复制每个节点（创建新节点并复制 val 和 next），把新节点直接插到原节点的后面，形成一个交错链表
			1→1′→2→2′→3→3′
			最后，从交错链表链表中分离出 1′→2′→3′

	*/
	if head == nil {
		return nil
	}

	// 复制每个节点，把新节点直接插到原节点的后面
	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
	}

	// 遍历交错链表中的原链表节点 复制 random
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next // 要复制的 random 是 cur.Random 的下一个节点
		}
	}

	// 奇偶索引分割链表
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil { // 从偶数开始就可以
		odd.Next = even.Next // 更新奇数节点时，奇数节点的后一个节点需要指向偶数节点的后一个节点
		odd = odd.Next
		even.Next = odd.Next // 更新偶数节点时，偶数节点的后一个节点需要指向奇数节点的后一个节点
		even = even.Next
	}
	// 取偶数索引即可
	odd.Next = nil
	return evenHead
}
