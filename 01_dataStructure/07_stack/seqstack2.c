/*
 * 程序名：seqstack2.c，此程序演示用顺序栈检查括号是否匹配，支持()[]{}三种括号。
 * 作者：C语言技术网(www.freecplus.net) 日期：20201230
*/
#include <stdio.h>
#include <string.h>

#define MAXSIZE 100      // 顺序栈的最大长度。   // xxxxx

typedef char ElemType;    // 自定义顺序栈的数据元素为字符。 // xxxx

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

// 检查表达式括号是否匹配，返回值：0-不匹配，1-匹配。
int checkbrackets(char *str);

int main()
{
  char str[101];  
  memset(str,0,sizeof(str));

  printf("请输入待检查的字符串：");
  fgets(str,100,stdin);    // 不建议用gets函数，gets函数编译时可能会出现警号。
  str[strlen(str)-1]=0;    // 删除str最后的换行。
  printf("输入的字符串是=%s=\n",str);

  if (checkbrackets(str) == 0) { printf("匹配失败。\n"); return -1; }

  printf("匹配成功。\n");

  return 1;
}

// 检查表达式括号是否匹配，返回值：0-不匹配，1-匹配。
int checkbrackets(char *str)
{
  SeqStack SS;     // 创建顺序栈。
  InitStack(&SS);  // 初始化顺序栈。

  ElemType ee;     // 创建一个数据元素。

  int ipos=0,len=strlen(str);

  for (ipos=0;ipos<len;ipos++)
  {
    // 如果是左括号，入栈。
    if ( (str[ipos]=='(') || (str[ipos]=='[') || (str[ipos]=='{') ) Push(&SS,&str[ipos]);
    //if ( (str[ipos]=='(') || (str[ipos]=='[') || (str[ipos]=='{') ) Push(&SS,str+ipos); // 用指针运算也可以。

    // 如果是右括号，出栈，并判断是否匹配。
    if ( (str[ipos]==')') || (str[ipos]==']') || (str[ipos]=='}') ) 
    {
      // 一定要判断出栈结果，如果栈中没有元素，匹配失败。
      if (Pop(&SS,&ee) != 1) return 0; 

      // 只要出现了不匹配的情况，函数返回。
      if ( (str[ipos]==')') && (ee!='(') ) return 0; 
      if ( (str[ipos]==']') && (ee!='[') ) return 0; 
      if ( (str[ipos]=='}') && (ee!='{') ) return 0; 
    }
  }

  // 如果全部的括号都匹配，栈一定是空的。
  return IsEmpty(&SS);
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
    printf("SS[%d],value=%c\n",kk,SS->data[kk]);     // 用数组的下标访问。  // xxx
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
