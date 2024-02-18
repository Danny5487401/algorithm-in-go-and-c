<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [链表](#%E9%93%BE%E8%A1%A8)
  - [线性表和单链表对比](#%E7%BA%BF%E6%80%A7%E8%A1%A8%E5%92%8C%E5%8D%95%E9%93%BE%E8%A1%A8%E5%AF%B9%E6%AF%94)
  - [单链表分类](#%E5%8D%95%E9%93%BE%E8%A1%A8%E5%88%86%E7%B1%BB)
  - [单链表代码](#%E5%8D%95%E9%93%BE%E8%A1%A8%E4%BB%A3%E7%A0%81)
    - [创建和销毁](#%E5%88%9B%E5%BB%BA%E5%92%8C%E9%94%80%E6%AF%81)
    - [单链表的插入与删除(按位序---下标索引)](#%E5%8D%95%E9%93%BE%E8%A1%A8%E7%9A%84%E6%8F%92%E5%85%A5%E4%B8%8E%E5%88%A0%E9%99%A4%E6%8C%89%E4%BD%8D%E5%BA%8F---%E4%B8%8B%E6%A0%87%E7%B4%A2%E5%BC%95)
      - [插入](#%E6%8F%92%E5%85%A5)
      - [删除](#%E5%88%A0%E9%99%A4)
    - [查找](#%E6%9F%A5%E6%89%BE)
    - [单链表的插入与删除(按结点)](#%E5%8D%95%E9%93%BE%E8%A1%A8%E7%9A%84%E6%8F%92%E5%85%A5%E4%B8%8E%E5%88%A0%E9%99%A4%E6%8C%89%E7%BB%93%E7%82%B9)
  - [单链表的排序](#%E5%8D%95%E9%93%BE%E8%A1%A8%E7%9A%84%E6%8E%92%E5%BA%8F)
  - [单链表的反转](#%E5%8D%95%E9%93%BE%E8%A1%A8%E7%9A%84%E5%8F%8D%E8%BD%AC)
  - [双链表](#%E5%8F%8C%E9%93%BE%E8%A1%A8)
  - [循环单链表](#%E5%BE%AA%E7%8E%AF%E5%8D%95%E9%93%BE%E8%A1%A8)
  - [循环双链表](#%E5%BE%AA%E7%8E%AF%E5%8F%8C%E9%93%BE%E8%A1%A8)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# 链表
## 线性表和单链表对比
![](.link_list_images/liner_list_vs_link_list.png)


## 单链表分类
![](.link_list_images/link_list_class.png)
- 带头节点
- 不带头节点

## 单链表代码
![](.link_list_images/link_list_code.png)
这边别名注释：
LNode为结构体，LinkList是结构体指针

- LinkList LL ; 强调声明一个链表
- LNode *node ; 强调声明一个结点

### 创建和销毁
![](.link_list_images/link_list_code1.png)


### 单链表的插入与删除(按位序---下标索引)
![](.link_list_images/link_list_delete_n_insert.png)

#### 插入
![](.link_list_images/insert_process.png)
Note: 写法顺序不能反

#### 删除
![](.link_list_images/delete_process.png)

### 查找
- 按位序
- 按元素

### 单链表的插入与删除(按结点)
![](.link_list_images/insert_process1.png)


## 单链表的排序
![](.link_list_images/sort_link_list.png)

## 单链表的反转
![](.link_list_images/reverse_link_list.png)
- 反转部分
- 反转全部

![](.link_list_images/reverse_process.png)
方式一：原地反转链表流程
1. 取出所有数据
2. 然后一个一个插数据

方式二：新空间反转链表流程
![](.link_list_images/reverse_process_1.png)

## 双链表
![](.link_list_images/double_link_list.png)


## 循环单链表
![](.link_list_images/cycle_link_list.png)

## 循环双链表
![](.link_list_images/cycle_double_link_list.png)