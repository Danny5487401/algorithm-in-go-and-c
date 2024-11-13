package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 填充每个节点的下一个右侧节点指针 https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/description/
func TestConnect(t *testing.T) {
	convey.Convey("填充每个节点的下一个右侧节点指针: 给定一个*完美二叉树*,让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL", t, func() {
		testCase := []struct {
			input  *Node
			sum    int
			target *Node
		}{
			{
				CreateNodeByArray([]int{1, 2, 3, 4, 5, 6, 7}),
				22,
				CreateNode2(),
			},
			{
				CreateNodeByArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}),
				5,
				CreateNode3(),
			},
		}

		for _, tst := range testCase {
			rsp := connectBFS(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 方式一: DFS
func connect(root *Node) *Node {
	/*
			DFS 这棵树，从根节点 1 出发，向左递归到 2，再向左递归到 4。

			这三个节点正好是每一层的第一个节点（类似链表头），用一个数组 pre 记录，即 pre[0] 为节点 1，pre[1] 为节点 2，pre[2] 为节点 4。pre 的下标就是节点的深度

			继续递归到 5（深度为 2），从 pre[2] 中拿到节点 4，把 4 的 next 指向 5。然后更新 pre[2] 为节点 5，这样在后面递归到节点 6 时，就可以从 pre[2] 中拿到节点 5，把 5 的 next 指向 6 了。


		1. 创建一个空数组 pre（因为一开始不知道二叉树有多深）。
		2. DFS 这棵二叉树，递归参数为当前节点 node，以及当前节点的深度 depth。每往下递归一层，就把 depth 加一。
		3. 如果 depth 等于 pre 数组的长度，说明 node 是这一层最左边的节点，把 node 添加到 pre 的末尾。
		4. 否则，把 pre[depth] 的 next 指向 node，然后更新 pre[depth] 为 node。
		5. 递归边界：如果 node 是空节点，直接返回。
		6. 递归入口：dfs(root,0)。
		7. 最后返回 root。


	*/
	var pre []*Node
	var dfs func(*Node, int)
	dfs = func(n *Node, depth int) {
		if n == nil {
			return
		}
		if depth == len(pre) { // node 是这一层最左边的节点
			pre = append(pre, n)
		} else { // pre[depth] 是 node 左边的节点
			pre[depth].Next = n
			pre[depth] = n
		}
		depth++
		dfs(n.Left, depth)
		dfs(n.Right, depth)
	}
	dfs(root, 0)
	return root

}

// 方式二: BFS
func connectBFS(root *Node) *Node {
	// 只需要把每一层的节点，从左到右依次用 next 指针连接起来。
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	for len(queue) > 0 {

		tmp := queue
		//queue = nil // 每层归0
		for i, _ := range tmp {
			node := queue[0]
			queue = queue[1:]
			if i > 0 { // 从第二个开始
				tmp[i-1].Next = node
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

	}
	return root

}

func CreateNode2() *Node {
	node7 := &Node{
		Val: 7,
	}
	node6 := &Node{
		Val:  6,
		Next: node7,
	}
	node5 := &Node{
		Val:  5,
		Next: node6,
	}
	node4 := &Node{
		Val:  4,
		Next: node5,
	}
	node3 := &Node{
		Val:   3,
		Left:  node6,
		Right: node7,
	}
	node2 := &Node{
		Val:   2,
		Left:  node4,
		Right: node5,
		Next:  node3,
	}
	node1 := &Node{
		Val:   1,
		Left:  node2,
		Right: node3,
	}

	return node1
}
func CreateNode3() *Node {
	node15 := &Node{
		Val: 15,
	}
	node14 := &Node{
		Val:  14,
		Next: node15,
	}
	node13 := &Node{
		Val:  13,
		Next: node14,
	}
	node12 := &Node{
		Val:  12,
		Next: node13,
	}
	node11 := &Node{
		Val:  11,
		Next: node12,
	}
	node10 := &Node{
		Val:  10,
		Next: node11,
	}
	node9 := &Node{
		Val:  9,
		Next: node10,
	}
	node8 := &Node{
		Val:  8,
		Next: node9,
	}
	node7 := &Node{
		Val:   7,
		Left:  node14,
		Right: node15,
	}
	node6 := &Node{
		Val:   6,
		Left:  node12,
		Right: node13,
		Next:  node7,
	}
	node5 := &Node{
		Val:   5,
		Left:  node10,
		Right: node11,
		Next:  node6,
	}
	node4 := &Node{
		Val:   4,
		Left:  node8,
		Right: node9,
		Next:  node5,
	}
	node3 := &Node{
		Val:   3,
		Left:  node6,
		Right: node7,
	}
	node2 := &Node{
		Val:   2,
		Left:  node4,
		Right: node5,
		Next:  node3,
	}
	node1 := &Node{
		Val:   1,
		Left:  node2,
		Right: node3,
	}

	return node1
}
