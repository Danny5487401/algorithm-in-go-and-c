package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 填充每个节点的下一个右侧节点指针 II https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/description/
func TestConnectII(t *testing.T) {
	convey.Convey("填充每个节点的下一个右侧节点指针: 给定一个*二叉树*,让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL", t, func() {
		testCase := []struct {
			input  *Node
			sum    int
			target *Node
		}{
			{
				CreateNodeByArray([]int{1, 2, 3, 4, 5, NULL, 7}),
				22,
				CreateNode2(),
			},
		}

		for _, tst := range testCase {
			rsp := connectBFS(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func connectII(root *Node) *Node {
	// 同 116
	return nil
}
