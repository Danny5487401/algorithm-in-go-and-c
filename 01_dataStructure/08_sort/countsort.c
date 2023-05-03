// 计数排序（基础版），作者：C语言技术网（www.freecplus.net）码农有道。
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// 获取待排序数组的最大元素的值。
int arrmax(int *arr,unsigned int len)
{
  int ii=0;
  int imax=0;

  for (ii=0;ii<len;ii++) if (imax<arr[ii]) imax=arr[ii];

  return imax;
}

// 计数排序主函数，arr-待排序数组的地址，len-数组的长度。
void countsort(int *arr,unsigned int len) 
{
  if (len<2) return;

  int imax=arrmax(arr,len);  // 获取待排序数组的最大元素的值。
  int arrtmp[imax+1];        // 临时数组的大小为imax+1。

  memset(arrtmp,0,sizeof(arrtmp));  // 初始化临时数组。

  int ii,jj,kk;
  
  // 临时数组计数。
  for (ii=0;ii<len;ii++) arrtmp[arr[ii]]++;

  // 把临时数组计数的内容填充到arr中。
  ii=0;
  for (jj=0;jj<imax+1;jj++)
  {
    for (kk=0;kk<arrtmp[jj];kk++) arr[ii++]=jj;
  }
}

int main() 
{
  int arr[]={2,3,8,7,1,2,2,2,7,3,9,8,2,1,4,2,4,6,9,2};

  int len=sizeof(arr)/sizeof(int);

  int xx; for (xx=0;xx<len;xx++) printf("%2d ",arr[xx]); printf("\n");

  countsort(arr,len);

  // 显示排序结果。
  int yy; for (yy=0;yy<len;yy++) printf("%2d ",arr[yy]); printf("\n");

  // system("pause");  // widnows下的C启用本行代码。

  return 0;
}
