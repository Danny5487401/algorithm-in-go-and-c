/*
 * 程序名：linklist2.c，此程序演示带头结点的单链表的实现，数据元素是结构体。
 * 作者：C语言技术网(www.freecplus.net) 日期：20201230
*/
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef struct
{
  int  no;        // 超女编号。
  char name[31];  // 超女姓名。
}ElemType;        // 自定义顺序表的数据元素为结构体。

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

// 打印链表中全部的元素，包括超女编号和姓名。
void PrintList1(LinkList LL);

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

int main()
{
  LinkList LL=NULL; // 声明链表指针变量。

  LL=InitList1();     // 初始化链表。

  // 如果要在函数中对变量进行赋值，必须把变量的地址传入函数。
  // 指针变量简称指针，如果要在函数中对指针变量赋值，也必须把指针的地址传入函数。
  // LL是指针，在InitList2函数中，需要把头结点的地址赋值给LL，所以要传入LL的地址。
  // 指针是变量，用于存放变量的地址，指针不是地址，指针里存放的内容才是地址。
  // 所以，这里要把指针变量LL的地址传给InitList2()函数。
  // 各位菜鸡，明白鸟吗？
  // InitList2(&LL);     // 初始化链表，传入指针变量LL的地址。

  // InitList3(LL);      // 初始化链表，C++的引用。

  // 如果参数采用转指针LL的值，LL的值只能传进去，无法返回，这种方法是不行的。
  // InitList4(LL);

  printf("LL=%p\n",LL);

  ElemType ee;      // 创建一个数据元素。

  printf("在表中插入元素（1、2、3、4、5、6、7、8、9、10）。\n");
  ee.no=1;  strcpy(ee.name,"妲己");   InsertList(LL, 1, &ee);
  ee.no=2;  strcpy(ee.name,"褒姒");   InsertList(LL, 1, &ee);
  ee.no=3;  strcpy(ee.name,"西施");   InsertList(LL, 1, &ee);
  ee.no=4;  strcpy(ee.name,"王昭君"); InsertList(LL, 1, &ee);
  ee.no=5;  strcpy(ee.name,"貂蝉");   InsertList(LL, 1, &ee);
  ee.no=6;  strcpy(ee.name,"杨玉环"); InsertList(LL, 1, &ee);
  ee.no=7;  strcpy(ee.name,"赵飞燕"); InsertList(LL, 1, &ee);
  ee.no=8;  strcpy(ee.name,"萧美娘"); InsertList(LL, 1, &ee);
  ee.no=9;  strcpy(ee.name,"陈圆圆"); InsertList(LL, 1, &ee);
  ee.no=10; strcpy(ee.name,"李师师"); InsertList(LL, 1, &ee);

  printf("length=%d\n",LengthList(LL));

  PrintList(LL);

  printf("在第5个位置插入元素（13）。\n");
  memset(&ee,0,sizeof(ElemType)); ee.no=13; InsertList(LL, 5, &ee);  

  PrintList(LL);

  printf("在表头插入元素（11），表尾插入元素（12）。\n");
  memset(&ee,0,sizeof(ElemType)); ee.no=11; PushFront(LL, &ee);
  memset(&ee,0,sizeof(ElemType)); ee.no=12; PushBack(LL, &ee);

  PrintList(LL);

  printf("删除表中第7个结点。\n");
  DeleteNode(LL,7); PrintList(LL);

  printf("删除表中第一个结点。\n");
  PopFront(LL); PrintList(LL);

  printf("删除表中最后一个结点。\n");
  PopBack(LL); PrintList(LL);

  LNode *tmp;

  if ( (tmp=LocateNode(LL,3)) != NULL)
    printf("第3个结点的地址是=%p，ee.no=%d\n",tmp,tmp->data.no);

  memset(&ee,0,sizeof(ElemType)); ee.no=8;
  if ( (tmp=LocateElem(LL,&ee)) != NULL)
    printf("元素值为8的结点的地址是=%p\n",tmp);
  else
    printf("元素值为8的结点的地址是NULL，没找着。\n");

  printf("在结点%p之后插入66\n",tmp);
  memset(&ee,0,sizeof(ElemType)); ee.no=66;
  InsertNextNode(tmp,&ee);  PrintList(LL);

  printf("在结点%p之前插入55\n",tmp);
  memset(&ee,0,sizeof(ElemType)); ee.no=55;
  InsertPriorNode(tmp,&ee);  PrintList(LL);

  PrintList1(LL);

  DestroyList1(LL); LL=NULL;  // 销毁链表，LL置为空。
  // DestroyList2(&LL);       // 销毁链表，传入指针的地址，LL在函数中会置为空。
  // DestroyList3(LL);       // 销毁链表，C++的引用，LL在函数中会置为空。

  printf("LL=%p\n",LL);

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

  LNode *tmp=pp->next;   // tmp为将要删除的结点。
  pp->next=pp->next->next;    // 写成p->next=tmp->next更简洁。

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
    printf("%-3d", pp->data.no);  
    pp=pp->next;
  }

  printf("\n");

  /*
  // 以下代码用于显示全部结点的地址和元素的值。
  LNode *pp=LL;  // 从第0个结点开始。

  while (pp != NULL)
  {
    printf("%p,%p,%-3d\n",pp,pp->next,pp->data.no);  
    pp=pp->next;
  }
  */
}

// 打印链表中全部的元素，包括超女编号和姓名。
void PrintList1(LinkList LL)
{
  if (LL == NULL) { printf("链表LL不存在。\n"); return; } // 判断链表是否存在。

  LNode *pp=LL->next;  // 从第1个结点开始。

  while (pp != NULL)
  {
    printf("no=%d,name=%s\n",pp->data.no,pp->data.name);  
    pp=pp->next;
  }
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

// 判断链表是否为空，返回值：0-非空或失败，1-空。
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
    if (pp->data.no == ee->no) return pp;

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

