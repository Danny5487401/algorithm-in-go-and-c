/*
 * 程序名：btree3.c，此程序演示中序线索二叉树的创建及求前驱后继的方法。
 * 作者：C语言技术网(www.freecplus.net) 日期：20200202
*/
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

///////////////////////////////////////////////
// 线索二叉树的数据结构。
typedef struct TBTNode
{
  char   data;               // 存放结点的数据元素。
  struct TBTNode *lchild;    // 指向左子结点地址的指针。
  struct TBTNode *rchild;    // 指向右子结点地址的指针。
  unsigned char ltag,rtag;   // 左右指针的类型，0-非线索指针，1-线索指针。
}TBTNode,*TBTTree;
///////////////////////////////////////////////

// 访问结点元素。
void visit(TBTNode *pp);

// 采用递归的方法对二叉树的先序遍历。
void PreOrder(TBTTree TT);

// 采用递归的方法对二叉树的中序遍历。
void InOrder(TBTTree TT);

// 采用递归的方法对二叉树的后序遍历。
void PostOrder(TBTTree TT);

// 中序遍历二叉树线索化主函数。
void CrtInThread(TBTTree TT);

// 中序遍历线索化二叉树的递归函数，参数pp是当前节点，pre是前驱结点。
void InThread(TBTTree pp,TBTTree *pre);  // C语言的pre参数要用指针的指针。
// void InThread(TBTTree pp,TBTTree &pre); // C++语言的pre参数为引用，如果启用它，需要用C++的编译器编译。

//////////////////////////////////////////////////
// 获取中序遍历线索二叉树的第一个结点。
TBTNode *FirstNode(TBTTree TT);

// 获取中序遍历线索二叉树的后继结点。
TBTNode *NextNode(TBTNode *pp);

// 中序线索二叉树的遍历函数。
void InOrder1(TBTTree TT);
//////////////////////////////////////////////////

///////////////////////////////////////////////////
// 获取中序遍历线索二叉树的最后一个结点。
TBTNode *LastNode(TBTTree TT);

// 获取中序遍历线索二叉树的前驱结点。
TBTNode *PriorNode(TBTNode *pp);

// 从右到左中序线索二叉树的遍历函数。
void InOrder2(TBTTree TT);
//////////////////////////////////////////////////

int main()
{
  TBTTree TT=0; // 声明树指针变量。

  /*
  // 手工构造一个如下结构的二叉树。
             1
          /     \
         2       3
        / \     /
       4   5   6
      / \ /   /
     7  8 9  0
  */

  // 分配根节点。
  TT=(TBTNode *)malloc(sizeof(TBTNode));
  TT->data='1';
  TT->lchild=TT->rchild=0;
  TT->ltag=TT->rtag=0;

  // 第二层第一个节点。
  TT->lchild=(TBTNode *)malloc(sizeof(TBTNode));
  TT->lchild->data='2'; 
  TT->lchild->lchild=TT->lchild->rchild=0;
  TT->lchild->ltag=TT->lchild->rtag=0;

  // 第二层第二个节点。
  TT->rchild=(TBTNode *)malloc(sizeof(TBTNode));
  TT->rchild->data='3'; 
  TT->rchild->lchild=TT->rchild->rchild=0;
  TT->rchild->ltag=TT->rchild->rtag=0;

  // 第三层第一个节点。
  TT->lchild->lchild=(TBTNode *)malloc(sizeof(TBTNode));
  TT->lchild->lchild->data='4'; 
  TT->lchild->lchild->lchild=TT->lchild->lchild->rchild=0;
  TT->lchild->lchild->ltag=TT->lchild->lchild->rtag=0;

  // 第三层第二个节点。
  TT->lchild->rchild=(TBTNode *)malloc(sizeof(TBTNode));
  TT->lchild->rchild->data='5'; 
  TT->lchild->rchild->lchild=TT->lchild->rchild->rchild=0;
  TT->lchild->rchild->ltag=TT->lchild->rchild->rtag=0;

  // 第三层第三个节点。
  TT->rchild->lchild=(TBTNode *)malloc(sizeof(TBTNode));
  TT->rchild->lchild->data='6'; 
  TT->rchild->lchild->lchild=TT->rchild->lchild->rchild=0;
  TT->rchild->lchild->ltag=TT->rchild->lchild->rtag=0;

  // 第四层第一个节点。
  TT->lchild->lchild->lchild=(TBTNode *)malloc(sizeof(TBTNode));
  TT->lchild->lchild->lchild->data='7'; 
  TT->lchild->lchild->lchild->lchild=TT->lchild->lchild->lchild->rchild=0;
  TT->lchild->lchild->lchild->ltag=TT->lchild->lchild->lchild->rtag=0;

  // 第四层第二个节点。
  TT->lchild->lchild->rchild=(TBTNode *)malloc(sizeof(TBTNode));
  TT->lchild->lchild->rchild->data='8'; 
  TT->lchild->lchild->rchild->lchild=TT->lchild->lchild->rchild->rchild=0;
  TT->lchild->lchild->rchild->ltag=TT->lchild->lchild->rchild->rtag=0;

  // 第四层第三个节点。
  TT->lchild->rchild->lchild=(TBTNode *)malloc(sizeof(TBTNode));
  TT->lchild->rchild->lchild->data='9'; 
  TT->lchild->rchild->lchild->lchild=TT->lchild->rchild->lchild->rchild=0;
  TT->lchild->rchild->lchild->ltag=TT->lchild->rchild->lchild->rtag=0;

  // 第四层第四个节点。
  TT->rchild->lchild->rchild=(TBTNode *)malloc(sizeof(TBTNode));
  TT->rchild->lchild->rchild->data='0'; 
  TT->rchild->lchild->rchild->lchild=TT->rchild->lchild->rchild->rchild=0;
  TT->rchild->lchild->rchild->ltag=TT->rchild->lchild->rchild->rtag=0;

  //////////////////////////////////////////////////////////
  // 二叉树的先序遍历。
  printf("先序遍历结果："); PreOrder(TT); printf("\n");

  // 二叉树的中序遍历。
  printf("中序遍历结果："); InOrder(TT); printf("\n");

  // 二叉树的后序遍历。
  printf("后序遍历结果："); PostOrder(TT); printf("\n");
  // 以上遍历二叉树的代码要放在线索化之前。
  //////////////////////////////////////////////////////////

  // 中序遍历二叉树线索化主函数。
  CrtInThread(TT);

  printf("中序线索二叉树从左到右遍历结果："); InOrder1(TT); printf("\n");
  printf("中序线索二叉树从右到左遍历结果："); InOrder2(TT); printf("\n");

  return 0;
}

// 访问结点元素。
void visit(TBTNode *pp)
{
  printf("%c ",pp->data);  // 访问结点元素就是把值输出来，意思一下就行了。
}

// 二叉树的先序遍历。
void PreOrder(TBTTree TT)
{
  if (TT == NULL) return;

  visit(TT);               // 访问子树TT的根结点。

  PreOrder(TT->lchild);    // 遍历左子树。

  PreOrder(TT->rchild);    // 遍历右子树。
}

// 二叉树的中序遍历。
void InOrder(TBTTree TT)
{
  if (TT == NULL) return;

  InOrder(TT->lchild);     // 遍历左子树。

  visit(TT);               // 访问子树TT的根结点。

  InOrder(TT->rchild);     // 遍历右子树。
}

// 二叉树的后序遍历。
void PostOrder(TBTTree TT)
{
  if (TT == NULL) return;

  PostOrder(TT->lchild);     // 遍历左子树。

  PostOrder(TT->rchild);     // 遍历右子树。

  visit(TT);                 // 访问子树TT的根结点。
}

// 中序遍历线索化二叉树的递归函数，参数pp是当前节点，pre是前驱结点。
void InThread(TBTTree pp,TBTTree *pre)
{
  if (pp==NULL) return;

  // 递归线索化左子树。
  InThread(pp->lchild,pre);  

  // 建立当前结点的前驱线索。
  if (pp->lchild==NULL)
  {
    pp->lchild=(*pre); pp->ltag=1;
  }

  printf("%c ",pp->data);

  // 建立前驱结点的后继线索。
  if ( ((*pre)!=NULL) && ((*pre)->rchild==NULL) )
  {
    (*pre)->rchild=pp; (*pre)->rtag=1;
  }

  (*pre)=pp;  // 改前驱结点为当前结点。

  // 递归线索化右子树。
  InThread(pp->rchild,pre);
}

/*
// C++语言的pre参数为引用，如果启用它，需要用C++的编译器编译。
void InThread(TBTTree pp,TBTTree &pre) 
{
  if (pp==NULL) return;

  // 递归线索化左子树。
  InThread(pp->lchild,pre);  

  // 建立当前结点的前驱线索。
  if (pp->lchild==NULL)
  {
    pp->lchild=pre; pp->ltag=1;
  }

  printf("%c ",pp->data);

  // 建立前驱结点的后继线索。
  if ( (pre!=NULL) && (pre->rchild==NULL) )
  {
    pre->rchild=pp; pre->rtag=1;
  }

  pre=pp;  // 改前驱结点为当前结点。

  // 递归线索化右子树。
  InThread(pp->rchild,pre);
}
*/

// 中序遍历二叉树线索化主函数。
void CrtInThread(TBTTree TT)
{
  if (TT==NULL) return;

  TBTNode *pre=NULL;

  // 中序遍历线索化二叉树的递归函数，参数pp是当前节点，pre是前驱结点。
  InThread(TT,&pre);  // C语言指针的指针。
  // InThread(TT,pre);  // C++引用，如果启用它，需要用C++的编译器编译。

  // 处理最后一个结点的右指针。
  pre->rchild=NULL;
  pre->rtag=1;

  printf(" 线索化完成。\n");
}

////////////////////////////////////////////////////////////////////////////
// 获取中序遍历线索二叉树的第一个结点。
TBTNode *FirstNode(TBTTree TT)
{
  // 向左走到尽头。
  while (TT->ltag==0) TT=TT->lchild;

  return TT;
}

// 获取中序遍历线索二叉树的后继结点。
TBTNode *NextNode(TBTNode *pp)
{
  // 如果当前结点的右指针存放的是线索，直接指返回。
  // 如果当前结点的右指针存放的是结点，返回右子树中序遍历的第一个结点。
  if (pp->rtag==0) return FirstNode(pp->rchild);
  else return pp->rchild;
}

// 从左到右中序线索二叉树的遍历函数。
void InOrder1(TBTTree TT)
{
  // 得到第一个结点。
  TBTNode *pp=FirstNode(TT);

  while (pp!=NULL)
  {
    visit(pp);        // 访问结点。
    pp=NextNode(pp);  // 指针移动到后继结点。
  }
}

////////////////////////////////////////////////////////////////////////////
// 获取中序遍历线索二叉树的最后一个结点。
TBTNode *LastNode(TBTTree TT)
{
  // 向右走到尽头。
  while (TT->rtag==0) TT=TT->rchild;

  return TT;
}

// 获取中序遍历线索二叉树的前驱结点。
TBTNode *PriorNode(TBTNode *pp)
{
  // 如果当前结点的左指针存放的是线索，直接指返回。
  // 如果当前结点的左指针存放的是结点，返回左子树中序遍历的最后个结点。
  if (pp->ltag==0) return LastNode(pp->lchild);
  else return pp->lchild;
}

// 从右到左中序线索二叉树的遍历函数。
void InOrder2(TBTTree TT)
{
  // 得到最后一个结点。
  TBTNode *pp=LastNode(TT);

  while (pp!=NULL)
  {
    visit(pp);        // 访问结点。
    pp=PriorNode(pp);  // 指针移动到前驱结点。
  }
}
////////////////////////////////////////////////////////////////////////////

