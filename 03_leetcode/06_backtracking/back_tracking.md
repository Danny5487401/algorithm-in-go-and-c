<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [回溯法（back tracking）（探索与回溯法](#%E5%9B%9E%E6%BA%AF%E6%B3%95back-tracking%E6%8E%A2%E7%B4%A2%E4%B8%8E%E5%9B%9E%E6%BA%AF%E6%B3%95)
  - [分类](#%E5%88%86%E7%B1%BB)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# 回溯法（back tracking）（探索与回溯法


是一种选优搜索法，又称为试探法，按选优条件向前搜索，以达到目标。但当探索到某一步时，发现原先选择并不优或达不到目标，就退回一步重新选择，
这种走不通就退回再走的技术为回溯法，而满足回溯条件的某个状态的点称为“回溯点.

![](.back_tracking_images/back_tracking.png) 
![](.back_tracking_images/back_tracking2.png)



![](.back_tracking_images/back_tracking_info.png)

dfs(i)-->dfs(i+1)
注意：参数 i 的含义不是第 i 个，而是 >= i 的部分


## 分类
[电话号码的字母组合](17_letter_combinations_of_a_phone_number_test.go)

1. [子集型回溯](78_subsets_test.go)

站在输入的视角: 每个元素选或不选
![](.back_tracking_images/subsequence_back_tracking.png)
站在答案的视角：每个节点都是答案
![](.back_tracking_images/sub_sequence_back_tracking2.png)
![](.back_tracking_images/sub_sequence_back_tracking3.png)

- 相同案例：[分割回文串](131_palindrome_partition_test.go)

![](.back_tracking_images/palindrome_partition1.png)
![](.back_tracking_images/sub_sequnce_back_tracking4.png)

2. [组合型回溯：可以减枝优化](./77_combinations_test.go)

假设遍历路径 path 长度为 m, 那么还需要查找 d=k-m 个数。如果是从[1,i]倒序这i个数中选数，如果i<d,那么肯定选不出k个数，直接return,这就是剪枝。

![](.back_tracking_images/back_tracking_combinaton2.png)


[组合总和 III](./216_combination_sum3_test.go)
![](.back_tracking_images/combinationSum3.png)



[括号生成](./22_generate_parentheses_test.go)

![](.back_tracking_images/GenerateParentheses.png)
![](.back_tracking_images/Generate Parentheses2.png)
![](.back_tracking_images/Generate Parentheses3.png)

dfs[i,open]
- dfs[i+1,open+1] 选左括号
- dfs[i+1,open] 选右括号


3. 排列型回溯

![](.back_tracking_images/Permutations.png)

![]( .back_tracking_images/permutations2.png)