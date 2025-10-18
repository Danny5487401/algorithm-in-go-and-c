<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [数组](#%E6%95%B0%E7%BB%84)
  - [差分数组 difference array](#%E5%B7%AE%E5%88%86%E6%95%B0%E7%BB%84-difference-array)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


# 数组


## 差分数组 difference array


a=[1,3,3,5,8] ,其中的相邻元素两两作差（右边减左边），得到数组 [2,0,2,3], 然后在开头补上 a[0]，得到差分数组 d=[1,2,0,2,3]


如果从左到右累加 d 中的元素，我们就「还原」回了 a 数组 [1,3,3,5,8]


现在把连续子数组 a[1],a[2],a[3] 都加上 10，得到 a′ =[1,13,13,15,8]。得到差分数组, d ′ =[1,12,0,2,−7]

对比 d 和 d ′ ，可以发现只有 d[1] 和 d[4] 变化了，这意味着对 a 中连续子数组的操作，可以转变成对差分数组 d 中两个数的操作。

对于数组 a，定义其差分数组（difference array）为

d[i]={
a[0],   i=0
a[i]−a[i−1],  i≥1




性质 1：从左到右累加 d 中的元素，可以得到数组 a。

性质 2：如下两个操作是等价的。

把 a 的子数组 a[i],a[i+1],…,a[j] 都加上 x。
把 d[i] 增加 x，把 d[j+1] 减少 x。 