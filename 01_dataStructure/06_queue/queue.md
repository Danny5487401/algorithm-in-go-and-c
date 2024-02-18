<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [queue队列](#queue%E9%98%9F%E5%88%97)
  - [队列的顺序存储](#%E9%98%9F%E5%88%97%E7%9A%84%E9%A1%BA%E5%BA%8F%E5%AD%98%E5%82%A8)
    - [队空和队满判断](#%E9%98%9F%E7%A9%BA%E5%92%8C%E9%98%9F%E6%BB%A1%E5%88%A4%E6%96%AD)
    - [队列的顺序存储代码实现](#%E9%98%9F%E5%88%97%E7%9A%84%E9%A1%BA%E5%BA%8F%E5%AD%98%E5%82%A8%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [队列的链式存储](#%E9%98%9F%E5%88%97%E7%9A%84%E9%93%BE%E5%BC%8F%E5%AD%98%E5%82%A8)
  - [双端队列--实际没有什么价值](#%E5%8F%8C%E7%AB%AF%E9%98%9F%E5%88%97--%E5%AE%9E%E9%99%85%E6%B2%A1%E6%9C%89%E4%BB%80%E4%B9%88%E4%BB%B7%E5%80%BC)
  - [队列的应用](#%E9%98%9F%E5%88%97%E7%9A%84%E5%BA%94%E7%94%A8)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# queue队列
![](.queue_images/queue.png)
![](.queue_images/queue2.png)

## 队列的顺序存储
![](.queue_images/sequnce_queue.png)
![](.queue_images/sequnce_queue2.png)
1. 出队，头指针往后移，从1-->3
2. 入队，尾指针往后移，可以从8-->9-->1

用数组实现的队列（循环队列）
![](.queue_images/sequence_queue_code.png)

### 队空和队满判断
![](.queue_images/sequence_queue_code2.png)
1. 通过牺牲一个存储单元

![](.queue_images/extra_length_sequence_code3.png)
2. 增加辅助变量 


### 队列的顺序存储代码实现
![](.queue_images/queue_code1.png)
![](.queue_images/queue_code2.png)
![](.queue_images/queue_code3.png)
![](.queue_images/queue_code4.png)


## 队列的链式存储
![](.queue_images/link_queue.png)

![](.queue_images/link_queue2.png)

入队
![](.queue_images/link_enqueue.png)

出队
![](.queue_images/link_dequeue.png)


## 双端队列--实际没有什么价值
![](.queue_images/double_end_queue.png)

## 队列的应用
1. 树的层次应用
![](.queue_images/btree_queue.png)
![](.queue_images/btree_queue2.png)

根节点8入队，把根节点的左右子节点3和2入队，然后根节点8出队。
把3节点的左右节点4和5入队，把3出队。把2节点的左右节点3和6入队等等

2. cpu的资源调度
![](.queue_images/queue_in_cpu.png)

3. 数据缓冲区
![](.queue_images/queue_in_buffer.png)