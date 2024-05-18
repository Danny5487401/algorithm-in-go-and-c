<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Binary Search](#binary-search)
  - [在排序数组中查找元素的第一个和最后一个位置](#%E5%9C%A8%E6%8E%92%E5%BA%8F%E6%95%B0%E7%BB%84%E4%B8%AD%E6%9F%A5%E6%89%BE%E5%85%83%E7%B4%A0%E7%9A%84%E7%AC%AC%E4%B8%80%E4%B8%AA%E5%92%8C%E6%9C%80%E5%90%8E%E4%B8%80%E4%B8%AA%E4%BD%8D%E7%BD%AE)
  - [[寻找峰值]](#%E5%AF%BB%E6%89%BE%E5%B3%B0%E5%80%BC)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Binary Search


## [在排序数组中查找元素的第一个和最后一个位置](./34_find_first_and_last_position_of_element_in_sorted_array_test.go)

返回有序数组 >= 8 的第一个位置

![](.binary_search_images/binary_search1.png)
![](.binary_search_images/binary_search2.png)
![](.binary_search_images/binary_search3.png)

循环不变量：
- L-1 红色：小于8 
- R+1 蓝色：大于等于8

循环结束后，因为 R+1 = L，所以也可以用 L 表示


四种情况
* 大于等于
* 大于
* 小于
* 小于等于


## [寻找峰值](162_find_peak_element_test.go)
![](.binary_search_images/Find Peak Element1.png)
![](.binary_search_images/Find Peak Element2.png)

- 红色：峰顶左侧元素
- 蓝色：峰顶及其右侧元素
- 白色：不确定

根据此定义，最右侧 6 肯定是蓝色

题目 任意两个相邻数不相同

## [寻找旋转排序数组的最小值]

两段式递增，寻找最小值 

![](.binary_search_images/Find Minimum in Rotated Sorted Array1.png)
![](.binary_search_images/Find Minimum in Rotated Sorted Array2.png)




## 参考

- [二分查找——红蓝染色法](https://blog.csdn.net/qq_45808700/article/details/129247507)