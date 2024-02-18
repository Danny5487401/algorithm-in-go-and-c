<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [algorithm-in-go-and-c](#algorithm-in-go-and-c)
  - [第一章 数据结构与算法](#%E7%AC%AC%E4%B8%80%E7%AB%A0-%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E7%AE%97%E6%B3%95)
  - [第二章  ](#%E7%AC%AC%E4%BA%8C%E7%AB%A0)
  - [第三章 LeetCode](#%E7%AC%AC%E4%B8%89%E7%AB%A0-leetcode)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# algorithm-in-go-and-c
c 和 go 实现 数据结构与算法


## 第一章 数据结构与算法
Note: 使用 GBK 编码
- [1 linear list线性表](01_dataStructure/01_linear_list/linear_list.md)
    - [1.1 顺序表 静态实现，数据元素是整数](01_dataStructure/01_linear_list/seqlist1.c)
    - [1.2 顺序表 静态实现，数据元素是结构体](01_dataStructure/01_linear_list/seqlist2.c)
    - [1.3 顺序表 动态实现，数据元素是整数](01_dataStructure/01_linear_list/seqlist3.c)
    - [1.4 顺序表 动态实现，数据元素是结构体](01_dataStructure/01_linear_list/seqlist4.c)
- [2 link list链表](01_dataStructure/02_link_list/link_list.md)
    - [2.1 带头结点的单链表的实现，数据元素是整数](01_dataStructure/02_link_list/linklist1.c)
    - [2.2 带头结点的单链表的实现，数据元素是结构体](01_dataStructure/02_link_list/linklist2.c)
    - [2.3 不带头结点的单链表的实现](01_dataStructure/02_link_list/linklist3.c)
    - [2.4 带头结点的双链表的实现，数据元素是整数](01_dataStructure/02_link_list/linklist4.c)
    - [2.5 带头结点的循环单链表的实现，数据元素是整数](01_dataStructure/02_link_list/linklist5.c)
    - 2.6 常见的考题


- [3 B 树](01_dataStructure/03_btree/btree.md)
    - [3.1 二叉树的层次遍历](01_dataStructure/03_btree/btree1.c)
    - [3.2 二叉树的前序遍历、中序遍历和后序遍历，包括递归和非递归两种方法](01_dataStructure/03_btree/btree2.c)
    - [3.3 中序线索二叉树的创建及求前驱后继的方法](01_dataStructure/03_btree/btree3.c)
    - [3.4 二叉排序树的各种操作，包括插入、删除、查找](01_dataStructure/03_btree/btree4.c)


- [4 图](01_dataStructure/04_graph/graph.md)
    - 4.1 BFS (Breadth First Search 广度优先遍历）
    - 4.2 DFS (Depth First Search 深度优先搜索)
    - 4.3 单源最短路径算法(Floyd(弗洛伊德)算法)
    - 4.4 拓扑排序
    - 4.5 关键路径
- 5 hash哈希表

- [6 queue队列](01_dataStructure/06_queue/queue.md)
    - [6.1 循环队列的数组实现，队尾指针指向队尾的下一个元素，没有length的辅助变量](01_dataStructure/06_queue/seqqueue1.c)
    - [6.2 循环队列的数组实现，队尾指针指向队尾的下一个元素，增加了length的辅助变量](01_dataStructure/06_queue/seqqueue2.c)
    - [6.3 循环队列的数组实现，队尾指针指向队尾元素，没有length的辅助变量](01_dataStructure/06_queue/seqqueue3.c)
    - [6.4 循环队列的数组实现，队尾指针指向队尾元素，增加了length的辅助变量](01_dataStructure/06_queue/seqqueue4.c)
    - [6.5 队列的链表实现（带头结点）](01_dataStructure/06_queue/linkqueue1.c)

- [7 stack栈](01_dataStructure/07_stack/stack.md)
    - [7.1 顺序栈的实现，数据元素是整数](01_dataStructure/07_stack/seqstack1.c)
    - [7.2 链栈的实现，数据元素是整数](01_dataStructure/07_stack/linkstack1.c)
    - [7.3 顺序栈检查括号是否匹配，支持()[]{}三种括号](01_dataStructure/07_stack/seqstack2.c)
    - [7.4 用顺序栈实现中缀表达式转后缀表达式](01_dataStructure/07_stack/seqstack3.c)

- [8 排序算法](01_dataStructure/08_sort/sort.md)
    - 8.1 冒泡排序
    - 8.2 桶排序
    - 8.3 计数排序
    - [8.4 堆排序](01_dataStructure/08_sort/heapsort.c)
    - [8.5 插入排序](01_dataStructure/08_sort/insertsort.c)
    - 8.6 归并排序
        - [递归的方法实现](01_dataStructure/08_sort/mergesort.c)
        - [循环的方法实现](01_dataStructure/08_sort/mergesort1.c)
    - 8.7 快速排序
    - 8.8 基数排序
    - [8.9 选择排序](01_dataStructure/08_sort/selectsort1.c)
    - [8.10 希尔排序](01_dataStructure/08_sort/shellsort.c)

- [9 数组和广义表](01_dataStructure/09_array/array.md)
- [10 串](01_dataStructure/10_string/string.md)
- [11 查找](01_dataStructure/11_search/search.md)
    - [11.1 顺序查找](01_dataStructure/11_search/seqsearch.c)
    - [11.1 二分查找](01_dataStructure/11_search/binsearch.c)
- [12 递归](01_dataStructure/12_recursive/recursive.md)

## [第二章  ](02_algorithm/algorithm.md)

- [1 蛮力法 Exhausitive Attack](02_algorithm/01_Exhaustive_Attack.md)
- [2 分治法 Divide and Conquer](02_algorithm/02_divide_n_conquer.md)
- [3 贪心法 Greedy Method](02_algorithm/03_greedy_method.md)
- [4 动态规划 dynamic programming](02_algorithm/04_dynamic_programming.md)


## 第三章 LeetCode
- hash
  - [1 两数之和](03_leetcode/01_hash/01_two_sum_test.go)
  - [49 字母异位词分组](03_leetcode/01_hash/49_group_anagrams_test.go)
  - [128 最长连续序列](03_leetcode/01_hash/128_longest_consecutive_sequence_test.go)

## 参考
- [The Algorithms 组织开源的多种语言算法实现](https://github.com/TheAlgorithms)