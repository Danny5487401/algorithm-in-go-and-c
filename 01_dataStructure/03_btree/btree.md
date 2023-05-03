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

