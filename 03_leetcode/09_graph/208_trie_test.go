package _9_graph

// 实现 Trie (前缀树) https://leetcode.cn/problems/implement-trie-prefix-tree/?envType=study-plan-v2&envId=top-100-liked
// 可参考 github.com/dghubble/trie

type Node struct {
	son [26]*Node // 26 叉树的每个节点
	end bool      // 是否为终止节点
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{&Node{}}
}

func (t Trie) Insert(word string) {
	cur := t.root
	for _, c := range word {
		c -= 'a'
		if cur.son[c] == nil { // 如果 word[i] 不是 cur 的儿子，那么创建一个新的节点 node 作为 cur 的儿子
			cur.son[c] = &Node{}
		}
		cur = cur.son[c]
	}
	cur.end = true
}

func (t Trie) find(word string) int {
	cur := t.root
	for _, c := range word {
		c -= 'a'
		if cur.son[c] == nil {
			return 0
		}
		cur = cur.son[c]
	}
	if cur.end {
		return 2
	}
	return 1
}

func (t Trie) Search(word string) bool {
	return t.find(word) == 2
}

func (t Trie) StartsWith(prefix string) bool {
	return t.find(prefix) != 0
}
