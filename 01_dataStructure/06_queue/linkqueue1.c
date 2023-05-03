/*
 * 程序名：linkqueue1.c，此程序演示队列的链表实现（带头结点）。
 * 作者：C语言技术网(www.freecplus.net) 日期：20201230
*/
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef int ElemType;    // 自定义队列的数据元素为整数。

typedef struct LNode
{
  ElemType data;         // 存储队列中的元素。
  struct LNode *next;    // next指针。
}LNode;

typedef struct
{
  LNode *front,*rear;     // 队列的头指针和尾指针。
}LinkQueue,*PLinkQueue;

// 队列QQ的初始化操作。
int InitQueue(PLinkQueue QQ);                     

// 销毁队列QQ。
void DestroyQueue(PLinkQueue QQ);

// 清空队列。
void Clear(PLinkQueue QQ);                    

// 元素入队，返回值：0-失败；1-成功。
int InQueue(PLinkQueue QQ, ElemType *ee);

// 打印队列中全部的元素。
void PrintQueue(PLinkQueue QQ);                    

// 求队列的长度，返回值：>=0-队列QQ元素的个数。
int  Length(PLinkQueue QQ);                   

// 判断队列是否已满，链式队列不存在队满的说法。
int IsFull(PLinkQueue QQ);

// 判断队列是否为空，返回值：1-空，0-非空或失败。
int  IsEmpty(PLinkQueue QQ);                    

// 元素出队，返回值：0-失败；1-成功。
int OutQueue(PLinkQueue QQ, ElemType *ee);

// 获取队头元素，返回值：0-失败；1-成功。
// 只查看队头元素的值，元素不出队。
int GetHead(PLinkQueue QQ, ElemType *ee);

int main()
{
  LinkQueue QQ;     // 创建队列。

  memset(&QQ,0,sizeof(LinkQueue));

  InitQueue(&QQ);  // 初始化队列。

  ElemType ee;     // 创建一个数据元素。

  printf("元素（1、2、3、4、5、6、7、8、9、10）入队。\n");
  ee=1;  InQueue(&QQ, &ee);
  ee=2;  InQueue(&QQ, &ee);
  ee=3;  InQueue(&QQ, &ee);
  ee=4;  InQueue(&QQ, &ee);
  ee=5;  InQueue(&QQ, &ee);
  ee=6;  InQueue(&QQ, &ee);
  ee=7;  InQueue(&QQ, &ee);
  ee=8;  InQueue(&QQ, &ee);
  ee=9;  InQueue(&QQ, &ee);
  ee=10; InQueue(&QQ, &ee);

  printf("队列的长度是%d\n",Length(&QQ));

  PrintQueue(&QQ);

  if (OutQueue(&QQ,&ee)==1)  printf("出队的元素值为%d\n",ee);
  if (OutQueue(&QQ,&ee)==1)  printf("出队的元素值为%d\n",ee);
  if (OutQueue(&QQ,&ee)==1)  printf("出队的元素值为%d\n",ee);
  if (OutQueue(&QQ,&ee)==1)  printf("出队的元素值为%d\n",ee);
  if (OutQueue(&QQ,&ee)==1)  printf("出队的元素值为%d\n",ee);
  if (OutQueue(&QQ,&ee)==1)  printf("出队的元素值为%d\n",ee);
  if (OutQueue(&QQ,&ee)==1)  printf("出队的元素值为%d\n",ee);

  printf("队列的长度是%d\n",Length(&QQ));
  PrintQueue(&QQ);

  printf("元素（11、12、13、14、15）入队。\n");
  ee=11;  InQueue(&QQ, &ee);
  ee=12;  InQueue(&QQ, &ee);
  ee=13;  InQueue(&QQ, &ee);
  ee=14;  InQueue(&QQ, &ee);
  ee=15;  InQueue(&QQ, &ee);

  printf("队列的长度是%d\n",Length(&QQ));

  PrintQueue(&QQ);

  // 只查看队头元素的值，元素不出队。
  if (GetHead(&QQ,&ee)==1)  printf("队头的元素值为%d\n",ee);

  DestroyQueue(&QQ); // 销毁队列QQ。

  return 0;
}

// 初始化队列，返回值：0-失败；1-成功。
int InitQueue(PLinkQueue QQ)
{
  LNode *head = (LNode *)malloc(sizeof(LNode));  // 分配头结点。

  if (head == NULL) return 0;  // 内存不足，返回失败。

  head->next=NULL;  // 头结点的下一结点暂时不存在，置空。

  QQ->front=QQ->rear=head;

  return 1;
}

// 清空队列。
void Clear(PLinkQueue QQ)
{
  if (QQ == NULL) return; // 判断队列是否存在。

  if (QQ->front == NULL) { printf("队列QQ未初始化。\n"); return; } // 队列未初始化。

  // 清空队列QQ是指释放链表全部的数据结点，但保留头结点。
  LNode *tmp=QQ->front->next,*tmpnext;

  while(tmp!=NULL)
  {
    tmpnext=tmp->next;  // tmpnext保存下一结点的地址。
    free(tmp);          // 释放当前结点。
    tmp=tmpnext;        // tmp指针移动到下一结点。
  }

  QQ->rear=QQ->front;   // 尾指针指向头结点。

  QQ->front->next=NULL;
}

// 求队列的长度，返回值：>=0-队列QQ元素的个数。
int Length(PLinkQueue QQ)
{
  if (QQ == NULL) return 0;  // 检查空指针。

  if (QQ->front == NULL) { printf("队列QQ未初始化。\n"); return 0; } // 队列未初始化。

  LNode *pp=QQ->front->next;  // 头结点不算，从第1个结点开始。

  int length=0;

  while (pp != NULL) { pp=pp->next; length++; }

  return length;

}

// 销毁队列QQ。
void DestroyQueue(PLinkQueue QQ)
{
  if (QQ == NULL) return;    // 检查空指针。

  // 销毁队列QQ是指释放链表全部的结点，包括头结点。
  LNode *tmp=QQ->front,*tmpnext;

  while(tmp!=NULL)
  {
    tmpnext=tmp->next;  // tmpnext保存下一结点的地址。
    free(tmp);          // 释放当前结点。
    tmp=tmpnext;        // tmp指针移动到下一结点。
  }

  QQ->front=QQ->rear=NULL;  // 防止野指针。

  return;
}

// 判断队列是否为空，返回值：1-空，0-非空或失败。
int IsEmpty(PLinkQueue QQ)
{
  if (QQ == NULL) return 0;  // 检查空指针。

  if (QQ->front == NULL) { printf("队列QQ未初始化。\n"); return 0; } // 队列未初始化。

  if (QQ->front->next == NULL) return 1;

  return 0;
}

// 判断队列是否已满，链式队列不存在队满的说法。
int IsFull(PLinkQueue QQ)
{
  if (QQ == NULL) return 0;   // 检查空指针。

  return 1;
}

// 元素入队，返回值：0-失败；1-成功。
int InQueue(PLinkQueue QQ, ElemType *ee)
{
  if ( (QQ == NULL) || (ee == NULL) ) return 0;   // 检查空指针。

  if (QQ->front == NULL) { printf("队列QQ未初始化。\n"); return 0; } // 队列未初始化。

  LNode *tmp = (LNode *)malloc(sizeof(LNode));  // 分配一个结点。
  if (tmp == NULL) return 0;  // 内存不足，返回失败。

  // 考虑数据元素为结构体的情况，这里采用了memcpy的方法而不是直接赋值。
  memcpy(&tmp->data,ee,sizeof(ElemType));
  tmp->next=NULL;

  // 把tmp追加到QQ->rear之后。
  QQ->rear->next=tmp;
  QQ->rear=tmp;

  return 1;
}

// 打印队列中全部的元素。
void PrintQueue(PLinkQueue QQ)
{
  if (QQ == NULL) return;   // 检查空指针。

  if (QQ->front == NULL) { printf("队列QQ未初始化。\n"); return; } // 队列未初始化。

  LNode *pp=QQ->front->next;  // 从第1个数据结点开始。

  while (pp != NULL)
  {
    printf("%-3d", pp->data);  // 如果元素ee为结构体，这行代码要修改。
    pp=pp->next;
  }

  printf("\n");
}

// 元素出队，返回值：0-失败；1-成功。
int OutQueue(PLinkQueue QQ, ElemType *ee)
{
  if ( (QQ == NULL) || (ee == NULL) ) return 0;   // 检查空指针。

  if (QQ->front == NULL) { printf("队列QQ未初始化。\n"); return 0; } // 队列未初始化。

  if (IsEmpty(QQ) == 1) { printf("队列为空。\n"); return 0; }

  LNode *tmp=QQ->front->next;

  memcpy(ee,&tmp->data,sizeof(ElemType));  
  QQ->front->next=tmp->next;

  // 如果出队的是最后一个结点。
  if (QQ->rear==tmp) QQ->rear=QQ->front;
  
  free(tmp);

  return 1;
}

// 获取队头元素，返回值：0-失败；1-成功。
// 只查看队头元素的值，元素不出队。
int GetHead(PLinkQueue QQ, ElemType *ee)
{
  if ( (QQ == NULL) || (ee == NULL) ) return 0;   // 检查空指针。

  if (QQ->front == NULL) { printf("队列QQ未初始化。\n"); return 0; } // 队列未初始化。

  if (IsEmpty(QQ) == 1) { printf("队列为空。\n"); return 0; }

  memcpy(ee,&QQ->front->next->data,sizeof(ElemType));  

  return 1;
}
