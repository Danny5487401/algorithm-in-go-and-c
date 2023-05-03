/*
 * 程序名：seqqueue4.c，此程序演示循环队列的数组实现，队尾指针指向队尾元素，增加了length的辅助变量。
 * 作者：C语言技术网(www.freecplus.net) 日期：20201230
*/
#include <stdio.h>
#include <string.h>

#define MAXSIZE 10       // 循环队列的最大长度，最多可以存放MAXSIZE个元素。

typedef int ElemType;    // 自定义循环队列的数据元素为整数。

typedef struct
{
  ElemType data[MAXSIZE];   // 用数组存储循环队列中的元素。
  int front;                // 队列的头指针。
  int rear;                 // 队列的尾指针，指向队尾元素。
  int length;               // 队列的实际长度。    // xxx
}SeqQueue,*PSeqQueue;

// 循环队列QQ的初始化操作。
void InitQueue(PSeqQueue QQ);                     

// 销毁循环队列QQ。
void DestroyQueue(PSeqQueue QQ);

// 元素入队，返回值：0-失败；1-成功。
int InQueue(PSeqQueue QQ, ElemType *ee);

// 元素出队，返回值：0-失败；1-成功。
int OutQueue(PSeqQueue QQ, ElemType *ee);

// 求循环队列的长度，返回值：>=0-队列QQ元素的个数。
int  Length(PSeqQueue QQ);                   

// 清空循环队列。
void Clear(PSeqQueue QQ);                    

// 判断循环队列是否为空，返回值：1-空，0-非空或失败。
int  IsEmpty(PSeqQueue QQ);                    

// 判断循环队列是否已满，返回值：1-已满，0-未满或失败。
int IsFull(PSeqQueue QQ);

// 打印循环队列中全部的元素。
void PrintQueue(PSeqQueue QQ);                    

// 获取队头元素，返回值：0-失败；1-成功。
// 只查看队头元素的值，元素不出队。
int GetHead(PSeqQueue QQ, ElemType *ee);

int main()
{
  SeqQueue QQ;     // 创建循环队列。

  InitQueue(&QQ);  // 初始化循环队列。

  ElemType ee;     // 创建一个数据元素。

  printf("元素（1、2、3、4、5、6、7、8、9、10、11）入队。\n");
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
  ee=11; InQueue(&QQ, &ee);

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

  return 0;
}

// 初始化循环队列
void InitQueue(PSeqQueue QQ)
{
  Clear(QQ); // 清空循环队列。
}

// 清空循环队列。
void Clear(PSeqQueue QQ)
{
  if (QQ == NULL) return;  // 检查空指针。

  QQ->front=0;
  QQ->rear=MAXSIZE-1;    // xxxx
  QQ->length=0;
  memset(QQ->data,0,sizeof(ElemType)*MAXSIZE);  // 数组元素清零。
}

// 求循环队列的长度，返回值：>=0-队列QQ元素的个数。
int Length(PSeqQueue QQ)
{
  if (QQ == NULL) return 0;  // 检查空指针。

  return QQ->length;    // xxx
}

// 销毁循环队列QQ。
void DestroyQueue(PSeqQueue QQ)
{
  // 静态循环队列无需释放内存，不需要销毁操作。

  Clear(QQ); // 清空循环队列。

  return;
}

// 判断循环队列是否为空，返回值：1-空，0-非空或失败。
int IsEmpty(PSeqQueue QQ)
{
  if (QQ == NULL) return 0;   // 检查空指针。

  // if (QQ->front == QQ->rear) return 1;  

  if (QQ->length == 0) return 1;    // xxx

  return 0;
}

// 判断循环队列是否已满，返回值：1-已满，0-未满或失败。
int IsFull(PSeqQueue QQ)
{
  if (QQ == NULL) return 0;   // 检查空指针。

  // if ( ((QQ->rear+1)%MAXSIZE) == QQ->front) return 1;
  if (QQ->length == MAXSIZE) return 1;    // xxx

  return 0;
}

// 元素入队，返回值：0-失败；1-成功。
int InQueue(PSeqQueue QQ, ElemType *ee)
{
  if ( (QQ == NULL) || (ee == NULL) ) return 0;   // 检查空指针。

  if (IsFull(QQ) == 1)
  {
    printf("循环队列已满，不能插入。\n"); return 0;
  }

  // xxxx 先移动队尾指针，然后再插入数据。
  QQ->rear=(QQ->rear+1)%MAXSIZE;  // 队尾指针后移。

  memcpy(&QQ->data[QQ->rear],ee,sizeof(ElemType));  // 用数组的下标访问。
  // memcpy(QQ->data+QQ->rear,ee,sizeof(ElemType));    // 采用指针运算也可以。

  QQ->length++;    // xxx

  return 1;
}

// 打印循环队列中全部的元素。
void PrintQueue(PSeqQueue QQ)
{
  if (QQ == NULL) return;   // 检查空指针。

  if (IsEmpty(QQ) == 1) { printf("队列为空。\n"); return; }

  int kk,qlen=Length(QQ);
  for (kk = 0; kk < qlen; kk++)
  {
    // 用数组的下标访问。
    printf("data[%d],value=%d\n",(QQ->front+kk)%MAXSIZE,QQ->data[(QQ->front+kk)%MAXSIZE]);     
   
    // 采用指针运算也可以。
    // printf("data[%d],value=%d\n",(QQ->front+kk)%MAXSIZE,*(QQ->data+(QQ->front+kk)%MAXSIZE));
  }
}

// 元素出队，返回值：0-失败；1-成功。
int OutQueue(PSeqQueue QQ, ElemType *ee)
{
  if ( (QQ == NULL) || (ee == NULL) ) return 0;   // 检查空指针。

  if (IsEmpty(QQ) == 1) { printf("队列为空。\n"); return 0; }

  memcpy(ee,&QQ->data[QQ->front],sizeof(ElemType));  // 用数组的下标访问。
  // memcpy(ee,QQ->data+QQ->front,sizeof(ElemType));    // 采用指针运算也可以。

  QQ->front=(QQ->front+1)%MAXSIZE;  // 队列头指针后移。

  QQ->length--;    // xxx

  return 1;
}

// 获取队头元素，返回值：0-失败；1-成功。
// 只查看队头元素的值，元素不出队。
int GetHead(PSeqQueue QQ, ElemType *ee)
{
  if ( (QQ == NULL) || (ee == NULL) ) return 0;   // 检查空指针。

  if (IsEmpty(QQ) == 1) { printf("队列为空。\n"); return 0; }

  memcpy(ee,&QQ->data[QQ->front],sizeof(ElemType));  // 用数组的下标访问。
  // memcpy(ee,QQ->data+QQ->front,sizeof(ElemType));    // 采用指针运算也可以。

  return 1;
}
