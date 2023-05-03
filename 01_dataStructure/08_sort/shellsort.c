// 希尔排序，作者：C语言技术网（www.freecplus.net）码农有道。
#include <stdlib.h>
#include <stdio.h>

// 对希尔排序中的单个组进行排序。
// arr-待排序的数组，len-数组总的长度，ipos-分组的起始位置，istep-分组的步长（增量）。
void groupsort(int *arr, int len, int ipos,int istep)
{
  int itmp;  // 当前需要排序的元素的值。
  int ii;    // 需要排序元素的计数器。
  int jj;    // 插入排序时，需要后移元素的计数器。

  for (ii=ipos+istep;ii<len;ii=ii+istep) 
  {
    itmp=arr[ii];    // 待排序元素
      
    // 从已排序的最右边开始，把大于当前排序的元素后移。
    // for (jj=ii-istep;(jj>=0&&arr[jj]>itmp);jj=jj-istep)
    for (jj=ii-istep;jj>=0;jj=jj-istep)
    {
      if (arr[jj]<=itmp) break;  

      arr[jj+istep]=arr[jj];  // 逐个元素后移。
    }

    arr[jj+istep]=itmp; // 插入当前排序元素。
  }
}

// 希尔排序，arr是待排序数组的首地址，len是数组的大小。
void shellsort(int *arr,unsigned int len)
{
  int ii,istep;

  // istep为步长，每次减为原来的一半取整数，最后一次必定为1。
  for (istep=len/2;istep>0;istep=istep/2)
  {
    // 共istep个组，对每一组都执行插入排序。
    for (ii=0;ii<istep;ii++)
    {
      groupsort(arr,len,ii,istep);
    }
  }
}

int main(int argc,char *argv[])
{
  int arr[]={44,3,38,5,47,15,36,26,27,2,46,4,19,50,48};
  int len=sizeof(arr)/sizeof(int);

  shellsort(arr,len);   // 调用插入排序函数对数组排序。

  // 显示排序结果。
  int yy; for (yy=0;yy<len;yy++) printf("%2d ",arr[yy]); printf("\n");

  // system("pause");  // widnows下的C启用本行代码。

  return 0;
}
