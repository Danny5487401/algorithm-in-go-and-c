package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 两数相加 https://leetcode.cn/problems/add-two-numbers/description/
func TestAddTwoNumbers(t *testing.T) {
	convey.Convey("两数相加:每位数字都是按照 逆序 的方式存储的", t, func() {
		testCase := []struct {
			input1 *ListNode
			input2 *ListNode
			target *ListNode
		}{
			{
				GetListNodeBySlice([]int{2, 4, 3}),
				GetListNodeBySlice([]int{5, 6, 4}),
				GetListNodeBySlice([]int{7, 0, 8}),
			},

			{
				GetListNodeBySlice([]int{0}),
				GetListNodeBySlice([]int{0}),
				GetListNodeBySlice([]int{0}),
			},
			{
				GetListNodeBySlice([]int{9, 9, 9, 9, 9, 9, 9}),
				GetListNodeBySlice([]int{9, 9, 9, 9}),
				GetListNodeBySlice([]int{8, 9, 9, 9, 0, 0, 0, 1}),
			},
		}

		for _, tst := range testCase {
			rsp := addTwoIterate(tst.input1, tst.input2)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwo(l1, l2, 0)
}

// 方法一：递归
func addTwo(l1, l2 *ListNode, carry int) *ListNode {
	// 简化代码:
	// 1. 这里返回 l1,
	// 2. 如果 l2 比 l1 长, 交换
	if l1 == nil && l2 == nil {
		if carry != 0 {
			return &ListNode{Val: carry} // 如果进位了，就额外创建一个节点
		}
		return nil
	}
	if l1 == nil { // 如果 l1 是空的，那么此时 l2 一定不是空节点
		l1, l2 = l2, l1
	}

	sum := carry + l1.Val // 节点值和进位加在一起

	if l2 != nil {
		sum += l2.Val // 节点值和进位加在一起
		l2 = l2.Next  // 下一个节点
	}

	l1.Val = sum % 10                     // 每个节点保存一个数位
	l1.Next = addTwo(l1.Next, l2, sum/10) // 进位
	return l1
}

// 方法二：迭代
func addTwoIterate(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0 // 进位

	for l1 != nil || l2 != nil || carry != 0 { // 有一个不是空节点，或者还有进位，就继续迭代
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: carry % 10} // 每个节点保存一个数位
		carry = carry / 10

		cur = cur.Next

	}
	return dummy.Next
}
