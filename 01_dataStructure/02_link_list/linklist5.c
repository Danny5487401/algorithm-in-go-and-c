/*
 * 程序名：linklist5.c，此程序演示带头结点的循环单链表的实现，数据元素是整数。
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

// 销毁链表LL。
void DestroyList1(LinkList LL);

// 清空链表。
void ClearList(LinkList LL);                    

// 在链表LL的第ii个位置插入元素ee，返回值：0-失败；1-成功。
int  InsertList(LinkList LL, unsigned int ii, ElemType *ee);   

// 打印链表中全部的元素。
void PrintList(LinkList LL);                    

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

int main()
{
  LinkList LL=NULL; // 声明链表指针变量。

  LL=InitList1();     // 初始化链表。

  printf("LL=%p\n",LL);

  ElemType ee;      // 创建一个数据元素。

  LNode *tmp1;

  printf("在表中插入元素（1、2、3、4、5、6、7、8、9、10）。\n");
  ee=1;  InsertList(LL, 1, &ee);
  ee=2;  InsertList(LL, 1, &ee);
  ee=3;  InsertList(LL, 1, &ee);
  ee=4;  InsertList(LL, 1, &ee);
  ee=5;  InsertList(LL, 1, &ee);
  ee=6;  InsertList(LL, 1, &ee);
  ee=7;  InsertList(LL, 1, &ee);
  ee=8;  InsertList(LL, 1, &ee);
  ee=9;  InsertList(LL, 1, &ee);
  ee=10; InsertList(LL, 1, &ee);

  printf("length=%d\n",LengthList(LL));

  PrintList(LL);

  printf("在第5个位置插入元素（13）。\n");
  ee=13; InsertList(LL, 5, &ee);  

  PrintList(LL);

  printf("在表头插入元素（11），表尾插入元素（12）。\n");
  ee=11; PushFront(LL, &ee);
  ee=12; PushBack(LL, &ee);

  PrintList(LL);

  printf("删除表中第一个结点。\n");
  PopFront(LL); PrintList(LL);

  printf("删除表中最后一个结点。\n");
  PopBack(LL); PrintList(LL);

  LNode *tmp;

  if ( (tmp=LocateNode(LL,5)) != NULL)
    printf("第5个结点的地址是=%p，ee=%d\n",tmp,tmp->data);

  ee=8;
  if ( (tmp=LocateElem(LL,&ee)) != NULL)
    printf("元素值为8的结点的地址是=%p\n",tmp);
  else
    printf("元素值为8的结点的地址是NULL，没找着。\n");

  printf("在结点%p之后插入66\n",tmp);
  ee=66;
  InsertNextNode(tmp,&ee);  PrintList(LL);

  printf("在结点%p之前插入55\n",tmp);
  ee=55;
  InsertPriorNode(tmp,&ee);  PrintList(LL);

  DestroyList1(LL); LL=NULL;  // 销毁链表，LL置为空。

  printf("LL=%p\n",LL);

  return 0;
}

// 初始化链表LL，返回值：失败返回NULL，成功返回头结点的地址。
LNode *InitList1()
{
  LNode *head = (LNode *)malloc(sizeof(LNode));  // 分配头结点。

  if (head == NULL) return NULL;  // 内存不足，返回失败。

  head->next=head;  // 头结点的next指针指向自己。 // xxx

  return head;
}

// 销毁链表LL。
void DestroyList1(LinkList LL)
{
  LNode *head=LL;  // 保存头结点的地址。

  LL=LL->next;  // 从数据结点开始。

  // 释放链表全部的数据结点。
  LNode *tmp;

  while(LL!=head)      // xxx
  {
    tmp=LL->next;  // tmp保存下一结点的地址。
    free(LL);      // 释放当前结点。
    LL=tmp;        // LL指针移动到下一结点。
  }

  free(head);  // 再释放头结点。// xxx

  return;
}

// 清空链表。
void ClearList(LinkList LL)
{
  // 清空链表LL是指释放链表全部的数据结点，不包括头结点。
  if (LL == NULL) { printf("链表LL不存在。\n"); return; } // 判断链表是否存在。

  LNode *tmp1;
  LNode *tmp2=LL->next;  // 保留头结点，从头结点的下一个结点开始释放。

  while(tmp2!=LL) // xxx
  {
    tmp1=tmp2->next;
    free(tmp2);
    tmp2=tmp1;
  }

  LL->next=LL; // 把头结点的next指针指向自己。 // xxx

  return;
}

// 在链表LL的第ii个位置插入元素ee，返回值：0-失败；1-成功。
int InsertList(LinkList LL, unsigned int ii, ElemType *ee)
{
  if ( (LL == NULL) || (ee == NULL) ) { printf("链表LL或元素ee不存在。\n"); return 0; } // 判断表和元素是否存在。

  // 判断插入位置是否合法
  if (ii < 1) { printf("插入位置（%d）不合法，应该在大于0。\n",ii); return 0; }

  // 要在位序ii插入结点，必须找到ii-1结点。
  LNode *pp;

  if (ii == 1) 
  {
    pp=LL;  // 如果ii==1，不用找了，ii-1结点就是头结点。
  }
  else
  {
    pp=LL->next;  // xxx 指针pp指第一个数据结点，逐步往后移动，如果为LL，表示后面没结点了。
    int kk=1;     // kk指向的是第几个结点，pp每向后移动一次，kk就加1。

    while ( (pp != LL) && (kk < ii-1) )   // xxxx
    {
      pp=pp->next;    kk++;

      // printf("pp=%p,kk=%d\n",pp,kk);
    }

    if ( pp == LL ) { printf("位置（%d）不合法，超过了表长。\n",ii); return 0; }  // xxx
  }

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

  // 要在位序ii删除结点，必须找到ii-1结点。
  LNode *pp;

  if (ii == 1)
  {
    pp=LL;  // 如果ii==1，不用找了，ii-1结点就是头结点。
  }
  else
  {
    pp=LL->next;  // xxx 指针pp指第一个数据结点，逐步往后移动，如果为LL，表示后面没结点了。
    int kk=1;     // kk指向的是第几个结点，pp每向后移动一次，kk就加1。

    while ( (pp != LL) && (kk < ii-1) )   // xxxx
    {
      pp=pp->next;    kk++;

      // printf("pp=%p,kk=%d\n",pp,kk);
    }
  }

  if ( pp->next == LL ) { printf("位置（%d）不合法，超过了表长。\n",ii); return 0; }  // xxx

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
  while (pp->next != LL) pp=pp->next;   // xxx
  
  LNode *tmp = (LNode *)malloc(sizeof(LNode));  // 分配一个结点。
  if (tmp == NULL) return 0;  // 内存不足，返回失败。
  
  // 考虑数据元素为结构体的情况，这里采用了memcpy的方法而不是直接赋值。
  memcpy(&tmp->data,ee,sizeof(ElemType));

  // 处理next指针。
  pp->next=tmp;
  tmp->next=LL;   // xxx

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

  // 必须加上这个判断，空表。
  if ( LL->next == LL) { printf("链表LL为空，没有尾结点。\n"); return 0; } // 判断表是否为空。 // xxx

  // 要删除最后一个结点，必须找到最后一个结点的前一个结点。

  LNode *pp=LL;  // 从第0个结点开始。   // xxx

  // 找到倒数第二个结点（包括头结点）。
  while (pp->next->next != LL) pp=pp->next;
 
  // 释放最后一个结点。
  free(pp->next);
  pp->next=LL;   // xxx

  return 1;
}

// 打印链表中全部的元素。
void PrintList(LinkList LL)
{
  if (LL == NULL) { printf("链表LL不存在。\n"); return; } // 判断链表是否存在。

  LNode *pp=LL->next;  // 从第1个结点开始。

  while (pp != LL)   //  xxxxx
  {
    printf("%-3d", pp->data);  // 如果元素ee为结构体，这行代码要修改。
    pp=pp->next;
  }

  printf("\n");

  /*
  // 以下代码用于显示全部数据结点的地址和元素的值。
  LNode *pp=LL->next;  // 从第1个结点开始。

  printf("%p,%p\n",LL,LL->next);  // 显示头结点的指针值。

  while (pp != LL)
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

  while (pp != LL) { pp=pp->next; length++; } // xxxxx

  return length;

  // 不使用临时变量，如何计算链表（包括头结点）的长度？
  // if (LL==NULL) return 0;
  // return LengthList(LL->next)+1;
}

// 判断链表是否为空，返回值：0-非空或失败，1-空。
int IsEmpty(LinkList LL)
{
  if (LL == NULL) return 0;

  if (LL->next == LL) return 1; // xxxx

  return 0;
}

// 获取链表中第ii个结点，成功返回结点的地址，失败返回空。
// 注意，ii可以取值为0，表示头结点。
LNode *LocateNode(LinkList LL, unsigned int ii)
{
  if ( LL == NULL ) { printf("链表LL不存在。\n"); return NULL; } // 判断表和元素是否存在。

  if (ii == 0) return LL;  // 直接返回头结点。
  
  LNode *pp=LL->next;  // 指针pp指向第一个数据结点，逐步往后移动，如果为空，表示后面没结点了。 // xxx
  int kk=1;      // kk指向的是第几个结点，从数据结点开始，pp每向后移动一次，kk就加1。

  while ( (pp != LL) && (kk < ii) )
  { 
    pp=pp->next; kk++; 
  }

  if ( pp==LL ) { printf("位置（%d）不合法，超过了表长。\n",ii); return NULL; }  // xxx

  return pp;
}

// 查找元素ee在链表LL中的结点地址，如果没找到返回NULL，否则返回结点的地址。
LNode *LocateElem(LinkList LL, ElemType *ee)
{
  LNode *pp=LL->next;  // 从第1个数据结点开始。

  while (pp != LL)  // xxxx
  {
    // 如果数据元素是结构体，以下代码要修改。
    if (pp->data == *ee) return pp;

    pp = pp->next;
  }

  return NULL;
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


