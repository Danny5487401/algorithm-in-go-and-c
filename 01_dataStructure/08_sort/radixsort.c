#include <stdlib.h>
#include <stdio.h>
#include <string.h>

// 获取数组arr中最大值，arr-待排序的数组，len-数组arr的长度。
int arrmax(int *arr,unsigned int len)
{
  int ii,imax;

  imax=arr[0];

  for (ii=1;ii<len;ii++)
    if (arr[ii]>imax) imax=arr[ii];

  return imax;
}

// 对数组arr按指数位进行排序。
// arr-待排序的数组，len-数组arr的长度。
// exp-排序指数，exp=1：按个位排序；exp=10：按十位排序；......
void _radixsort(int *arr,unsigned int len,unsigned int exp)
{
  int ii;
  int result[len];       // 存放从桶中收集后数据的临时数组。
  int buckets[10]={0};   // 初始化10个桶。

  // 遍历arr，将数据出现的次数存储在buckets中。
  for (ii=0;ii<len;ii++)
    buckets[(arr[ii]/exp)%10]++;

  // 调整buckets各元素的值，调整后的值就是arr中元素在result中的位置。
  for (ii=1;ii<10;ii++)
    buckets[ii]=buckets[ii]+buckets[ii-1];

  // 将arr中的元素填充到result中。
  for (ii=len-1;ii>=0;ii--)
  {
    int iexp=(arr[ii]/exp)%10;
    result[buckets[iexp]-1]=arr[ii];
    buckets[iexp]--;
  }
  
  // 将排序好的数组result复制到数组arr中。
  memcpy(arr,result,len*sizeof(int));
}

// 基数排序主函数，arr-待排序的数组，len-数组arr的长度。
void radixsort(int *arr,unsigned int len)
{
  int imax=arrmax(arr,len);    // 获取数组arr中的最大值。

  int iexp;    // 排序指数，iexp=1：按个位排序；iexp=10：按十位排序；......

  // 从个位开始，对数组arr按指数位进行排序。
  for (iexp=1;imax/iexp>0;iexp=iexp*10)
  {
    _radixsort(arr,len,iexp);
    int yy; printf("exp=%-5d  ",iexp); for (yy=0;yy<len;yy++) printf("%2d ",arr[yy]); printf("\n");
  }
}

int main(int argc,char *argv[])
{
  int arr[]={144,203,738,905,347,215,836,26,527,602,946,504,219,750,848};
  int len=sizeof(arr)/sizeof(int);

  radixsort(arr,len);  // 基数排序。

  // 显示排序结果。
  int yy; for (yy=0;yy<len;yy++) printf("%2d ",arr[yy]); printf("\n");

  // system("pause");  // widnows下的C启用本行代码。

  return 0;
}
