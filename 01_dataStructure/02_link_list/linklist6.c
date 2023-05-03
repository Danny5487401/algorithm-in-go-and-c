/*
 * 程序名：linklist6.c，此程序演示链表的常用考题，在linklist1.c上增加了常用考题的解答。
 * 作者：C语言技术网(www.freecplus.net) 日期：20201230
*/
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef int ElemType;     // 自定义链表的数据元素为整数。

typedef struct LNode
{
  ElemType data;       // 存放结点的数据元素。
  struct LNode *next;  // 指向下一个结点的指针。
}LNode,*LinkList;

// 初始化链表LL，返回值：失败返回NULL，成功返回头结点的地址。
LNode *InitList1();

// 传入指针变量的地址的方法。
// 初始化链表，返回值：0-失败；1-成功。
int InitList2(LinkList *LL);

// C++引用的方法，在Linux下，需要用g++编译。
// 初始化链表，返回值：0-失败；1-成功。
// int InitList3(LinkList &LL);

// 如果参数采用转指针LL的值，LL的值只能传进去，无法返回，这种方法是不行的。
int InitList4(LinkList LL);

// 销毁链表LL。
void DestroyList1(LinkList LL);

// 销毁链表LL。
// 传入指针的地址的方法。
void DestroyList2(LinkList *LL);

// C++引用的方法，在Linux下，需要用g++编译。
// 传入指针的地址的方法。
// void DestroyList3(LinkList &LL);

// 清空链表。
void ClearList(LinkList LL);                    

// 在链表LL的第ii个位置插入元素ee，返回值：0-失败；1-成功。
int  InsertList(LinkList LL, unsigned int ii, ElemType *ee);   

// 打印链表中全部的元素。
void PrintList(LinkList LL);                    

// 反向打印链表中全部的元素。
void PrintList1(LNode *pp);

// 在链表LL的头部插入元素ee，返回值：0-失败；1-成功。
int  PushFront(LinkList LL, ElemType *ee);

// 在链表LL的尾部插入元素ee，返回值：0-失败；1-成功。
int  PushBack(LinkList LL, ElemType *ee);

// 删除链表LL中的第ii个结点，返回值：0-位置ii不合法；1-成功。
int  DeleteNode(LinkList LL, unsigned int ii);  

// 删除链表LL中第一个结点，返回值：0-位置不合法；1-成功。
int PopFront(LinkList LL);

// 删除链表LL中最后一个结点，返回值：0-位置不合法；1-成功。
int PopBack(LinkList LL);

// 求链表的长度，返回值：>=0-表LL结点的个数。
int  LengthList(LinkList LL);                   

// 判断链表是否为空，返回值：0-非空或失败，1-空。
int IsEmpty(LinkList LL);

// 获取链表中第ii个结点，成功返回结点的地址，失败返回空。
// 注意，ii可以取值为0，表示头结点。
LNode *LocateNode(LinkList LL, unsigned int ii);

// 查找元素ee在链表LL中的结点地址，如果没找到返回NULL，否则返回结点的地址。
LNode *LocateElem(LinkList LL, ElemType *ee);

// 在指定结点pp之后插入元素ee，返回值：0-失败；1-成功。
int InsertNextNode(LNode *pp, ElemType *ee);

// 在指定结点pp之前插入元素ee，返回值：0-失败；1-成功。
int InsertPriorNode(LNode *pp, ElemType *ee);

// 删除指定结点。
int DeleteNode1(LNode *pp);

// 采用归并的方法，将两个升序的链表La和Lb，合并成一个升序的链表Lc。
int MergeList(LinkList La,LinkList Lb,LinkList Lc);

// 把链表pp结点之后的结点原地逆置（反转），返回值：0-失败；1-成功。
void ReverseList(LNode *pp);

// 找出带头结点的单链表倒数第k(k>0)个结点。
LNode *FindKNode(LinkList LL,unsigned int kk);

// 判断单链表是否有环，并找到环的入口。
LNode *FindLoop(LinkList LL);

// 找出两个汇聚单链表的公共结点，如果没有汇聚，返回空，如果有，返回汇聚的第一个结点。
LNode *FindJoin(LinkList La,LinkList Lb);

// 假设线性表L=(a1,a2,a3,...,an-2,an-1,an)采用带头结点的单链表保存，数据结点的个数为偶数。
// 请设计一个空间复杂度为O(1)且时间上尽可能高效的算法，重新排列L中的各结点。
// 最后得到线性表L=(a1,an,a2,an-1,a3,an-2,...)。
void ReplaceList(LinkList LL);

int main()
{
  LinkList LL=NULL; // 声明链表指针变量。

  LL=InitList1();     // 初始化链表。

  printf("LL=%p\n",LL);

  ElemType ee;      // 创建一个数据元素。

  printf("在表中插入元素（1、2、3、4、5、6、7、8、9、10）。\n");
  ee=1;  PushBack(LL, &ee);
  ee=2;  PushBack(LL, &ee);
  ee=3;  PushBack(LL, &ee);
  ee=4;  PushBack(LL, &ee);
  ee=5;  PushBack(LL, &ee);
  ee=6;  PushBack(LL, &ee);
  ee=7;  PushBack(LL, &ee);
  ee=8;  PushBack(LL, &ee);
  ee=9;  PushBack(LL, &ee);
  ee=10; PushBack(LL, &ee);

  // 按正常顺序显示全部结点的值。
  PrintList(LL);

  // 第1题。
  /*
  // 反向显示全部结点的值。
  PrintList1(LL->next); printf("\n");
  */

  // 第2题。
  /*
  // 找出带头结点的单链表倒数第k(k>0)个结点。
  LNode *tmp=FindKNode(LL,6);
  if (tmp == NULL) 
    printf("倒数第6个结点没找着。\n");
  else
    printf("倒数第6个结点的值是%d。\n",tmp->data);
  */

  // 第3题。
  /*
  // 判断单链表是否有环，并找到环的入口。
  // 手工把链表设置成有环。
  LNode *tmp1=LocateNode(LL,4);
  LNode *tmp2=LocateNode(LL,10);
  tmp2->next=tmp1;
  printf("设置链表环入口的地址是%p，结点的值是%d。\n",tmp1,tmp1->data);

  // 判断单链表是否有环，并找到环的入口。
  LNode *tmp=FindLoop(LL);
  if (tmp == NULL) 
    printf("链表没环。\n");
  else
    printf("链表环入口的地址是%p，结点的值是%d。\n",tmp,tmp->data);
  */

  // 第4题。
  /*
  LinkList La=LL;   // 直接用LL。

  // 创建链表Lb。
  LinkList Lb=InitList1();     // 初始化链表。
  printf("在表Lb中插入元素（11、12、33）。\n");
  ee=11;  PushBack(Lb, &ee);
  ee=12;  PushBack(Lb, &ee);
  ee=13;  PushBack(Lb, &ee);

  // 把Lb最后一个节点的next指针指向La的第7个结点。
  LNode *tmp1=LocateNode(La,7); 
  LNode *tmp2=LocateNode(Lb,3); 
  tmp2->next=tmp1;

  PrintList(La);
  PrintList(Lb);

  // 找出两个汇聚单链表的公共结点，如果没有汇聚，返回空，如果有，返回汇聚的第一个结点。
  LNode *tmp=FindJoin(La,Lb);
  if (tmp == NULL)
    printf("链表没有汇聚。\n");
  else
    printf("链表汇聚的起始地址是%p，结点的值是%d。\n",tmp,tmp->data);
  */

  // 第5题。
  /*
  // 假设线性表L=(a1,a2,a3,...,an-2,an-1,an)采用带头结点的单链表保存，数据结点的个数为偶数。
  // 请设计一个空间复杂度为O(1)且时间上尽可能高效的算法，重新排列L中的各结点。
  // 最后得到线性表L=(a1,an,a2,an-1,a3,an-2,...)。
  ReplaceList(LL);
  PrintList(LL);
  */

  // 在测试FindLoop函数的时候，不能启用以下代码。
  // DestroyList1(LL); LL=NULL;  // 销毁链表，LL置为空。

  return 0;
}

// 如果参数采用转指针LL的值，LL的值只能传进去，无法返回，这种方法是不行的。
int InitList4(LinkList LL)
{
  LNode *head = (LNode *)malloc(sizeof(LNode));  // 分配头结点。

  if (head == NULL) return 0;  // 内存不足，返回失败。

  head->next=NULL;  // 头结点的下一结点暂时不存在，置空。

  LL=head;

  printf("LL1=%p\n",LL);

  return 1;
}

// 初始化链表LL，返回值：失败返回NULL，成功返回头结点的地址。
LNode *InitList1()
{
  LNode *head = (LNode *)malloc(sizeof(LNode));  // 分配头结点。

  if (head == NULL) return NULL;  // 内存不足，返回失败。

  head->next=NULL;  // 头结点的下一结点暂时不存在，置空。

  return head;
}

// 传入指针变量的地址的方法。
// 初始化链表，返回值：0-失败；1-成功。
int InitList2(LinkList *LL)
{
  // 在本函数中，LL是指针的指针，用于存放指针的地址。

  if ( *LL != NULL ) { printf("链表LL已存在，在初始化之前请先释放它。\n"); return 0; }

  LNode *head = (LNode *)malloc(sizeof(LNode));  // 分配头结点。

  if (head == NULL) return 0;  // 内存不足，返回失败。

  head->next=NULL;  // 头结点的下一结点暂时不存在，置空。

  *LL=head;

  return 1;
}

/*
// C++引用的方法。
// 初始化链表，返回值：0-失败；1-成功。
int InitList3(LinkList &LL)
{
  if ( LL != NULL ) { printf("链表L已存在，在初始化之前请先释放它。\n"); return 0; }

  LNode *head = (LNode *)malloc(sizeof(LNode));  // 分配头结点。

  if (head == NULL) return 0;  // 内存不足，返回失败。

  head->next=NULL;  // 头结点的下一结点暂时不存在，置空。

  LL=head;

  return 1;
}
*/

// 销毁链表LL。
void DestroyList1(LinkList LL)
{
  // 销毁链表LL是指释放链表全部的结点，包括头结点。
  LNode *tmp;

  while(LL!=NULL)
  {
    tmp=LL->next;  // tmp保存下一结点的地址。
    free(LL);      // 释放当前结点。
    LL=tmp;        // LL指针移动到下一结点。
  }

  // LL=NULL;   // LL在本函数中相当于局部变量，就算置空了也不会影响调用者传递的LL，所以LL=NULL没有意义。

  return;
}

// 销毁链表LL。
void DestroyList2(LinkList *LL)
{  
  // 如果函数的参数是指针的指针，可以启用以下代码。
  LNode *tmp1,*tmp2;

  tmp1=*LL;

  while(tmp1!=NULL)
  {
    tmp2=tmp1->next; // tmp保存下一结点的地址。
    free(tmp1);      // 释放当前结点。
    tmp1=tmp2;       // LL指针移动到下一结点。
  }

  *LL=NULL;  // 把链表的指针置为空，表示链表不存在了。

  return;
}

/*
// C++引用的方法。
// 传入指针的地址的方法。
void DestroyList3(LinkList &LL)
{
  // 销毁链表LL是指释放链表全部的结点，包括头结点。
  LNode *tmp;

  while(LL!=NULL)
  {
    tmp=LL->next;   // tmp保存下一结点的地址。
    free(LL);       // 释放当前结点。
    LL=tmp;         // LL指针移动到下一结点。
  }

  LL=NULL;  // 把链表的指针置为空，表示链表不存在了。

  return;
}
*/

// 清空链表。
void ClearList(LinkList LL)
{
  // 清空链表LL是指释放链表全部的结点，但不包括头结点。
  if (LL == NULL) { printf("链表LL不存在。\n"); return; } // 判断链表是否存在。

  LNode *tmp1;
  LNode *tmp2=LL->next;  // 保留头结点，从头结点的下一个结点开始释放。

  while(tmp2!=NULL)
  {
    tmp1=tmp2->next;
    free(tmp2);
    tmp2=tmp1;
  }

  LL->next=NULL; // 这行代码一定不能少，否则会留下野指针。

  return;
}

// 在链表LL的第ii个位置插入元素ee，返回值：0-失败；1-成功。
int InsertList(LinkList LL, unsigned int ii, ElemType *ee)
{
  if ( (LL == NULL) || (ee == NULL) ) { printf("链表LL或元素ee不存在。\n"); return 0; } // 判断表和元素是否存在。

  // 判断插入位置是否合法
  if (ii < 1) { printf("插入位置（%d）不合法，应该在大于0。\n",ii); return 0; }

  // 要在位序ii插入结点，必须找到ii-1结点。
  LNode *pp=LL;  // 指针pp指向头结点，逐步往后移动，如果为空，表示后面没结点了。
  int kk=0;      // kk指向的是第几个结点，从头结点0开始，pp每向后移动一次，kk就加1。

  while ( (pp != NULL) && (kk < ii-1) )
  {
    pp=pp->next; kk++;

    // printf("pp=%p,kk=%d\n",pp,kk);
  }

  if ( pp==NULL ) { printf("位置（%d）不合法，超过了表长。\n",ii); return 0; }

  LNode *tmp = (LNode *)malloc(sizeof(LNode));  // 分配一个结点。
  if (tmp == NULL) return 0;  // 内存不足，返回失败。
  
  // 考虑数据元素为结构体的情况，这里采用了memcpy的方法而不是直接赋值。
  memcpy(&tmp->data,ee,sizeof(ElemType));

  // 处理next指针。
  tmp->next=pp->next;
  pp->next=tmp;

  return 1;

  ///////////////////////////////////////
  // 以上代码可以用以下代码代替。
  // LNode *pp=LocateNode(LL,ii-1);  
  // return InsertNextNode(pp,ee);
  ///////////////////////////////////////
}

// 删除链表LL中的第ii个结点，返回值：0-位置ii不合法；1-成功。
int  DeleteNode(LinkList LL, unsigned int ii)  
{
  if (LL == NULL) { printf("链表L不存在。\n"); return 0; } // 判断链表是否存在。

  // 判断删除位置是否合法
  if (ii < 1) { printf("删除位置（%d）不合法，应该在大于0。\n",ii); return 0; }

  // 要删除位序ii结点，必须找到ii-1结点。
  LNode *pp=LL;  // 指针pp指向头结点，逐步往后移动，如果为空，表示后面没结点了。
  int kk=0;      // kk指向的是第几个结点，从头结点0开始，pp每向后移动一次，kk就加1。

  while ( (pp != NULL) && (kk < ii-1) )
  {
    pp=pp->next; kk++;
  }

  // 注意，以下行的代码与视频中的不一样，视频中的是 if ( pp==NULL )，有bug。
  if ( pp->next==NULL ) { printf("位置（%d）不合法，超过了表长。\n",ii); return 0; }

  LNode *tmp=pp->next;  // tmp为将要删除的结点。
  pp->next=pp->next->next;   // 写成p->next=tmp->next更简洁。

  free(tmp);

  return 1;
}

// 在链表LL的头部插入元素ee，返回值：0-失败；1-成功。
int  PushFront(LinkList LL, ElemType *ee)
{
  return InsertList(LL,1,ee);
}

// 在链表LL的尾部插入元素ee，返回值：0-失败；1-成功。
int PushBack(LinkList LL, ElemType *ee)
{
  if ( (LL == NULL) || (ee == NULL) ) { printf("链表LL或元素ee不存在。\n"); return 0; } // 判断表和元素是否存在。

  LNode *pp=LL;  // 从头结点开始。

  // 找到最后一个结点。
  while (pp->next != NULL) pp=pp->next;
  
  LNode *tmp = (LNode *)malloc(sizeof(LNode));  // 分配一个结点。
  if (tmp == NULL) return 0;  // 内存不足，返回失败。
  
  // 考虑数据元素为结构体的情况，这里采用了memcpy的方法而不是直接赋值。
  memcpy(&tmp->data,ee,sizeof(ElemType));

  // 处理next指针。
  tmp->next=NULL;
  pp->next=tmp;

  return 1;
}

// 删除链表LL中第一个结点，返回值：0-位置不合法；1-成功。
int PopFront(LinkList LL)
{
  return DeleteNode(LL, 1);
}

// 删除链表LL中最后一个结点，返回值：0-位置不合法；1-成功。
int PopBack(LinkList LL)
{
  if ( LL == NULL ) { printf("链表LL不存在。\n"); return 0; } // 判断表和元素是否存在。

  // 必须加上这个判断，否则下面的循环pp->next->next不成立。
  if ( LL->next == NULL) { printf("链表LL为空，没有尾结点。\n"); return 0; } // 判断表是否为空。

  // 要删除最后一个结点，必须找到最后一个结点的前一个结点。

  LNode *pp=LL;  // 从第0个结点开始。

  // 找到倒数第二个结点（包括头结点）。
  while (pp->next->next != NULL) pp=pp->next;
 
  // 释放最后一个结点。
  free(pp->next);
  pp->next=NULL;

  return 1;
}

// 打印链表中全部的元素。
void PrintList(LinkList LL)
{
  if (LL == NULL) { printf("链表LL不存在。\n"); return; } // 判断链表是否存在。

  LNode *pp=LL->next;  // 从第1个结点开始。

  while (pp != NULL)
  {
    printf("%-3d", pp->data);  // 如果元素ee为结构体，这行代码要修改。
    pp=pp->next;
  }

  printf("\n");

  /*
  // 以下代码用于显示全部结点的地址和元素的值。
  LNode *pp=LL;  // 从第0个结点开始。

  while (pp != NULL)
  {
    printf("%p,%p,%-3d\n",pp,pp->next,pp->data);  // 如果元素ee为结构体，这行代码要修改。
    pp=pp->next;
  }
  */
}

// 求链表的长度，返回值：>=0-表LL结点的个数。
int  LengthList(LinkList LL)
{
  if (LL == NULL) { printf("链表LL不存在。\n"); return 0; } // 判断链表是否存在。

  LNode *pp=LL->next;  // 头结点不算，从第1个结点开始。

  int length=0;

  while (pp != NULL) { pp=pp->next; length++; }

  return length;

  // 不使用临时变量，如何计算链表（包括头结点）的长度？
  // if (LL==NULL) return 0;
  // return LengthList(LL->next)+1;
}

// 判断链表是否为空，返回值：0-非空，1-空。
int IsEmpty(LinkList LL)
{
  if (LL == NULL) return 0;

  if (LL->next == NULL) return 1;

  return 0;
}

// 获取链表中第ii个结点，成功返回结点的地址，失败返回空。
// 注意，ii可以取值为0，表示头结点。
LNode *LocateNode(LinkList LL, unsigned int ii)
{
  if ( LL == NULL ) { printf("链表LL不存在。\n"); return NULL; } // 判断表和元素是否存在。
  
  LNode *pp=LL;  // 指针pp指向头结点，逐步往后移动，如果为空，表示后面没结点了。
  int kk=0;      // kk指向的是第几个结点，从头结点0开始，pp每向后移动一次，kk就加1。

  while ( (pp != NULL) && (kk < ii) )
  { 
    pp=pp->next; kk++; 
  }

  if ( pp==NULL ) { printf("位置（%d）不合法，超过了表长。\n",ii); return NULL; }

  return pp;
}

// 查找元素ee在链表LL中的结点地址，如果没找到返回NULL，否则返回结点的地址。
LNode *LocateElem(LinkList LL, ElemType *ee)
{
  LNode *pp=LL->next;  // 从第1个数据结点开始。

  while (pp != NULL)
  {
    // 如果数据元素是结构体，以下代码要修改。
    if (pp->data == *ee) return pp;

    pp = pp->next;
  }

  return NULL;
}

// 采用归并的方法，将两个升序的链表La和Lb，合并成一个升序的链表Lc。
int MergeList(LinkList La,LinkList Lb,LinkList Lc)
{
  if ( (La == NULL) || (Lb == NULL) || (Lc == NULL) ) { printf("表La、Lb、Lc至少有一个不存在。\n"); return 0; }

  La=La->next;
  Lb=Lb->next;

  LNode *pp;

  // 把La和Lb合并到Lc中。
  while ( (La != NULL) && (Lb != NULL) )
  {
    // 取La和Lb的较小者。
    if (La->data <= Lb->data)
    {
      pp=La; La=La->next;
    }
    else
    {
      pp=Lb; Lb=Lb->next;
    }

    // 把较小者追加到Lc中。
    Lc->next=(LNode *)malloc(sizeof(LNode));  // 分配一个新结点。
    Lc=Lc->next;
    memcpy(&Lc->data,&pp->data,sizeof(ElemType));
    Lc->next=NULL;
  }

  // 把链表La其它的元素追加到Lc中。
  while (La != NULL)
  {
    Lc->next=(LNode *)malloc(sizeof(LNode));  // 分配一个新结点。
    Lc=Lc->next;
    memcpy(&Lc->data,&La->data,sizeof(ElemType));
    Lc->next=NULL;
    La=La->next;
  }
    
  // 把链表Lb其它的元素追加到Lc中。
  while (Lb != NULL)
  {
    Lc->next=(LNode *)malloc(sizeof(LNode));  // 分配一个新结点。
    Lc=Lc->next;
    memcpy(&Lc->data,&Lb->data,sizeof(ElemType));
    Lc->next=NULL;
    Lb=Lb->next;
  }

  return 1;
}

// 在指定结点pp之后插入元素ee，返回值：0-失败；1-成功。
int InsertNextNode(LNode *pp, ElemType *ee)
{
  if (pp == NULL) { printf("结点pp不存在。\n"); return 0; }

  LNode *tmp = (LNode *)malloc(sizeof(LNode));
  if (tmp == NULL) return 0;

  memcpy(&tmp->data,ee,sizeof(ElemType));
  tmp->next=pp->next;
  pp->next=tmp;

  return 1;
}

// 在指定结点pp之前插入元素ee，返回值：0-失败；1-成功。
int InsertPriorNode(LNode *pp, ElemType *ee)
{
  if (pp == NULL) { printf("结点pp不存在。\n"); return 0; }

  // 在指定结点pp之前插入采用偷梁换柱的方法：
  // 1、分配一个新的结点；
  // 2、把pp结点的数据和指针复制到新结点中。
  // 3、把待插入元素的数据存入pp结点中。

  LNode *tmp = (LNode *)malloc(sizeof(LNode));
  if (tmp == NULL) return 0;

  // 把pp结点的数据和指针复制到tmp中。
  memcpy(&tmp->data,&pp->data,sizeof(ElemType));
  tmp->next=pp->next;

  // 把待插入的元素存入pp中。
  memcpy(&pp->data,ee,sizeof(ElemType));
  pp->next=tmp;
  
  return 1;
}

// 删除指定结点。
int DeleteNode1(LNode *pp)
{
  if (pp == NULL) { printf("结点pp不存在。\n"); return 0; }

  // 删除指定结点的思想是：1）把pp后继结点的数据和next指针复制到pp结点；2）删除pp结点的后继结点。
  LNode *tmp=pp->next;  // tmp指向pp的后继结点。
  memcpy(&pp->data,&tmp->data,sizeof(ElemType)); // 把后继结点的数据复制到pp结点中。
  pp->next=tmp->next;   // 把pp的next指向后继结点的next。
  free(tmp);  // 释放后继结点。

  // 写这个函数的目的是告诉大家这种方法是有问题的。
  // 问题：如果当前的pp结点是链表的最后一个结点，那么它的后继结点根本不存在。
  // 结论：此法不通，还是乖乖的从链表头部开始扫描。

  return 1;
}

// 把链表pp结点之后的结点原地逆置（反转），返回值：0-失败；1-成功。
void ReverseList(LNode *pp)
{
  LNode *ss;      // 当前结点。
  LNode *ssnext;  // 当前结点的下一结点。

  ss=pp->next;    // 从pp结点之后的结点开始反转。

  pp->next=NULL;  // pp->next指向空。

  while (ss != NULL)
  {
    ssnext=ss->next;  // 保留ss下一结点的地址。

    // 以下两行相当于在pp之后插入ss结点。
    ss->next=pp->next;  
    pp->next=ss;

    ss=ssnext;  // ss结点后移。
  }
}

// 反向打印链表中全部的元素。
void PrintList1(LNode *pp)
{
  if (pp == NULL) return;

  PrintList1(pp->next);

  printf("%-3d", pp->data);  

  return;
}

// 找出带头结点的单链表倒数第k(k>0)个结点。
LNode *FindKNode(LinkList LL,unsigned int kk)
{
  LNode *pp1=LL->next;
  LNode *pp2=LL->next;

  int ii=1;

  // 让pp1先跑kk个节点。
  while ( (pp1!=NULL) && (ii<kk) )
  {
    pp1=pp1->next; ii++;
  }
    
  if (pp1 == NULL) return NULL;

  // 然后pp1继续跑到尾，pp2也开始跑。
  while (1)
  {
    pp1=pp1->next; 
    if (pp1 == NULL) break;
    pp2=pp2->next;
  }
  
  return pp2;
}

// 判断单链表是否有环，并找到环的入口。
LNode *FindLoop(LinkList LL)
{
  LNode *fast=LL,*slow=LL;  // 设置快慢指针。

  while ( (fast != NULL) && (fast->next != NULL) ) // 防止链表没环操作空指针的情况。
  {
    // 慢指针走一步，快指针走两步。
    slow=slow->next;
    fast=fast->next->next;

    if (slow == fast) break;
  }

  // 如果快指针能跑到头，表示没环。
  if ( (fast==NULL) || (fast->next==NULL) ) return NULL;

  // 找到环的入口。
  LNode *p1=LL,*p2=slow;
  while (p1 != p2)
  {
    p1=p1->next; p2=p2->next;
  }

  return p1;
}

// 找出两个汇聚单链表的公共结点，如果没有汇聚，返回空，如果有，返回汇聚的第一个结点。
LNode *FindJoin(LinkList La,LinkList Lb)
{
  int lalen=LengthList(La); // 得到La的长度。
  int lblen=LengthList(Lb); // 得到Lb的长度。

  LNode *pa,*pb;
  
  // 如果La更长，La先走lalen-lblen步。
  for (pa=La;lblen<lalen;lalen--)
    pa=pa->next;

  // 如果Lb更长，Lb先走lblen-lalen步。
  for (pb=Lb;lalen<lblen;lblen--)
    pb=pb->next;

  // 然后La和Lb一起走，结点地址相同就是汇聚点。
  while (1)
  {
    if (pa == pb) break;

    pa=pa->next; pb=pb->next;

    if (pa == NULL) return NULL;   // 跑到头了还没有相同的地址，表示没有汇聚。
  }
  
  return pa;  // 返回汇聚的第一个结点的地址。
}

// 假设线性表L=(a1,a2,a3,...,an-2,an-1,an)采用带头结点的单链表保存，数据结点的个数为偶数。
// 请设计一个空间复杂度为O(1)且时间上尽可能高效的算法，重新排列L中的各结点。
// 最后得到线性表L=(a1,an,a2,an-1,a3,an-2,...)。
void ReplaceList(LinkList LL)
{
  // 第1步：找到链表的中间结点。
  LNode *fast=LL,*slow=LL;  // 设置快慢指针。

  while (fast->next != NULL) 
  {
    // 慢指针走一步，快指针走两步。
    slow=slow->next;
    fast=fast->next->next;  // 因为数据结点的个数为偶数，所以这里不会出现操作空指针的情况。
    // printf("slow=%p,data=%d\n",slow,slow->data);
    // printf("fast=%p,data=%d\n",fast,fast->data);
  }

  // slow指向链表前半段最后一个结点。
  // printf("slow=%p,data=%d\n",slow,slow->data);

  // 第2步：把链表的后半段原地逆置。
  ReverseList(slow);

  // PrintList(LL);

  // 第3步：把链表的后半段插入到前半段中。
  LNode *n1=LL->next;    // n1为前半段第一个数据结点。
  LNode *n2=slow->next;  // n2为后半段第一个数据结点。

  slow->next=NULL;  // 链表结束标志。
 
  LNode *tmp;

  while (n2 != NULL)
  {
    tmp=n2->next;      // 保存下一结点的地址。
    n2->next=n1->next;
    n1->next=n2;
    n2=tmp;
    n1=n1->next->next; // n1跳到下一个插入位置。
  }
}

