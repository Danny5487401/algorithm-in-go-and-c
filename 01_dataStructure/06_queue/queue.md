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