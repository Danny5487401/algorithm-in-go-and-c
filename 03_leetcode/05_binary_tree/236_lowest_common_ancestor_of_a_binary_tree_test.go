package _5_binary_tree

import (
	"reflect"
	"testing"
)

// 二叉树最近公共祖先 https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/?envType=study-plan-v2&envId=top-100-liked
func TestLowestCommonAncestor(t *testing.T) {
	root1 := Ints2TreeNode([]int{3, 5, 1, 6, 2, 9, 8, NULL, NULL, 7, 4})
	root2 := Ints2TreeNode([]int{3, 5, 1, 6, 2, 9, 8, NULL, NULL, 7, 4})
	root3 := Ints2TreeNode([]int{1, 2})

	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "ex1",
			args: args{
				root: root1,
				p:    root1.Left,  // 5
				q:    root1.Right, // 1
			},
			want: root1, // 3
		},
		{
			name: "ex2",
			args: args{
				root: root2,
				p:    root2.Left,             // 5
				q:    root2.Left.Right.Right, // 4
			},
			want: root2.Left, // 5
		},
		{
			name: "ex3",
			args: args{
				root: root3,
				p:    root3,      // 1
				q:    root3.Left, // 2
			},
			want: root3, // 1
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}

			//if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			//}
		})
	}

}

// 满足 x 是 p、q 的祖先且 x 的深度尽可能大
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		// 如果当前节点为空，为p,为q, 直接返回
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		// 如果左右都有
		return root
	}

	if left != nil {
		// 只有左边找到
		return left
	}
	return right
}
