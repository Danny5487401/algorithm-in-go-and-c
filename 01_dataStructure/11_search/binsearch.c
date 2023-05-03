/*
 * 程序名：binsearch.c，此程序演示折半查找。
 * 作者：C语言技术网(www.freecplus.net) 日期：20210325
*/
#include <stdio.h>
#include <string.h>

// 迭代实现。
// 在sstable中查找key，失败返回-1，成功返回key在sstable中的数组下标。
int Bin_Search1(int *sstable,unsigned int len,int key)
{
  int low,high,mid;
  low=0; high=len-1;    // 初始化low和high的值，设置初始区间。

  while (low<=high)
  {
    mid=(low+high)/2;   // 计算mid的值，取中间位置。

    if (sstable[mid]==key) return mid;     // 找到了待查找的元素。
    else if (sstable[mid]>key) high=mid-1; // 继续在前半区间进行查找。
    else low=mid+1;                        // 继续在后半区间进行查找。
  }

  return -1;
}

// 递归实现。
// 在sstable中查找key，失败返回-1，成功返回key在sstable中的数组下标。
int Bin_Search2(int *sstable,int key,int low,int high)
{
  if (low>high) return -1;  

  int mid=(low+high)/2;   // 计算mid的值，取中间位置。

  if (sstable[mid]==key) return mid;         // 找到了待查找的元素。
  else if (sstable[mid]>key) Bin_Search2(sstable,key,low,mid-1); // 继续在前半区间进行查找。
  else Bin_Search2(sstable,key,mid+1,high);  // 继续在后半区间进行查找。
}

int main()
{
  // 有序的顺序表。
  int sstable1[]={1,2,3,4,5,6,7,8,9,12};
  int len=sizeof(sstable1)/sizeof(int);

  // 迭代
  printf("result1=%d\n",Bin_Search1(sstable1,len,0));
  printf("result1=%d\n",Bin_Search1(sstable1,len,1));
  printf("result1=%d\n",Bin_Search1(sstable1,len,7));
  printf("result1=%d\n",Bin_Search1(sstable1,len,10));
  printf("result1=%d\n",Bin_Search1(sstable1,len,12));
  printf("result1=%d\n",Bin_Search1(sstable1,len,15));
  printf("\n");

  // 递归
  printf("result2=%d\n",Bin_Search2(sstable1, 0,0,len-1));
  printf("result2=%d\n",Bin_Search2(sstable1, 1,0,len-1));
  printf("result2=%d\n",Bin_Search2(sstable1, 7,0,len-1));
  printf("result2=%d\n",Bin_Search2(sstable1,10,0,len-1));
  printf("result2=%d\n",Bin_Search2(sstable1,12,0,len-1));
  printf("result2=%d\n",Bin_Search2(sstable1,15,0,len-1));

  return 0;
}
