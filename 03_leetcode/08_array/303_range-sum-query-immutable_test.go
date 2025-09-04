package _8_array

// 区域和检索 - 数组不可变 https://leetcode.cn/problems/range-sum-query-immutable/description/
type NumArray []int

func Constructor(nums []int) NumArray {
	s := make(NumArray, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	return s
}

// 计算索引 left 和 right （包含 left 和 right）之间的 nums 元素的 和
func (s NumArray) SumRange(left, right int) int {
	return s[right+1] - s[left]
}

// 前缀和解析: https://leetcode.cn/problems/range-sum-query-immutable/solutions/2693498/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/

/*
任意子数组的和，都可以表示为两个前缀和的差。


把 nums 记作 a，设其长度为 n

s[0]=0
s[1]=a[0]

s[n]=a[0]+...+a[n-1]




前 i 个数的和，加上 a[i]，就是前 i+1 个数的和

s[i+1]=s[i]+a[i]


示例中的数组 [−2,0,3,−5,2,−1]，对应的前缀和数组 s=[0,−2,−2,1,−4,−2,−3]。


子数组的元素和转化成两个前缀和的差，下标区间 [left,right] 的元素和等于前缀 [0,right] 的元素和减去另一个前缀 [0,left−1] 的元素和

s[right+1]−s[left]

子数组 [3,−5,2,−1] 的和，就可以 O(1) 地用 s[6]−s[2]=−3−(−2)=−1 算出来


为什么要定义 s[0]=0?
如果不定义 s[0]=0，就必须特判 left=0 的情况.
*/
