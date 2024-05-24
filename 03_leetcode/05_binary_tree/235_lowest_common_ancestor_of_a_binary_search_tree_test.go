package _5_binary_tree

import (
	"reflect"
	"testing"
)

// 二叉搜索树的最近公共祖先 https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/?envType=study-plan-v2&envId=top-100-liked
func TestLowestCommonAncestorBinarySearchTree(t *testing.T) {
	root1 := CreateTreeByArray([]int{6, 2, 8, 0, 4, 7, 9, 0, 0, 3, 5})
	root2 := CreateTreeByArray([]int{6, 2, 8, 0, 4, 7, 9, 0, 0, 3, 5})
	root3 := CreateTreeByArray([]int{1, 2})

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
				p:    root1.Left,  // 2
				q:    root1.Right, // 8
			},
			want: root1, // 3
		},
		{
			name: "ex2",
			args: args{
				root: root2,
				p:    root2.Left,       // 2
				q:    root2.Left.Right, // 4
			},
			want: root2.Left, // 5
		},
		{
			name: "ex3",
			args: args{
				root: root3,
				p:    root3,      // 2
				q:    root3.Left, // 1
			},
			want: root3, // 1
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor2(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}

			//if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			//}
		})
	}

}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	val := root.Val
	if p.Val < val && q.Val < val {
		// 在左子树中
		return lowestCommonAncestor(root.Left, p, q)
	}

	if p.Val > val && q.Val > val {
		// 在右子树中
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}
