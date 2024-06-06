/*
 * C语言版本。
 * 程序名：btree4.c，此程序演示二叉排序树的各种操作，包括插入、删除、查找。
 * 作者：C语言技术网(www.freecplus.net) 日期：20200202
*/
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

///////////////////////////////////////////////

typedef int ElemType;     // 自定义数据元素为整数。

// 二叉树的数据结构。
typedef struct BSTNode
{
  ElemType   data;           // 存放结点的数据元素。
  struct BSTNode *lchild;    // 指向左子结点地址的指针。
  struct BSTNode *rchild;    // 指向右子结点地址的指针。
}BSTNode,*BSTree;
///////////////////////////////////////////////

// 在树TT中插入关键字为ee的新结点（递归实现），返回值：0-树中已存在关键字为ee的结点；1-成功。
int Insert(BSTree *TT,ElemType *ee);
// 在树TT中插入关键字为ee的新结点（非递归实现），返回值：0-树中已存在关键字为ee的结点；1-成功。
int Insert1(BSTree *TT,ElemType *ee);

// 用数组arr中的序列构建二叉排序树TT。
// 以下两种写法都可以。
//void CreateBST(BSTree *TT,int arr[],int len);
void CreateBST(BSTree *TT,int *arr,int len);

// 在树TT中查找值为ee的结点（递归实现），成功返回结点的地址，失败返回NULL。
BSTNode *Find(BSTree TT,ElemType *ee);
// 在树TT中查找值为ee的结点（非递归实现），成功返回结点的地址，失败返回NULL。
BSTNode *Find1(BSTree TT,ElemType *ee);

// 在树TT中删除值为ee的结点，成功返回0，结点不存在返回1。
int Delete(BSTree *TT,ElemType *ee);

// 求二叉树的高度。
int TreeDepth(BSTree TT);

// 访问结点元素。
void visit(BSTNode *pp);

// 采用递归的方法对二叉树的先序遍历。
void PreOrder(BSTree TT);

// 采用递归的方法对二叉树的中序遍历。
void InOrder(BSTree TT);

// 采用递归的方法对二叉树的后序遍历。
void PostOrder(BSTree TT);

int main()
{
  BSTree TT=0; // 声明树指针变量。

  ElemType arr[]={50,66,60,26,21,30,70,68}; 

  /*
  // 用数组arr中的序列构建二叉排序树TT。
  // 构建的二叉排序树将如下：
              50
           /     \
          26      66
         /  \    /  \
        21  30 60   70
                   /
                  68
  */
  CreateBST(&TT,arr,sizeof(arr)/sizeof(ElemType));

  ElemType ee;

  // 在树TT中查找值为ee的结点，成功返回结点的地址，失败返回NULL。
  ee=30;
  BSTNode *ss=Find(TT,&ee);
  if (ss!=NULL) printf("查找成功，结点的地址是%p，值是%d。\n",ss,ss->data);
  else printf("查找失败。\n");

  // 二叉树的中序遍历。
  printf("中序遍历结果1："); InOrder(TT); printf("\n");

  // 在树TT中删除值为ee的结点，成功返回0，结点不存在返回1。
  ee=50;
  printf("从树中删除值为ee的结点，Delete()的返回值是%d。\n",Delete(&TT,&ee));

  // 二叉树的中序遍历。
  printf("中序遍历结果2："); InOrder(TT); printf("\n");

  return 0;
}

// 二叉树的先序遍历。
void PreOrder(BSTree TT)
{
  if (TT == NULL) return;

  visit(TT);               // 访问子树TT的根结点。
  PreOrder(TT->lchild);    // 遍历左子树。
  PreOrder(TT->rchild);    // 遍历右子树。
}

// 二叉树的中序遍历。
void InOrder(BSTree TT)
{
  if (TT == NULL) return;

  InOrder(TT->lchild);     // 遍历左子树。
  visit(TT);               // 访问子树TT的根结点。
  InOrder(TT->rchild);     // 遍历右子树。
}

// 二叉树的后序遍历。
void PostOrder(BSTree TT)
{
  if (TT == NULL) return;

  PostOrder(TT->lchild);     // 遍历左子树。
  PostOrder(TT->rchild);     // 遍历右子树。
  visit(TT);                 // 访问子树TT的根结点。
}

// 求二叉树的高度。
int TreeDepth(BSTree TT)
{
  if (TT==NULL) return 0;

  int ll=TreeDepth(TT->lchild);  // 求左子树的高度。

  int rr=TreeDepth(TT->rchild);  // 求右子树的高度。

  return ll>rr ? ll+1: rr+1;     // 取左、右子树的较大者再加上根结点的高度。
}

// 访问结点元素。
void visit(BSTNode *pp)
{
  printf("%d ",pp->data);  // 访问结点元素就是把值输出来，意思一下就行了。
}

// 在树TT中插入关键字为ee的新结点（递归实现），返回值：0-树中已存在关键字为ee的结点；1-成功。
int Insert(BSTree *TT,ElemType *ee)
{
  // 如果当前子树为空，创建子树的根结点。
  if ((*TT)==NULL)
  {
    (*TT)=(BSTree)malloc(sizeof(BSTNode));
    memcpy(&(*TT)->data,ee,sizeof(ElemType));  // 考虑数据元素ee为结构体的情况，采用memcpy而不是直接赋值。
    (*TT)->lchild=(*TT)->rchild=NULL;  // 当前子树的左右孩子指针置空。
    return 1;  // 返回成功。
  }

  if (*ee==(*TT)->data) return 0; // 如果ee已存在，返回失败。

  if (*ee<(*TT)->data) Insert(&(*TT)->lchild,ee);  // 递归向左插入。
  else Insert(&(*TT)->rchild,ee);  // 递归向右插入。
}

// 在树TT中插入关键字为ee的新结点（非递归实现），返回值：0-树中已存在关键字为ee的结点；1-成功。
int Insert1(BSTree *TT,ElemType *ee)
{
  // 如果树为空，创建树的根结点。
  if ((*TT)==NULL)
  {
    (*TT)=(BSTree)malloc(sizeof(BSTNode));
    memcpy(&(*TT)->data,ee,sizeof(ElemType));  // 考虑数据元素ee为结构体的情况，采用memcpy而不是直接赋值。
    (*TT)->lchild=(*TT)->rchild=NULL;  // 树的左右孩子指针置空。
    return 1;  // 返回成功。
  }

  BSTNode *pss=NULL;  // 记录双亲结点的地址。
  BSTNode *ss=(*TT);  // 用于查找的临时指针。
  int  rl=0;          // 记录ss是双亲结点的左子树还是右子树，0-左子树；1-右子树。

  while (ss!=NULL)
  {
    if (*ee==ss->data) return 0; // 如果ee已存在，返回失败。

    pss=ss;  // 记录双亲结点的地址。
    if (*ee<ss->data) { ss=ss->lchild; rl=0; }  // 继续向左走。
    else { ss=ss->rchild; rl=1; }  // 继续向右走。
  }

  // 分配新结点。
  ss=(BSTree)malloc(sizeof(BSTNode));
  memcpy(&ss->data,ee,sizeof(ElemType));  // 考虑数据元素ee为结构体的情况，采用memcpy而不是直接赋值。
  ss->lchild=ss->rchild=NULL;  // 当前子树的左右孩子指针置空。

  // 让双亲结点的左或右指针指向新分配的结点。
  if (rl==0) pss->lchild=ss;
  else pss->rchild=ss;

  return 1;  // 返回成功。
}

// 用数组arr中的序列构建二叉排序树TT。
// 以下两种写法都可以。
//void CreateBST(BSTree *TT,int arr[],int len)
void CreateBST(BSTree *TT,int *arr,int len)
{
  (*TT)=NULL;

  int ii=0;
  for (ii=0;ii<len;ii++)
  {
    Insert(TT,&arr[ii]);  // 把数组中的元素逐个插入树中。
    // Insert1(TT,&arr[ii]);  // 把数组中的元素逐个插入树中。
  }
}

// 在树TT中查找值为ee的结点（递归实现），成功返回结点的地址，失败返回NULL。
BSTNode *Find(BSTree TT,ElemType *ee)  
{
  if (TT==NULL) return NULL;    // 查找失败。

  if (*ee==TT->data) return TT;  // 查找成功。 

  if (*ee<TT->data) return Find(TT->lchild,ee); // 递归向左查找。
  else return Find(TT->rchild,ee);             // 递归向右查找。
}

// 在树TT中查找值为ee的结点（非递归实现），成功返回结点的地址，失败返回NULL。
BSTNode *Find1(BSTree TT,ElemType *ee)  
{
  BSTNode *ss=TT;  // 用于查找的临时指针，也可以直接用TT。

  while (ss!=NULL) 
  {
    if (*ee==ss->data) return ss;  // 成功找到。

    if (*ee<ss->data) ss=ss->lchild;  // 继续向左走。
    else ss=ss->rchild;  // 继续向右走。
  }

  return NULL;

  /*
  // 以下代码更精简，可以取代以上的代码。
  while ( (TT!=NULL) && (*ee!=TT->data) )
  {
    if (*ee<TT->data) TT=TT->lchild;  // 继续向左走。
    else TT=TT->rchild;  // 继续向右走。
  }

  return TT;
  */
}

// 在树TT中删除值为ee的结点，成功返回0，结点不存在返回1。
int Delete(BSTree *TT,ElemType *ee)  
{
  if ((*TT)==NULL) return 1; // 树为空，返回1。

  // 1）如果树只有根结点，并且待删除的结点就是根结点。
  if ( ((*TT)->data==*ee) && ((*TT)->lchild==NULL) && ((*TT)->rchild==NULL) )
  {
    free(*TT); (*TT)=0; return 0;
  }

  BSTNode *pss=NULL;  // 记录双亲结点的地址。
  BSTNode *ss=(*TT);  // 用于查找的临时指针。
  int  rl=0;          // 记录ss是双亲结点的左子树还是右子树，0-左子树；1-右子树。

  // 查找待删除的结点。
  while (ss!=NULL) 
  {
    if (*ee==ss->data) break;  // 成功找到。

    pss=ss;  // 记录双亲结点的地址。
    if (*ee<ss->data) { ss=ss->lchild; rl=0; }  // 继续向左走。
    else { ss=ss->rchild; rl=1; }  // 继续向右走。
  }

  // 结点不存在，返回1。
  if (ss==NULL) return 1;

  // 2）如果待删除的结点ss是叶结点，直接删除，不会破坏二叉排序树的性质。
  if ( (ss->lchild==NULL) && (ss->rchild==NULL) )
  {
    free(ss);   // 释放结点。

    // 修改双亲结点pss的左或右指针指向NULL。
    if (rl==0) pss->lchild=NULL;
    else pss->rchild=NULL;

    return 0;  // 返回成功。
  }
  
  // 3）如果待删除的结点ss只有左子树或右子树，则让子树代替自己。
  if ( (ss->lchild==NULL) || (ss->rchild==NULL) )
  {
    if (ss->lchild!=NULL)  // 有左子树。
    {
      // 修改双亲结点pss的左或右指针指向ss的左子树。
      if (rl==0) pss->lchild=ss->lchild;
      else pss->rchild=ss->lchild;
    }
    else  // 有右子树。
    {
      // 修改双亲结点pss的左或右指针指向ss的右子树。
      if (rl==0) pss->lchild=ss->rchild;
      else pss->rchild=ss->rchild;
    }

    return 0;  // 返回成功。
  }

  // 4）如果待删除的结点ss有左子树和右子树，让左子树最右侧的结点代替自己，然后再删除左子树最右侧的结点。
  // 也可以让右子树最左侧的结点代替自己，然后删除右子树最左侧的结点。
  BSTNode *pss1=ss;        // 记录双亲结点的地址。
  BSTNode *ss1=ss->lchild; // 用于查找的临时指针。
  int  rl1=0;              // 记录ss1是双亲结点的左子树还是右子树，0-左子树；1-右子树。

  // 左子树向右走到尽头。
  while (ss1->rchild!=NULL)
  {
    rl1=1; pss1=ss1; ss1=ss1->rchild;
  }
  
  // 把左子树最右侧的结点值复制到结点ss中。
  memcpy(&ss->data,&ss1->data,sizeof(ElemType));

  // 左子树最右侧的结点ss1必定无右子树。
  // 修改双亲结点pss1的左或右指针指向ss1的左子树，ss1的左子树可以为空。
  if (rl1==0) pss1->lchild=ss1->lchild;
  else pss1->rchild=ss1->lchild;

  free(ss1);  // 释放ss1。

  return 0;
}
