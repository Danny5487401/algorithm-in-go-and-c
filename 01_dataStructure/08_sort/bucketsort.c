// 桶排序，作者：C语言技术网（www.freecplus.net）码农有道。
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

// 采用两层循环实现冒泡排序的方法。
// 参数arr是待排序数组的首地址，len是数组元素的个数。
void bubblesort(int *arr,unsigned int len)
{
  if (len<2) return; // 数组小于2个元素不需要排序。

  int ii;    // 排序的趟数的计数器。
  int jj;    // 每趟排序的元素位置计数器。
  int itmp;  // 比较两个元素大小时交换位置用到的临时变量。

  // 44,3,38,5,47,15,36,26,27,2,46,4,19,50,48  
  for (ii=len-1;ii>0;ii--)  // 一共进行len-1趟比较。
  {
    for (jj=0;jj<ii;jj++)  // 每趟只需要比较0......ii之间的元素，ii之后的元素是已经排序好的。
    {
      if (arr[jj]>arr[jj+1])  // 如果前面的元素大于后面的元素，则交换它位的位置。
      {
        itmp=arr[jj+1];
        arr[jj+1]=arr[jj];
        arr[jj]=itmp;
      }
    }
  }
}

// 桶排序主函数，参数arr是待排序数组的首地址，len是数组元素的个数。
void bucketsort(int *arr,unsigned int len)
{
  int buckets[5][5];   // 分配五个桶。
  int bucketssize[5];  // 每个桶中元素个数的计数器。

  // 初始化桶和桶计数器。
  memset(buckets,0,sizeof(buckets));
  memset(bucketssize,0,sizeof(bucketssize));

  // 把数组arr的数据放入桶中。
  int ii=0;
  for (ii=0;ii<len;ii++)
  {
    buckets[arr[ii]/10][bucketssize[arr[ii]/10]++]=arr[ii];
  }

  // 对每个桶进行冒泡排序。
  for (ii=0;ii<5;ii++)
  {
    bubblesort(buckets[ii],bucketssize[ii]);
  }

  // 把每个桶中的数据填充到数组arr中。
  int jj=0,kk=0;
  for (ii=0;ii<5;ii++)
  {
    for (jj=0;jj<bucketssize[ii];jj++)
      arr[kk++]=buckets[ii][jj];
  }
}

int main(int argc,char *argv[])
{
  int arr[]={21,3,30,44,15,36,6,10,9,19,25,48,5,23,47};
  int len=sizeof(arr)/sizeof(int);

  int xx; for (xx=0;xx<len;xx++) printf("%2d ",arr[xx]); printf("\n");

  bucketsort(arr,len);

  // 显示排序结果。
  int ii; for (ii=0;ii<len;ii++) printf("%2d ",arr[ii]); printf("\n");

  // system("pause");  // widnows下的C启用本行代码。

  return 0;
}
