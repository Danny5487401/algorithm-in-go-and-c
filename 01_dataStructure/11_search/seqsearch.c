/*
 * 程序名：seqsearch.c，此程序演示顺序查找。
 * 作者：C语言技术网(www.freecplus.net) 日期：20210325
*/
#include <stdio.h>
#include <string.h>

// 对无序的查找表进行顺序查找。
// 在sstable中查找key，失败返回-1，成功返回key在sstable中的数组下标。
int Seq_Search1(int *sstable,unsigned int len,int key)
{
  int ii;

  for (ii=0;ii<len;ii++)  // 从前往后或从后往前找都行。
    if (sstable[ii]==key) break;  // 找到了就break

  if (ii==len) return -1;  // 查找失败时，ii==len

  return ii;
}

/*
// 对无序的查找表进行顺序查找更精简的写法。
// 在sstable中查找key，失败返回-1，成功返回key在sstable中的数组下标。
int Seq_Search1(int *sstable,unsigned int len,int key)
{
  int ii;

  for (ii=0;ii<len&&sstable[ii]!=key;ii++)  // 从前往后或从后往前找都行。
    ;

  return ii==len?-1:ii; 
}
*/

// 设置哨兵对无序的查找表进行顺序查找。
// 在sstable中查找key，失败返回0，成功返回key在sstable中的数组下标。
int Seq_Search2(int *sstable,unsigned int len,int key)
{
  int ii;

  sstable[0]=key;  // 哨兵。

  for (ii=len-1;sstable[ii]!=key;ii--)  // 从后往前找。
    ;                                   // 注意这个空语句。

  return ii;  // 找不到时，ii为0
}

// 对有序的查找表进行顺序查找。
// 在sstable中查找key，失败返回-1，成功返回key在sstable中的数组下标。
int Seq_Search3(int *sstable,unsigned int len,int key)
{
  int ii;

  for (ii=0;ii<len;ii++)  // 从前往后或从后往前找都行。
  {
    if (sstable[ii]==key) break;  // 找到了就break
    if (sstable[ii] >key) return -1; // 不必再找了，返回-1
  }

  if (ii==len) return -1;  // 查找失败时，ii==len

  return ii;
}

// 采用跳跃的方法对有序的查找表进行顺序查找。
// 在sstable中查找key，失败返回-1，成功返回key在sstable中的数组下标。
int Seq_Search4(int *sstable,unsigned int len,int key)
{
  int ii=0;

  while (ii<len && sstable[ii]<key) ii=ii+3;  // 每次跳跃三个元素。

  if (ii<len && sstable[ii]==key) return ii;  // 查找成功。 
  else if (ii-1<len && sstable[ii-1]==key) return ii-1;  // 查找成功。
  else if (ii-2<len && sstable[ii-2]==key) return ii-2;  // 查找成功。

  return -1;  // 查找失败。
}

int main()
{
  // 对无序的查找表进行顺序查找。
  int sstable1[]={2,5,6,3,1,7,4,8,9};
  int len=sizeof(sstable1)/sizeof(int);

  printf("result1=%d\n",Seq_Search1(sstable1,len,7));
  printf("result1=%d\n",Seq_Search1(sstable1,len,15));
  printf("result1=%d\n",Seq_Search1(sstable1,len,2));
  printf("result1=%d\n",Seq_Search1(sstable1,len,9));
  printf("\n");

  // 设置哨兵对无序的查找表进行顺序查找。
  int sstable2[]={0,2,5,6,3,1,7,4,8,9};
  len=sizeof(sstable2)/sizeof(int);

  printf("result2=%d\n",Seq_Search2(sstable2,len,7));
  printf("result2=%d\n",Seq_Search2(sstable2,len,15));
  printf("result2=%d\n",Seq_Search2(sstable2,len,2));
  printf("result2=%d\n",Seq_Search2(sstable2,len,9));
  printf("\n");

  // 对有序的查找表进行顺序查找。
  int sstable3[]={1,2,3,4,5,6,7,8,9};
  len=sizeof(sstable3)/sizeof(int);

  printf("result3=%d\n",Seq_Search3(sstable3,len,7));
  printf("result3=%d\n",Seq_Search3(sstable3,len,15));
  printf("result3=%d\n",Seq_Search3(sstable3,len,1));
  printf("result3=%d\n",Seq_Search3(sstable3,len,9));
  printf("\n");

  // 采用跳跃的方法对有序的查找表进行顺序查找。
  int sstable4[]={1,2,3,4,5,6,7,8,9};
  len=sizeof(sstable4)/sizeof(int);

  printf("result4=%d\n",Seq_Search4(sstable4,len,0));
  printf("result4=%d\n",Seq_Search4(sstable4,len,7));
  printf("result4=%d\n",Seq_Search4(sstable4,len,15));
  printf("result4=%d\n",Seq_Search4(sstable4,len,1));
  printf("result4=%d\n",Seq_Search4(sstable4,len,9));
  printf("\n");

  return 0;
}
