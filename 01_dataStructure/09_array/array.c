#include <stdio.h>

typedef  int ElemType;

int main()
{
  int len=20;
  ElemType  arr[len];           // 数组的大小可以用常量，也可以用变量。

  printf("%p\n",arr+0);
  printf("%p\n",arr+1);
  printf("%p\n",arr+2);
  printf("%p\n",arr+3);
  printf("%p\n",arr+4);
}
