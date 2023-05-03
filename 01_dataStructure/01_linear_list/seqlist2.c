/*
 * 程序名：seqlist2.c，此程序演示顺序表的静态实现，数据元素是结构体。
 * 作者：C语言技术网(www.freecplus.net) 日期：20201230
*/
#include <stdio.h>
#include <string.h>

#define MAXSIZE 100       // 顺序表的最大长度

typedef struct
{
  int  no;        // 超女编号。
  char name[31];  // 超女姓名。
}ElemType;        // 自定义顺序表的数据元素为结构体。

typedef struct
{
  ElemType data[MAXSIZE];   // 用数组存储顺序表中的元素
  unsigned int length;      // 顺序表中元素的个数
}SeqList,*PSeqList;

// 顺序表LL的初始化操作。
void InitList(PSeqList LL);                     

// 销毁顺序表LL。
void DestroyList(PSeqList LL);

// 在顺序表LL的第ii个位置插入元素ee，返回值：0-失败；1-成功。
int  InsertList(PSeqList LL, unsigned int ii, ElemType *ee);   

// 在顺序表LL的头结点插入元素ee，返回值：0-失败；1-成功。
int  PushFront(PSeqList LL, ElemType *ee);

// 在顺序表LL的尾结点插入元素ee，返回值：0-失败；1-成功。
int  PushBack(PSeqList LL, ElemType *ee);

// 求顺序表的长度，返回值：>=0-表LL元素的个数。
int  LengthList(PSeqList LL);                   

// 获取顺序表中第ii个元素的值，存放在ee中，返回值：0-失败；1-成功。
int  GetElem(PSeqList LL, unsigned int ii, ElemType* ee);  

// 清空顺序表。
void ClearList(PSeqList LL);                    

// 判断顺序表是否为空，返回值：1-空，0-非空或失败。
int  IsEmpty(PSeqList LL);                    

// 打印顺序表中全部的元素。
void PrintList(PSeqList LL);                    

// 查找ee在顺序表LL中的位置，返回值：0-元素ee在表LL中不存在，>0元素ee在表LL中的位置。
int  LocateElem(PSeqList LL, ElemType *ee);

// 删除顺序表LL中的第ii个元素，返回值：0-位置ii不合法；1-成功。
int  DeleteElem(PSeqList LL, unsigned int ii);  

// 删除顺序表LL中头元素，返回值：0-位置ii不合法；1-成功。
int PopFront(PSeqList LL);

// 删除顺序表LL中尾元素，返回值：0-位置ii不合法；1-成功。
int PopBack(PSeqList LL);

int main()
{
  SeqList LL;     // 创建顺序表。

  ElemType ee;    // 创建一个数据元素。

  InitList(&LL);  // 初始化顺序表。

  printf("在表中插入元素（1、2、3、4、5、6、7、8、9、10）。\n");
  ee.no=1;  InsertList(&LL, 1, &ee); 
  ee.no=2;  InsertList(&LL, 1, &ee);
  ee.no=3;  InsertList(&LL, 1, &ee);
  ee.no=4;  InsertList(&LL, 1, &ee);
  ee.no=5;  InsertList(&LL, 1, &ee);
  ee.no=6;  InsertList(&LL, 1, &ee);
  ee.no=7;  InsertList(&LL, 1, &ee);
  ee.no=8;  InsertList(&LL, 1, &ee);
  ee.no=9;  InsertList(&LL, 1, &ee);
  ee.no=10; InsertList(&LL, 1, &ee);

  PrintList(&LL);

  printf("在表头插入元素（11），表尾插入元素（12）。\n");
  ee.no=11; PushFront(&LL, &ee);
  ee.no=12; PushBack(&LL, &ee);

  PrintList(&LL);

  printf("在第5个位置插入元素（13）。\n");
  ee.no=13; InsertList(&LL, 5, &ee);  

  PrintList(&LL);

  printf("删除表中第7个元素。\n");
  DeleteElem(&LL,7); PrintList(&LL);

  printf("删除表头元素。\n");
  PopFront(&LL); PrintList(&LL);

  printf("删除表尾元素。\n");
  PopBack(&LL); PrintList(&LL);

  GetElem(&LL,3,&ee);
  printf("第3个元素的值是%d。\n",ee.no);

  ee.no=8;
  printf("元素值为8的位置是=%d\n",LocateElem(&LL,&ee));

  return 0;
}

// 初始化顺序表
void InitList(PSeqList LL)
{
  ClearList(LL); // 清空顺序表。
}

// 清空顺序表。
void ClearList(PSeqList LL)
{
  if (LL == NULL) return;  // 检查空指针。

  LL->length=0;
  memset(LL->data,0,sizeof(ElemType)*MAXSIZE);
}

// 求顺序表的长度，返回值：>=0-表LL元素的个数。
int LengthList(PSeqList LL)
{
  if (LL == NULL) return 0;  // 检查空指针。

  return LL->length;
}

// 获取顺序表中第ii个元素的值，存放在ee中，返回值：0-失败；1-成功。
int  GetElem(PSeqList LL, unsigned int ii, ElemType* ee)
{
  if ( (LL == NULL) || (ee == NULL) ) return 0;   // 检查空指针。
  
  // 判断位置ii是否合法
  if ( (ii < 1) || (ii > LL->length) ) return 0;

  memcpy(ee,LL->data+ii-1,sizeof(ElemType));  // 采用memcpy是为了支持ee为结构体的情况。

  return ;
}

// 在顺序表LL的第ii个位置插入元素ee，返回值：0-失败；1-成功。
int InsertList(PSeqList LL, unsigned int ii, ElemType *ee)
{
  if ( (LL == NULL) || (ee == NULL) ) return 0;   // 检查空指针。

  if (LL->length >= MAXSIZE)
  {
    printf("顺序表已满，不能插入。\n"); return 0;
  }

  // 判断插入位置是否合法
  if ( (ii < 1) || (ii > LL->length+1) )
  {
    printf("插入位置（%d）不合法，应该在（%d-%d）之间。\n",ii,1,LL->length+1); return 0;
  }

  // 注意，元素后移只能用循环，不能用以下注释掉的方法，当元素是结构体时，以下方法不稳定。
  // if ( ii < LL->length+1)
  //   memcpy(LL->data+ii,LL->data+ii-1,(LL->length-ii+1)*sizeof(ElemType));

  // 把ii和ii之后的元素后移。
  int kk;
  for (kk=LL->length;kk>=ii;kk--)
  {
    memcpy(LL->data+kk,LL->data+kk-1,sizeof(ElemType));  // 采用memcpy是为了支持e为结构体的情况。
  }

  memcpy(LL->data+ii-1,ee,sizeof(ElemType));  // 采用memcpy是为了支持ee为结构体的情况。
  LL->length++;       // 表的长度加1。

  return 1;
}

// 在顺序表LL的头结点插入元素ee，返回值：0-失败；1-成功。
int  PushFront(PSeqList LL, ElemType *ee)
{
  return InsertList(LL,1,ee);
}

// 在顺序表LL的尾结点插入元素ee，返回值：0-失败；1-成功。
int PushBack(PSeqList LL, ElemType *ee)
{
  return InsertList(LL,LL->length+1,ee);
}

// 删除顺序表LL中的第ii个元素，返回值：0-位置ii不合法；1-成功。
int DeleteElem(PSeqList LL, unsigned int ii)
{
  if (LL == NULL) return 0;   // 检查空指针。

  // 判断删除位置ii是否合法
  if ( (ii < 1) || (ii > LL->length) )
  {
    printf("删除位置（%d）不合法，应该在（%d-%d）之间\n",ii,1,LL->length); return 0;
  }

  // 把ii之后的元素前移。
  int kk;
  for (kk=ii;kk<=LL->length;kk++)
  {
    memcpy(LL->data+kk-1,LL->data+kk,sizeof(ElemType));  // 采用memcpy是为了支持ee为结构体的情况。
  }

  LL->length--;       // 表的长度减1。

  return 1;
}

// 查找ee在顺序表LL中的位置，返回值：0-元素ee在表LL中不存在，>0元素ee在表LL中的位置。
int LocateElem(PSeqList LL, ElemType *ee)
{
  if ( (LL == NULL) || (ee == NULL) ) return 0;   // 检查空指针。

  int kk;

  for (kk = 0; kk < LL->length; kk++)
  {
    // 元素ee为结构体，这行代码与元素为整数时的方式不一样，比较超女结构体的编号。
    if ((LL->data+kk)->no == ee->no) return kk+1; // 在表中对应序号应为kk+1。
  }
    
  return 0;
}

// 删除顺序表LL中头元素，返回值：0-位置ii不合法；1-成功。
int PopFront(PSeqList LL)
{
  return DeleteElem(LL, 1);
}

// 删除顺序表LL中尾元素，返回值：0-位置ii不合法；1-成功。
int PopBack(PSeqList LL)
{
  return DeleteElem(LL, LL->length);
}

// 判断顺序表是否为空，返回值：1-空，0-非空或失败。
int IsEmpty(PSeqList LL)
{
  if (LL == NULL) return 0;   // 检查空指针。

  if (LL->length == 0) return 1;

  return 0;
}

// 打印顺序表中全部的元素。
void PrintList(PSeqList LL)
{
  if (LL == NULL) return;   // 检查空指针。

  if (LL->length == 0) { printf("表为空。\n"); return; }

  int kk;
  for (kk = 0; kk < LL->length; kk++)
  {
    // 元素ee为结构体，这行代码与元素为整数时的方式不一样，这里输出超女结构体的编号。
    printf("%-3d", (LL->data+kk)->no);  // 输出超女的编号。
  }

  printf("\n");
}

// 销毁顺序表LL。
void DestroyList(PSeqList LL)
{
  // 静态顺序表无需释放内存，不需要销毁操作。
  return;
}

