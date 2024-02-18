<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [树](#%E6%A0%91)
  - [基本术语](#%E5%9F%BA%E6%9C%AC%E6%9C%AF%E8%AF%AD)
  - [树的性质](#%E6%A0%91%E7%9A%84%E6%80%A7%E8%B4%A8)
  - [binary_tree二叉树](#binary_tree%E4%BA%8C%E5%8F%89%E6%A0%91)
    - [性质](#%E6%80%A7%E8%B4%A8)
  - [二叉树和度为2的树的区别](#%E4%BA%8C%E5%8F%89%E6%A0%91%E5%92%8C%E5%BA%A6%E4%B8%BA2%E7%9A%84%E6%A0%91%E7%9A%84%E5%8C%BA%E5%88%AB)
    - [满二叉树](#%E6%BB%A1%E4%BA%8C%E5%8F%89%E6%A0%91)
    - [完全二叉树](#%E5%AE%8C%E5%85%A8%E4%BA%8C%E5%8F%89%E6%A0%91)
    - [二叉排序树](#%E4%BA%8C%E5%8F%89%E6%8E%92%E5%BA%8F%E6%A0%91)
  - [二叉树的存储结构](#%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E5%AD%98%E5%82%A8%E7%BB%93%E6%9E%84)
    - [完全二叉树的顺序存储](#%E5%AE%8C%E5%85%A8%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E9%A1%BA%E5%BA%8F%E5%AD%98%E5%82%A8)
    - [二叉树的链式存储](#%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E9%93%BE%E5%BC%8F%E5%AD%98%E5%82%A8)
  - [二叉树的遍历](#%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E9%81%8D%E5%8E%86)
    - [层次遍历](#%E5%B1%82%E6%AC%A1%E9%81%8D%E5%8E%86)
      - [过程](#%E8%BF%87%E7%A8%8B)
    - [三种遍历方式](#%E4%B8%89%E7%A7%8D%E9%81%8D%E5%8E%86%E6%96%B9%E5%BC%8F)
      - [遍历方式代码实现](#%E9%81%8D%E5%8E%86%E6%96%B9%E5%BC%8F%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [二叉树的遍历](#%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E9%81%8D%E5%8E%86-1)
  - [线索二叉树](#%E7%BA%BF%E7%B4%A2%E4%BA%8C%E5%8F%89%E6%A0%91)
    - [线索二叉树数据结构](#%E7%BA%BF%E7%B4%A2%E4%BA%8C%E5%8F%89%E6%A0%91%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84)
    - [线索二叉树求前继和后继](#%E7%BA%BF%E7%B4%A2%E4%BA%8C%E5%8F%89%E6%A0%91%E6%B1%82%E5%89%8D%E7%BB%A7%E5%92%8C%E5%90%8E%E7%BB%A7)
      - [先序线索二叉树求后继](#%E5%85%88%E5%BA%8F%E7%BA%BF%E7%B4%A2%E4%BA%8C%E5%8F%89%E6%A0%91%E6%B1%82%E5%90%8E%E7%BB%A7)
      - [中序线索二叉树求前继和后继](#%E4%B8%AD%E5%BA%8F%E7%BA%BF%E7%B4%A2%E4%BA%8C%E5%8F%89%E6%A0%91%E6%B1%82%E5%89%8D%E7%BB%A7%E5%92%8C%E5%90%8E%E7%BB%A7)
  - [树的存储结构](#%E6%A0%91%E7%9A%84%E5%AD%98%E5%82%A8%E7%BB%93%E6%9E%84)
    - [1. 孩子兄弟表示法----主要](#1-%E5%AD%A9%E5%AD%90%E5%85%84%E5%BC%9F%E8%A1%A8%E7%A4%BA%E6%B3%95----%E4%B8%BB%E8%A6%81)
    - [2. 双亲表示法](#2-%E5%8F%8C%E4%BA%B2%E8%A1%A8%E7%A4%BA%E6%B3%95)
    - [3. 孩子表示法](#3-%E5%AD%A9%E5%AD%90%E8%A1%A8%E7%A4%BA%E6%B3%95)
  - [二叉排序（搜索）树(binary sort tree)](#%E4%BA%8C%E5%8F%89%E6%8E%92%E5%BA%8F%E6%90%9C%E7%B4%A2%E6%A0%91binary-sort-tree)
    - [查找,插入,创建,删除](#%E6%9F%A5%E6%89%BE%E6%8F%92%E5%85%A5%E5%88%9B%E5%BB%BA%E5%88%A0%E9%99%A4)
    - [二叉排序树的查找效率分析](#%E4%BA%8C%E5%8F%89%E6%8E%92%E5%BA%8F%E6%A0%91%E7%9A%84%E6%9F%A5%E6%89%BE%E6%95%88%E7%8E%87%E5%88%86%E6%9E%90)
  - [平衡二叉树 ---代码实现过于复杂](#%E5%B9%B3%E8%A1%A1%E4%BA%8C%E5%8F%89%E6%A0%91----%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0%E8%BF%87%E4%BA%8E%E5%A4%8D%E6%9D%82)
    - [插入和删除 数据后不平衡](#%E6%8F%92%E5%85%A5%E5%92%8C%E5%88%A0%E9%99%A4-%E6%95%B0%E6%8D%AE%E5%90%8E%E4%B8%8D%E5%B9%B3%E8%A1%A1)
    - [树高和节点关系](#%E6%A0%91%E9%AB%98%E5%92%8C%E8%8A%82%E7%82%B9%E5%85%B3%E7%B3%BB)
    - [哈夫曼huffman树（最优二叉树）](#%E5%93%88%E5%A4%AB%E6%9B%BChuffman%E6%A0%91%E6%9C%80%E4%BC%98%E4%BA%8C%E5%8F%89%E6%A0%91)
      - [huffman概念](#huffman%E6%A6%82%E5%BF%B5)
      - [构造huffman树](#%E6%9E%84%E9%80%A0huffman%E6%A0%91)
      - [huffman编码](#huffman%E7%BC%96%E7%A0%81)
  - [B树](#b%E6%A0%91)
    - [多叉排序树](#%E5%A4%9A%E5%8F%89%E6%8E%92%E5%BA%8F%E6%A0%91)
    - [5阶B树特性](#5%E9%98%B6b%E6%A0%91%E7%89%B9%E6%80%A7)
    - [5阶B树插入和删除](#5%E9%98%B6b%E6%A0%91%E6%8F%92%E5%85%A5%E5%92%8C%E5%88%A0%E9%99%A4)
      - [插入](#%E6%8F%92%E5%85%A5)
      - [删除](#%E5%88%A0%E9%99%A4)
  - [B+树----B树变形树](#b%E6%A0%91----b%E6%A0%91%E5%8F%98%E5%BD%A2%E6%A0%91)
    - [b+树考题](#b%E6%A0%91%E8%80%83%E9%A2%98)
  - [B树 V B+树](#b%E6%A0%91-v-b%E6%A0%91)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# 树

![](.btree_images/tree.png)

## 基本术语

![](.btree_images/tree_words.png)
![](.btree_images/tree_words2.png)

## 树的性质

![](.btree_images/tree_attribute.png)
![](.btree_images/tree_attribute1.png)
![](.btree_images/tree_attribute2.png)
![](.btree_images/tree_attribute3.png)

## binary_tree二叉树

![](.btree_images/binary_tree.png)

### 性质

![](.btree_images/binary_attribute1.png)
![](.btree_images/binart_attribute2.png)

## 二叉树和度为2的树的区别

![](.btree_images/binary_tree_vs_other_tree.png)

### 满二叉树

![](.btree_images/full_binary_tree.png)

### 完全二叉树

![](.btree_images/whole_binary_tree.png)

### 二叉排序树

![](.btree_images/sorted_binary_tree.png)

## 二叉树的存储结构

### 完全二叉树的顺序存储

![](.btree_images/full_binary_tree_sequence_store.png)
![](.btree_images/full_binary_store.png)
用顺序存储的方法普通二叉树容易造成空间的浪费

### 二叉树的链式存储

![](.btree_images/binary_chain_store.png)

## 二叉树的遍历

![](.btree_images/read_binary.png)

### 层次遍历

![](.btree_images/layer_read.png)
![](03_dataStructure/03_btree/.btree_images/layer_read_process.png)
借助队列，出队时，帮左右子元素入队。

#### 过程

1. 使用辅助队列，根节点1入队
   ![](.btree_images/layer_process1.png)

2. 来个循环，出个节点1，把他的所有子节点2，3，4入队
   ![](.btree_images/layer_process2.png)

3. 同理，出队2，把2的子节点5，6入队，直到队列为空
   ![](.btree_images/layer_process3.png)

### 三种遍历方式

![](.btree_images/three_ways_of_read_binary_tree.png)

1. 先序遍历
   ![](.btree_images/left_root_right.png)
   ![](.btree_images/left_root_right1.png)
   ![](.btree_images/tips_of_left_root_right.png)

2. 中序遍历
   ![](.btree_images/left_root_right2.png)
   ![](.btree_images/tips_of_left_root_right1.png)

3. 后序遍历
   ![](.btree_images/left_root_right3.png)
   ![](.btree_images/tips_of_left_root_right2.png)

#### 遍历方式代码实现

递归方案实现先序遍历
![](.btree_images/iterate_preorder.png)

递归方案实现中序遍历
![](.btree_images/iterate_inorder.png)

递归方案实现后序遍历
![](.btree_images/iterate_postoder.png)

栈的方式实现中序遍历
![](.btree_images/stack_oder.png)

## 二叉树的遍历

由遍历序列构造二叉树
![](03_dataStructure/03_btree/.btree_images/construct_tree.png)

先序+中序遍历 构造二叉树
![](.btree_images/construct_tree1.png)
![](.btree_images/construct_tree2.png)

层次+中序遍历
![](.btree_images/construct_tree3.png)

## 线索二叉树

![](.btree_images/index_tree.png)
线索二叉树，可以理解是用了二叉树中的空指针。

二叉树线索化
![](.btree_images/index_tree2.png)

### 线索二叉树数据结构

![](.btree_images/threaded_tree_struct.png)

### 线索二叉树求前继和后继

![](.btree_images/threaded_tree_pre_post_order.png)

#### 先序线索二叉树求后继

![](.btree_images/preorder_get_postorder.png)

#### 中序线索二叉树求前继和后继

![](.btree_images/inorder_tree_pre_post.png)

## 树的存储结构

![](.btree_images/tree_storage.png)

### 1. 孩子兄弟表示法----主要

![](.btree_images/tree_storage1.png)
![](.btree_images/tree_storage2.png)

### 2. 双亲表示法

![](.btree_images/tree_storage3.png)

### 3. 孩子表示法

![](.btree_images/tree_storage4.png)

## 二叉排序（搜索）树(binary sort tree)

![](.btree_images/two_trees_info.png)
![](.btree_images/bst.png)

中序遍历---升序

### 查找,插入,创建,删除

![](.btree_images/search_bst.png)
![](.btree_images/insert_bst.png)

![](.btree_images/init_bst3.png)
![](.btree_images/init_bst.png)
![](.btree_images/init_bst2.png)

二叉排序树的节点删除
![](.btree_images/delete_ast.png)
![](.btree_images/delete_ast_node.png)

### 二叉排序树的查找效率分析

![](.btree_images/ast_performance_of_searching.png)
查找成功 的平均查找长度（ASL average search length）

![](.btree_images/ast_performance_of_searching_fail.png)
查找失败 的平均查找长度（ASL average search length）

![](.btree_images/ast_performance_of_searching_fail_n_suc.png)

## 平衡二叉树 ---代码实现过于复杂

![](.btree_images/balance_tree.png)
![](.btree_images/avl.png)

### 插入和删除 数据后不平衡

![](.btree_images/unbalanced_avl.png)

调整方式：调整最小不平衡子树，如70是最小的
![](.btree_images/sort_avl.png)

从最小不平衡子树出发   
![](.btree_images/sort_avl2.png)
![](.btree_images/sort_avl3.png)
![](.btree_images/sort_avl4.png)

1. 左左更高
   ![](.btree_images/sort_avl5.png)
   向右旋转，然后30放在右侧左子树
   ![](.btree_images/sort_avl6.png)
   ![](.btree_images/sort_avl7.png)

2. 左右更高
   ![](.btree_images/sort_avl8.png)
   ![](.btree_images/sort_avl9.png)
   向左旋转, 向右旋转
   ![](.btree_images/sort_avl10.png)
   ![](.btree_images/sort_avl11.png)
   ![](.btree_images/sort_avl12.png)
   ![](.btree_images/sort_avl13.png)

总结：  
![](.btree_images/sort_avl14.png)

### 树高和节点关系

![](.btree_images/avl_height.png)
![](.btree_images/avl_height1.png)

### 哈夫曼huffman树（最优二叉树）

在通信上广泛应用

把经常访问放在跟根结点最近的地方
![](.btree_images/hf_tree.png)
![](.btree_images/hf_tree1.png)

二进制表示:书写简单，不容易泄密
![](.btree_images/hf_tree2.png)

1. 浪费空间
2. 人物长度固定

![](.btree_images/hf_tree3.png)

1. 长度不固定
2. 长度短

#### huffman概念

![](.btree_images/huffman_idea.png)
![](.btree_images/huffman_idea2.png)

#### 构造huffman树

![](.btree_images/construct_huffman.png)
![](.btree_images/construct_huffman2.png)
![](.btree_images/construct_huffman3.png)

#### huffman编码

![](.btree_images/huffman_encoding.png)
![](.btree_images/huffman_encoding2.png)
![](.btree_images/huffman_encoding3.png)

## B树
B树的基本思想是保证每个节点至少填满一半的关键字，以达到较好的性能。其优点是对于大规模数据的插入和删除操作，它仍然能够保持较好的性能，但缺点是每个节点都需要存储一定数量的指针，导致存储空间的浪费。

### 多叉排序树

![](.btree_images/multi_tree.png)
![](.btree_images/multi_tree1.png)
n个关键字--分为n+1个区间，注意线画在关键字中间

### 5阶B树特性

![](.btree_images/b_tree_info.png)
![](.btree_images/b_tree_info1.png)
![](.btree_images/btree_exam.png)
![](.btree_images/btree_exam1.png)

### 5阶B树插入和删除

#### 插入

![](.btree_images/btree_insert.png)
![](.btree_images/btree_insert1.png)
![](.btree_images/btree_insert2.png)
![](.btree_images/btree_insert3.png)
![](.btree_images/btree_insert4.png)
![](.btree_images/btree_insert5.png)
![](.btree_images/btree_insert6.png)

#### 删除

1. 删除69
   ![](.btree_images/btree_delete1.png)

2. 删除50,向左向右都可以借，但是不是直接借，需要通过父节点中转
   ![](.btree_images/btree_delete2.png)
   ![](.btree_images/btree_delete3.png)


3. 删除28
   ![](.btree_images/btree_delete4.png)
   ![](.btree_images/btree_delete5.png)
   ![](.btree_images/btree_delete6.png)
   ![](.btree_images/btree_delete7.png)
   ![](.btree_images/btree_delete8.png)

4. 删除70
   ![](.btree_images/btree_delete9.png)
   ![](.btree_images/btree_delete10.png)

## B+树----B树变形树

B+树在B树的基础上进行了改进，将所有的关键字都存储在叶子节点中，非叶子节点只存储关键字和指向子节点的指针。
这种设计能够大大减少存储空间的浪费，并且能够通过叶子节点间的链表快速定位到某个范围内的数据。
B+树在磁盘等外部存储设备上的表现比B树更优，因此被广泛应用于数据库和文件系统中

![](.btree_images/b+tree.png)
![](.btree_images/ b+tree4.png)

1. 查找:可以从根节点通过分块查找，也可以通过叶节点采用顺序查找
2. 索引是是子树最大值,线连在子树中间

关键字和记录是分开的
![](.btree_images/b+tree1.png)

数据比较多，增加索引表
![](.btree_images/b+tree2.png)
![](.btree_images/b+tree3.png)

### b+树考题

![](.btree_images/b+tree_exam.png)

## B树 V B+树

![](.btree_images/btree_vs_b+tree.png)

