/*
 * 程序名：seqstack1.c，此程序演示顺序栈的实现，数据元素是整数。
 * 作者：C语言技术网(www.freecplus.net) 日期：20201230
*/
#include <stdio.h>
#include <string.h>

#define MAXSIZE 10       // 顺序栈的最大长度。

typedef int ElemType;    // 自定义顺序栈的数据元素为整数。

typedef struct
{
  ElemType data[MAXSIZE];   // 用数组存储顺序栈中的元素。
  int top;                  // 栈顶指针，从0到MAXSIZE-1，-1表示空栈。
                            // 也可以从1到MAXSIZE，0表示空栈。
}SeqStack,*PSeqStack;

// 顺序栈SS的初始化操作。
void InitStack(PSeqStack SS);                     

// 销毁顺序栈SS。
void DestroyStack(PSeqStack SS);

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

int main()
{
  SeqStack SS;     // 创建顺序栈。

  InitStack(&SS);  // 初始化顺序栈。

  printf("栈的长度是%d\n",Length(&SS));

  ElemType ee;     // 创建一个数据元素。

  printf("元素（1、2、3、4、5、6、7、8、9、10）入栈。\n");
  ee=1;  Push(&SS, &ee);
  ee=2;  Push(&SS, &ee);
  ee=3;  Push(&SS, &ee);
  ee=4;  Push(&SS, &ee);
  ee=5;  Push(&SS, &ee);
  ee=6;  Push(&SS, &ee);
  ee=7;  Push(&SS, &ee);
  ee=8;  Push(&SS, &ee);
  ee=9;  Push(&SS, &ee);
  ee=10; Push(&SS, &ee);

  printf("栈的长度是%d\n",Length(&SS));

  // 只查看栈顶元素的值，元素不出栈。
  if (GetTop(&SS,&ee)==1)  printf("栈顶的元素值为%d\n",ee);

  PrintStack(&SS);

  if (Pop(&SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(&SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(&SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(&SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(&SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(&SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(&SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(&SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(&SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(&SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(&SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(&SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);

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

// 销毁顺序栈SS。
void DestroyStack(PSeqStack SS)
{
  // 静态顺序栈无需释放内存，不需要销毁操作。

  Clear(SS); // 清空顺序栈。

  return;
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
