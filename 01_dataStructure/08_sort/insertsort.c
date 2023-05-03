// 插入排序，作者：C语言技术网（www.freecplus.net）码农有道。
#include <stdlib.h>
#include <stdio.h>

// 参数arr是待排序数组的首地址，len是数组元素的个数。
void insertsort(int *arr,unsigned int len)
{
  if (len<2) return; // 数组小于2个元素不需要排序。

  int itmp;  // 当前需要排序的元素的值。
  int ii;    // 需要排序元素的计数器。
  int jj;    // 插入排序时，需要后移元素的计数器。

  for (ii=1;ii<len;ii++)
  {
    itmp=arr[ii];    // 待排序元素

    // 从已排序的最右边开始，把大于当前排序的元素后移。
    // for (jj=ii-1;(jj>=0&&arr[jj]>itmp);jj--)
    for (jj=ii-1;jj>=0;jj--)
    {
      if (arr[jj]<=itmp) break;

      arr[jj+1]=arr[jj]; // 逐个元素后移。
    }

    arr[jj+1]=itmp; // 插入当前排序元素。
  }
}


int main(int argc,char *argv[])
{
  int arr[]={44,3,38,5,47,15,36,26,27,2,46,4,19,50,48};
  int len=sizeof(arr)/sizeof(int);

  insertsort(arr,len);   // 调用插入排序函数对数组排序。

  // 显示排序结果。
  int yy; for (yy=0;yy<len;yy++) printf("%2d ",arr[yy]); printf("\n");

  // system("pause");  // widnows下的C启用本行代码。

  return 0;
}
