# 查找
![](.search_images/search_idea.png)

## 顺序查找
![](.search_images/sequence_search.png)
![](.search_images/sequence_search2.png)

### 优化方式
1. 哨兵
![](.search_images/sentinal_search.png)
![](.search_images/sentinal_search1.png)

2. 先排序
![](.search_images/sort_search.png)

3. 跳跃
![](.search_images/skip_search.png)


## 查找表的效率
1. 被查找的概率相同  
![](.search_images/search_table_performance.png)
![](.search_images/sorted_search_table_performance.png)

查找判定树来分析效率  
![](.search_images/search_tree_performance.png)

跳跃查找判定树
![](.search_images/skip_list_search_tree1.png)
![](.search_images/skip_list_search_tree2.png)
- ASL成功：没有跳跃5 跳跃3.44
- ASL失败：没有跳跃5.4 跳跃4

2. 被查找的概率不同
![](.search_images/search_table_ratio.png)


## 折半查找
![](.search_images/binary_search.png)
![](.search_images/binary_search1.png)
![](.search_images/binary_search2.png)
![](.search_images/binary_search3.png)
没有找到：high>low

### 效率分析
![](.search_images/binary_search_performance.png)
![](.search_images/binary_search_performance1.png) 
![](.search_images/binary_search_performance2.png)


## 分块查找
![](.search_images/block_search.png)
![](.search_images/block_search1.png)
![](.search_images/binary_block_search2.png)
块内无序，块间有序

### 效率分析
![](.search_images/block_search_perfomance.png)
