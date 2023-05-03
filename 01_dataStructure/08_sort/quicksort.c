// 快速排序，作者：C语言技术网（www.freecplus.net）码农有道。
#include <stdlib.h>
#include <stdio.h>

void quicksort(int *arr,unsigned int len)
{
  if (len<2) return;  // 数组的元素小于2个就不用排序了。

  int itmp=arr[0];  // 选取最左边的数作为中心轴。
  int ileft=0;      // 左下标。
  int iright=len-1; // 右下标。
  int imoving=2;    // 当前应该移动的下标，1-左下标；2-右下标。

  while (ileft<iright)
  {
    if (imoving==2) // 移动右下标的情况。
    {
      // 如果右下标位置元素的值大于等于中心轴，继续移动右下标。
      if (arr[iright]>=itmp) { iright--; continue; }
      
      // 如果右下标位置元素的值小于中心轴，把它填到左下标的坑中。
      arr[ileft]=arr[iright];
      ileft++;    // 左下标向右移动。
      imoving=1;  // 下次循环将移动左下标。
      continue;
    }

    if (imoving==1) // 移动左下标的情况。
    {
      // 如果左下标位置元素的值小等于中心轴，继续移动左下标。
      if (arr[ileft]<=itmp) { ileft++; continue; }

      // 如果左下标位置元素的值大于中心轴，把它填到右下标的坑中。
      arr[iright]=arr[ileft];
      iright--;   // 右下标向左移动。
      imoving=2;  // 下次循环将移动右下标。
      continue;
    }
  }

  // 如果循环结束，左右下标重合，把中心轴的值填进去。
  arr[ileft]=itmp;

  quicksort(arr,ileft);             // 对中心轴左边的序列进行排序。
  quicksort(arr+ileft+1,len-ileft-1); // 对中心轴右边的序列进行排序。
}

int main(int argc,char *argv[])
{
  int arr[]={44,3,38,5,47,15,36,26,27,2,46,4,19,50,48};
  int len=sizeof(arr)/sizeof(int);

  quicksort(arr,len);   // 调用插入排序函数对数组排序。

  // 显示排序结果。
  int yy; for (yy=0;yy<len;yy++) printf("%2d ",arr[yy]); printf("\n");

  // system("pause");  // widnows下的C启用本行代码。

  return 0;
}
