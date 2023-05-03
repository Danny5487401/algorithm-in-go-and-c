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