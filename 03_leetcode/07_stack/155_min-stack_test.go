package _7_stack

import "math"

// 最小栈 https://leetcode.cn/problems/min-stack/?envType=study-plan-v2&envId=top-100-liked
type MinStack struct {
	// preMin[0]=nums[0]
	// preMin[1]=min(nums[0],nums[1])
	// preMin[2]=min(nums[0],nums[1],nums[2])=min(preMin[1],nums[2])
	// preMin[i]=min(preMin[i−1],nums[i])
	vals []pair //维护前缀最小值
}

type pair struct{ val, preMin int } // 栈中除了保存添加的元素，还保存前缀最小值

func Constructor() MinStack {
	return MinStack{
		vals: []pair{
			{0, math.MaxInt}, // 初始化的时候，在栈底加一个 ∞ 哨兵，作为 preMin[−1]
		},
	}

}

func (this *MinStack) Push(val int) {
	// 设当前栈的大小是 n。添加元素 val 后，额外维护 preMin[n]=min(preMin[n−1],val)，其中 preMin[n−1] 是添加 val 之前，栈顶保存的前缀最小值

	this.vals = append(this.vals, pair{val, min(this.GetMin(), val)})

}

func (this *MinStack) Pop() {
	this.vals = this.vals[:len(this.vals)-1]
}

func (this *MinStack) Top() int {
	return this.vals[len(this.vals)-1].val
}

func (this *MinStack) GetMin() int {
	return this.vals[len(this.vals)-1].preMin
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
