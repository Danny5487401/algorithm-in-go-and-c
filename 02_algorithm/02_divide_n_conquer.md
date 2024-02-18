<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [分治法（Divide-and-Conquer Algorithm）](#%E5%88%86%E6%B2%BB%E6%B3%95divide-and-conquer-algorithm)
  - [分治与动态规划](#%E5%88%86%E6%B2%BB%E4%B8%8E%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92)
  - [分类](#%E5%88%86%E7%B1%BB)
  - [应用](#%E5%BA%94%E7%94%A8)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# 分治法（Divide-and-Conquer Algorithm）

分治分治，即分而治之。分治，就是把一个复杂的问题分成两个或更多的相同或相似的子问题，再把子问题分成更小的子问题……直到最后子问题可以简单的直接求解，原问题的解即子问题的解的合并。

![](02_algorithm/.02_divide_n_conquer_images/fenzhi.png)

## 分治与动态规划
当子问题相互独立时，能且只能使用分治。存在重叠子问题时，动态规划是更好的算法。 In a word, 分治法 —— 各子问题独立；动态规划 —— 各子问题重叠。


## 分类

- 递归分治


- 非递归分治
  !![](.02_divide_n_conquer_images/non_iterate_divide2.png)
  ![](.02_divide_n_conquer_images/non_iterate_divide.png)

## 应用

1. 最大值和最小值
   ![](.02_divide_n_conquer_images/largest_n_smallest.png)
   改进:区分问题
   ![](.02_divide_n_conquer_images/largest1.png)
   ![](.02_divide_n_conquer_images/largest2.png)

改进：分组
![](.02_divide_n_conquer_images/largest3.png)
![](.02_divide_n_conquer_images/largest4.png)


## 参考

