/*
 * 程序名：linkstack1.c，此程序演示链栈的实现，数据元素是整数。
 * 作者：C语言技术网(www.freecplus.net) 日期：20201230
*/
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef int ElemType;    // 自定义链栈的数据元素为整数。

typedef struct SNode
{
  ElemType data;       // 链栈中的元素。
  struct SNode *next;  // 指向下一个结点的指针。
}SNode,*LinkStack;

// 链栈SS的初始化操作，分配头结点，返回头结点的地址。
SNode *InitStack();

// 销毁链栈SS。
void DestroyStack(LinkStack SS);

// 元素入栈，返回值：0-失败；1-成功。
int Push(LinkStack SS, ElemType *ee);

// 元素出栈，返回值：0-失败；1-成功。
int Pop(LinkStack SS, ElemType *ee);

// 求链栈的长度，返回值：栈SS中元素的个数。
int  Length(LinkStack SS);                   

// 清空链栈。
void Clear(LinkStack SS);                    

// 判断链栈是否为空，返回值：1-空，0-非空或失败。
int  IsEmpty(LinkStack SS);                    

// 打印链栈中全部的元素。
void PrintStack(LinkStack SS);                    

// 获取栈顶元素，返回值：0-失败；1-成功。
// 只查看栈顶元素的值，元素不出栈。
int GetTop(LinkStack SS, ElemType *ee);

int main()
{
  LinkStack SS;     // 创建链栈。

  SS=InitStack();  // 初始化链栈。

  printf("栈的长度是%d\n",Length(SS));

  ElemType ee;     // 创建一个数据元素。

  printf("元素（1、2、3、4、5、6、7、8、9、10）入栈。\n");
  ee=1;  Push(SS, &ee);
  ee=2;  Push(SS, &ee);
  ee=3;  Push(SS, &ee);
  ee=4;  Push(SS, &ee);
  ee=5;  Push(SS, &ee);
  ee=6;  Push(SS, &ee);
  ee=7;  Push(SS, &ee);
  ee=8;  Push(SS, &ee);
  ee=9;  Push(SS, &ee);
  ee=10; Push(SS, &ee);

  printf("栈的长度是%d\n",Length(SS));

  if (GetTop(SS,&ee)==1)  printf("栈顶的元素值为%d\n",ee);

  PrintStack(SS);

  if (Pop(SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);
  if (Pop(SS,&ee)==1)  printf("出栈的元素值为%d\n",ee);

  // 销毁链栈SS。
  DestroyStack(SS); SS=0;  // 销毁链栈后把SS置为空，防止野指针。

  return 0;
}

// 链栈SS的初始化操作，分配头结点，返回头结点的地址。
SNode *InitStack()
{
  SNode *head = (SNode *)malloc(sizeof(SNode));  // 分配头结点。

  if (head == NULL) return NULL;  // 内存不足，返回失败。

  head->next=NULL;  // 头结点的下一结点暂时不存在，置空。

  return head;
}


// 清空链栈。
void Clear(LinkStack SS)
{
  // 清空链栈SS是指释放链栈全部的结点，但不包括头结点。
  if (SS == NULL) return;  // 检查空指针。

  SNode *tmp1;
  SNode *tmp2=SS->next;  // 保留头结点，从头结点的下一个结点开始释放。

  while(tmp2!=NULL)
  {
    tmp1=tmp2->next;
    free(tmp2);
    tmp2=tmp1;
  }

  SS->next=NULL; // 这行代码一定不能少，否则会留下野指针。

  return;
}

// 求链栈的长度，返回值：栈SS中元素的个数。
int  Length(LinkStack SS)
{
  if (SS == NULL) return 0; // 检查空指针。

  SNode *pp=SS->next;  // 头结点不算，从第1个结点开始。

  int length=0;

  while (pp != NULL) { pp=pp->next; length++; }

  return length;
}

// 销毁链栈SS。
void DestroyStack(LinkStack SS)
{
  // 销毁链栈SS是指释放链栈全部的结点，包括头结点。
  SNode *tmp;

  while(SS!=NULL)
  {
    tmp=SS->next;  // tmp保存下一结点的地址。
    free(SS);      // 释放当前结点。
    SS=tmp;        // SS指针移动到下一结点。
  }

  // SS=NULL;   // SS在本函数中相当于局部变量，就算置空了也不会影响调用者传递的SS，所以SS=NULL没有意义。

  return;
}

// 判断链栈是否为空，返回值：1-空，0-非空或失败。
int IsEmpty(LinkStack SS)
{
  if (SS == NULL) return 0;   // 检查空指针。

  if (SS->next == NULL) return 1;

  return 0;
}

// 打印链栈中全部的元素。
void PrintStack(LinkStack SS)
{
  if (SS == NULL) return; // 检查空指针。

  if (SS->next == NULL) { printf("栈为空。\n"); return; }

  int kk=0;
  SNode *pp=SS->next;  // 从第1个结点开始。

  while (pp != NULL)
  {
    printf("SS[%d],value=%d\n",kk++,pp->data);     
    pp=pp->next;
  }
}

// 元素出栈，返回值：0-失败；1-成功。
int Pop(LinkStack SS, ElemType *ee)
{
  if ( (SS == NULL) || (ee == NULL) ) return 0;  // 检查空指针。

  if (SS->next == NULL) { printf("栈为空。\n"); return 0; }

  SNode *tmp=SS->next;

  memcpy(ee,&tmp->data,sizeof(ElemType));  

  SS->next=tmp->next;

  free(tmp);

  return 1;
}

// 获取栈顶元素，返回值：0-失败；1-成功。
// 只查看栈顶元素的值，元素不出栈。
int GetTop(LinkStack SS, ElemType *ee)
{
  if ( (SS == NULL) || (ee == NULL) ) return 0;  // 检查空指针。

  if (SS->next == NULL) { printf("栈为空。\n"); return 0; }

  memcpy(ee,&SS->next->data,sizeof(ElemType));  

  return 1;
}

// 元素入栈，返回值：0-失败；1-成功。
int Push(LinkStack SS, ElemType *ee)
{
  if ( (SS == NULL) || (ee == NULL) ) return 0;  // 检查空指针。

  SNode *tmp = (SNode *)malloc(sizeof(SNode));  // 分配一个结点。
  if (tmp == NULL) return 0;  // 内存不足，返回失败。

  // 考虑数据元素为结构体的情况，这里采用了memcpy的方法而不是直接赋值。
  memcpy(&tmp->data,ee,sizeof(ElemType));

  // 处理next指针。
  tmp->next=SS->next;
  SS->next=tmp;
  
  return 1;
}
