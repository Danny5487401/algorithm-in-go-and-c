package _2_two_pinter

// 环形链表II https://leetcode.cn/problems/linked-list-cycle-ii/?envType=study-plan-v2&envId=top-100-liked

func detectCycle(head *ListNode) *ListNode {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			for slow != head {
				// 公式 a-c =(k-1)(b+c)
				// slow 从相遇点出发，head 从头节点出发，两个都走c步，slow 到入口，head 到入口的距离正好是环长的倍数，所以两个都继续走，肯定在入口相遇
				head = head.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
