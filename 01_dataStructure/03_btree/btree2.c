/*
 * 程序名：btree2.c，此程序演示二叉树的前序遍历、中序遍历和后序遍历，包括递归和非递归两种方法。
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


////////////////////////////////////////////////////////////////
// 栈的数据结构。
#define MAXSIZE 30       // 顺序栈的最大长度。

typedef BiTree ElemType;     // 自定义栈的数据元素为二叉树。

typedef struct
{
  ElemType data[MAXSIZE];   // 用数组存储顺序栈中的元素。
  int top;                  // 栈顶指针，从0到MAXSIZE-1，-1表示空栈。
                            // 也可以从1到MAXSIZE，0表示空栈。
}SeqStack,*PSeqStack;
////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////
// 顺序栈SS的初始化操作。
void InitStack(PSeqStack SS);                     

// 元素入栈，返回值：0-失败；1-成功。
int Push(PSeqStack SS, ElemType *ee);

// 元素出栈，返回值：0-失败；1-成功。
int Pop(PSeqStack SS, ElemType *ee);

// 求顺序栈的长度，返回值：栈SS中元素的个数。
int Length(PSeqStack SS);                   

// 清空顺序栈。
void Clear(PSeqStack SS);                    

// 判断顺序栈是否为空，返回值：1-空，0-非空或失败。
int IsEmpty(PSeqStack SS);                    

// 判断顺序栈是否已满，返回值：1-已满，0-未满或失败。
int IsFull(PSeqStack SS);

// 打印顺序栈中全部的元素。
void PrintStack(PSeqStack SS);                    

// 获取栈顶元素，返回值：0-失败；1-成功。
// 只查看栈顶元素的值，元素不出栈。
int GetTop(PSeqStack SS, ElemType *ee);
////////////////////////////////////////////////////////////////

// 求二叉树的高度。
int TreeDepth(BiTree TT);

// 访问结点元素。
void visit(BiTNode *pp);

// 采用递归的方法对二叉树的先序遍历。
void PreOrder(BiTree TT);
// 不采用递归的方法对二叉树的先序遍历。
void PreOrder1(BiTree TT);

// 采用递归的方法对二叉树的中序遍历。
void InOrder(BiTree TT);
// 不采用递归的方法对二叉树的中序遍历。
void InOrder1(BiTree TT);

// 采用递归的方法对二叉树的后序遍历。
void PostOrder(BiTree TT);
// 不采用递归的方法对二叉树的后序遍历。
void PostOrder1(BiTree TT);

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

  // 二叉树的先序遍历。
  printf("先序遍历结果1："); PreOrder(TT); printf("\n");
  printf("先序遍历结果2："); PreOrder1(TT); printf("\n");

  // 二叉树的中序遍历。
  printf("中序遍历结果1："); InOrder(TT); printf("\n");
  printf("中序遍历结果2："); InOrder1(TT); printf("\n");

  // 二叉树的后序遍历。
  printf("后序遍历结果1："); PostOrder(TT); printf("\n");
  printf("后序遍历结果2："); PostOrder1(TT); printf("\n");


  // 求二叉树的高度。
  printf("树的高度是：%d\n",TreeDepth(TT));

  return 0;
}

// 初始化顺序栈
void InitStack(PSeqStack SS)
{
  Clear(SS); // 清空顺序栈。
}

// 清空顺序栈。
void Clear(PSeqStack SS)
{
  if (SS == NULL) return; // 检查空指针。

  SS->top=-1;  // 栈顶指针置为-1。
  memset(SS->data,0,sizeof(ElemType)*MAXSIZE);  // 数组元素清零。
}

// 求顺序栈的长度，返回值：栈SS中元素的个数。
int Length(PSeqStack SS)
{
  if (SS == NULL) return 0; // 检查空指针。

  return SS->top+1;
}


// 判断顺序栈是否为空，返回值：1-空，0-非空或失败。
int IsEmpty(PSeqStack SS)
{
  if (SS == NULL) return 0;  // 检查空指针。

  if (SS->top == -1) return 1;

  return 0;
}

// 判断顺序栈是否已满，返回值：1-已满，0-未满或失败。
int IsFull(PSeqStack SS)
{
  if (SS == NULL) return 0;  // 检查空指针。

  if (SS->top >= MAXSIZE-1) return 1;

  return 0;
}

// 元素入栈，返回值：0-失败；1-成功。
int Push(PSeqStack SS, ElemType *ee)
{
  if ( (SS == NULL) || (ee == NULL) ) return 0;  // 检查空指针。

  if (IsFull(SS) == 1)
  {
    printf("顺序栈已满，不能插入。\n"); return 0;
  }

  SS->top++;  // 栈指针先加1。

  memcpy(&SS->data[SS->top],ee,sizeof(ElemType));  // 用数组的下标访问。
  // memcpy(SS->data+SS->top,ee,sizeof(ElemType));    // 采用指针运算也可以。

  return 1;
}

// 打印顺序栈中全部的元素。
void PrintStack(PSeqStack SS)
{
  if (SS == NULL) return;  // 检查空指针。

  if (SS->top == -1) { printf("栈为空。\n"); return; }

  int kk;
  for (kk = 0; kk <= SS->top; kk++)
  {
    printf("SS[%d],value=%d\n",kk,SS->data[kk]);     // 用数组的下标访问。
    // printf("SS[%d],value=%d\n",kk,*(SS->data+kk));   // 采用指针运算也可以。
  }
}

// 元素出栈，返回值：0-失败；1-成功。
int Pop(PSeqStack SS, ElemType *ee)
{
  if ( (SS == NULL) || (ee == NULL) ) return 0;  // 检查空指针。

  if (SS->top == -1) { printf("栈为空。\n"); return 0; }

  memcpy(ee,&SS->data[SS->top],sizeof(ElemType));  // 用数组的下标访问。
  // memcpy(ee,SS->data+SS->top,sizeof(ElemType));    // 采用指针运算也可以。

  SS->top--;  // 栈指针减1。

  return 1;
}

// 获取栈顶元素，返回值：0-失败；1-成功。
// 只查看栈顶元素的值，元素不出栈。
int GetTop(PSeqStack SS, ElemType *ee)
{
  if ( (SS == NULL) || (ee == NULL) ) return 0;  // 检查空指针。

  if (IsEmpty(SS) == 1) { printf("栈为空。\n"); return 0; }

  memcpy(ee,&SS->data[SS->top],sizeof(ElemType));  // 用数组的下标访问。
  // memcpy(ee,SS->data+SS->top,sizeof(ElemType));    // 采用指针运算也可以。

  return 1;
}

// 访问结点元素。
void visit(BiTNode *pp)
{
  printf("%c ",pp->data);  // 访问结点元素就是把值输出来，意思一下就行了。
}

// 二叉树的先序遍历。
void PreOrder(BiTree TT)
{
  if (TT == NULL) return;

  visit(TT);               // 访问子树TT的根结点。
  PreOrder(TT->lchild);    // 遍历左子树。
  PreOrder(TT->rchild);    // 遍历右子树。
}

// 二叉树的中序遍历。
void InOrder(BiTree TT)
{
  if (TT == NULL) return;

  InOrder(TT->lchild);     // 遍历左子树。
  visit(TT);               // 访问子树TT的根结点。
  InOrder(TT->rchild);     // 遍历右子树。
}

// 二叉树的后序遍历。
void PostOrder(BiTree TT)
{
  if (TT == NULL) return;

  PostOrder(TT->lchild);     // 遍历左子树。
  PostOrder(TT->rchild);     // 遍历右子树。
  visit(TT);                 // 访问子树TT的根结点。
}

// 不采用递归的方法对二叉树的先序遍历。
void PreOrder1(BiTree TT)
{
  SeqStack SS;     // 创建顺序栈。

  InitStack(&SS);  // 初始化顺序栈。

  ElemType ee=TT;  // 队列的元素是二叉树。

  while ( (ee!=NULL) || (IsEmpty(&SS)!=1) )
  {
    if (ee!=NULL)
    {
      visit(ee);               // 入栈前访问。
      Push(&SS,&ee);           // 入栈。
      ee=ee->lchild;           // 继续处理左子树。
    }
    else
    {
      Pop(&SS,&ee);            // 出栈。
      ee=ee->rchild;           // 处理右子树。 
    }
  }
}

// 不采用递归的方法对二叉树的中序遍历。
void InOrder1(BiTree TT)
{
  SeqStack SS;     // 创建顺序栈。

  InitStack(&SS);  // 初始化顺序栈。

  ElemType ee=TT;  // 队列的元素是二叉树。

  while ( (ee!=NULL) || (IsEmpty(&SS)!=1) )
  {
    if (ee!=NULL)
    {
      Push(&SS,&ee);        // 入栈。
      ee=ee->lchild;        // 继续处理左子树。
    }
    else
    {
      Pop(&SS,&ee);        // 出栈。
      visit(ee);           // 访问出栈元素。
      ee=ee->rchild;       // 处理右子树。
    }
  }
}

// 不采用递归的方法对二叉树的后序遍历。
void PostOrder1(BiTree TT)
{
  SeqStack SS;      // 创建顺序栈。

  InitStack(&SS);   // 初始化顺序栈。

  ElemType ee=TT;   // 队列的元素是二叉树。
  
  ElemType pp=NULL; // 最近访问过的结点。

  while ( (ee!=NULL) || (IsEmpty(&SS)!=1) )
  {
    if (ee!=NULL)    // 从根结点开始，沿树的左结点，依次入栈，直到左结点为空。
    {
      Push(&SS,&ee);
      ee=ee->lchild;
    }
    else   // 转向右。
    {
      // 读取栈顶元素，不出栈。
      GetTop(&SS,&ee);       

      // 如果右子树存在并且没有被访问过。
      if ( (ee->rchild!=NULL) && (ee->rchild!=pp) )
      {
        ee=ee->rchild;  // 转向右。
        Push(&SS,&ee);  // 右子树的根结点入栈。
        ee=ee->lchild;  // 再转向右子树的左结点。
      }
      else  
      {
        Pop(&SS,&ee);  // 将结点弹出。
        visit(ee);     // 访问结点。
        pp=ee;         // 记下最近访问过的结点。
        ee=NULL;       // 结点访问完后，重置ee为空，继续读取栈顶元素。
      }
    }
  }
}

// 求二叉树的高度。
int TreeDepth(BiTree TT)
{
  if (TT==NULL) return 0;

  int ll=TreeDepth(TT->lchild);  // 求左子树的高度。

  int rr=TreeDepth(TT->rchild);  // 求右子树的高度。

  return ll>rr ? ll+1: rr+1;  // 取左、右子树的较大者再加上根结点的高度。
}
