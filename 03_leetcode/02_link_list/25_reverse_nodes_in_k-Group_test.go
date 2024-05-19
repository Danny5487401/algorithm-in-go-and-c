package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// K 个一组翻转链表 https://leetcode.cn/problems/reverse-nodes-in-k-group/?envType=study-plan-v2&envId=top-100-liked

func TestReverseKGroup(t *testing.T) {
	convey.Convey(" K 个一组翻转链表", t, func() {
		testCase := []struct {
			input  *ListNode
			k      int
			target *ListNode
		}{
			{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
				2,
				&ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					}},
			},

			{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
				3,
				&ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					}},
			},
		}

		for _, tst := range testCase {
			rsp := reverseKGroup(tst.input, tst.k)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var n = 0
	cur := head
	for cur != nil {
		n++ // 获取链表长度
		cur = cur.Next
	}

	var dummy = &ListNode{
		Next: head,
	}
	p0 := dummy // 反转段的上一个节点
	for n >= k {
		n -= k
		// 同反转中间部分链表 https://leetcode.cn/problems/reverse-linked-list-ii/description/

		var pre *ListNode = nil
		var cur = p0.Next // 开始反转的起点
		for i := 0; i < k; i++ {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		// 记录下次循环的开始
		nxt := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = nxt
	}

	return dummy.Next
}
