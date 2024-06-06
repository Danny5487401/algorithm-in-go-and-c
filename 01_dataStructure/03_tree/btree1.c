/*
 * 程序名：btree1.c，此程序演示二叉树的层次遍历。
 * 作者：C语言技术网(www.freecplus.net) 日期：20200202
*/
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

///////////////////////////////////////////////
// 二叉树的数据结构。
typedef struct BiTNode
{
  char   data;               // 存放结点的数据元素。
  struct BiTNode *lchild;    // 指向左子结点地址的指针。
  struct BiTNode *rchild;    // 指向右子结点地址的指针。
}BiTNode,*BiTree;
///////////////////////////////////////////////


///////////////////////////////////////////////
// 队列的数据结构。
#define MAXSIZE 30       // 循环队列的最大长度，最多可以存放MAXSIZE-1个元素。

typedef BiTree ElemType;     // 自定义队列的数据元素为二叉树。

typedef struct
{
  ElemType data[MAXSIZE];   // 用数组存储循环队列中的元素。
  int front;                // 队列的头指针。
  int rear;                 // 队列的尾指针，指向队尾的下一个元素。
}SeqQueue,*PSeqQueue;
///////////////////////////////////////////////

///////////////////////////////////////////////
// 队列操作的函数。

// 循环队列QQ的初始化操作。
void InitQueue(PSeqQueue QQ);

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
///////////////////////////////////////////////


// 二叉树的层次遍历。
void LevelOrder(BiTree TT);

int main()
{
  BiTree TT=0; // 声明树指针变量。

  /*
  // 手工构造一个如下结构的二叉树。
             1
          /     \
         2       3
        / \     /
       4   5   6
      / \ / \
     7  8 9  0
  */

  // 分配根节点。
  TT=(BiTNode *)malloc(sizeof(BiTNode));
  TT->data='1';
  TT->lchild=TT->rchild=0;

  // 第二层第一个节点。
  TT->lchild=(BiTNode *)malloc(sizeof(BiTNode));
  TT->lchild->data='2'; TT->lchild->lchild=TT->lchild->rchild=0;

  // 第二层第二个节点。
  TT->rchild=(BiTNode *)malloc(sizeof(BiTNode));
  TT->rchild->data='3'; TT->rchild->lchild=TT->rchild->rchild=0;

  // 第三层第一个节点。
  TT->lchild->lchild=(BiTNode *)malloc(sizeof(BiTNode));
  TT->lchild->lchild->data='4'; TT->lchild->lchild->lchild=TT->lchild->lchild->rchild=0;

  // 第三层第二个节点。
  TT->lchild->rchild=(BiTNode *)malloc(sizeof(BiTNode));
  TT->lchild->rchild->data='5'; TT->lchild->rchild->lchild=TT->lchild->rchild->rchild=0;

  // 第三层第三个节点。
  TT->rchild->lchild=(BiTNode *)malloc(sizeof(BiTNode));
  TT->rchild->lchild->data='6'; TT->rchild->lchild->lchild=TT->rchild->lchild->rchild=0;

  // 第四层第一个节点。
  TT->lchild->lchild->lchild=(BiTNode *)malloc(sizeof(BiTNode));
  TT->lchild->lchild->lchild->data='7'; TT->lchild->lchild->lchild->lchild=TT->lchild->lchild->lchild->rchild=0;

  // 第四层第二个节点。
  TT->lchild->lchild->rchild=(BiTNode *)malloc(sizeof(BiTNode));
  TT->lchild->lchild->rchild->data='8'; TT->lchild->lchild->rchild->lchild=TT->lchild->lchild->rchild->rchild=0;

  // 第四层第三个节点。
  TT->lchild->rchild->lchild=(BiTNode *)malloc(sizeof(BiTNode));
  TT->lchild->rchild->lchild->data='9'; TT->lchild->rchild->lchild->lchild=TT->lchild->rchild->lchild->rchild=0;

  // 第四层第四个节点。
  TT->lchild->rchild->rchild=(BiTNode *)malloc(sizeof(BiTNode));
  TT->lchild->rchild->rchild->data='0'; TT->lchild->rchild->rchild->lchild=TT->lchild->rchild->rchild->rchild=0;

  // 二叉树层次遍历。
  printf("层次遍历结果："); LevelOrder(TT); printf("\n");

  return 0;
}

// 访问出队元素。
void visit(BiTNode *pp)
{
  printf("%c ",pp->data);  // 访问出队元素就是把值输出来，意思一下就行了。
}

// 二叉树层次遍历。
void LevelOrder(BiTree TT)
{
  SeqQueue QQ;     // 创建循环队列。

  InitQueue(&QQ);  // 初始化循环队列。

  ElemType ee=TT;  // 队列的元素是二叉树。
  
  InQueue(&QQ,&ee);  // 把根结点当成队列的元素入队。
  
  while (IsEmpty(&QQ)!=1)  // 队列不为空。
  {
    OutQueue(&QQ,&ee);  // 队头元素出队。

    visit(ee);  // 访问出队元素。

    if (ee->lchild != NULL) InQueue(&QQ,&ee->lchild);  // 如果出队元素有左结点，左结点入队。

    if (ee->rchild != NULL) InQueue(&QQ,&ee->rchild);  // 如果出队元素有右结点，右结点入队。
  }
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

  QQ->front=QQ->rear=0;
  memset(QQ->data,0,sizeof(ElemType)*MAXSIZE);  // 数组元素清零。
}

// 求循环队列的长度，返回值：>=0-队列QQ元素的个数。
int Length(PSeqQueue QQ)
{
  if (QQ == NULL) return 0;  // 检查空指针。

  return (QQ->rear-QQ->front+MAXSIZE)%MAXSIZE;
}

// 判断循环队列是否为空，返回值：1-空，0-非空或失败。
int IsEmpty(PSeqQueue QQ)
{
  if (QQ == NULL) return 0;    // 检查空指针。

  if (QQ->front == QQ->rear) return 1;

  return 0;
}

// 判断循环队列是否已满，返回值：1-已满，0-未满或失败。
int IsFull(PSeqQueue QQ)
{
  if (QQ == NULL) return 0;   // 检查空指针。

  if ( ((QQ->rear+1)%MAXSIZE) == QQ->front) return 1;

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

  memcpy(&QQ->data[QQ->rear],ee,sizeof(ElemType));  // 用数组的下标访问。
  // memcpy(QQ->data+QQ->rear,ee,sizeof(ElemType));    // 采用指针运算也可以。

  QQ->rear=(QQ->rear+1)%MAXSIZE;  // 队尾指针后移。

  return 1;
}

// 元素出队，返回值：0-失败；1-成功。
int OutQueue(PSeqQueue QQ, ElemType *ee)
{
  if ( (QQ == NULL) || (ee == NULL) ) return 0;   // 检查空指针。

  if (IsEmpty(QQ) == 1) { printf("队列为空。\n"); return 0; }

  memcpy(ee,&QQ->data[QQ->front],sizeof(ElemType));  // 用数组的下标访问。
  // memcpy(ee,QQ->data+QQ->front,sizeof(ElemType));    // 采用指针运算也可以。

  QQ->front=(QQ->front+1)%MAXSIZE;  // 队列头指针后移。

  return 1;
}

