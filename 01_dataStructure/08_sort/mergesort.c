// 采用递归的方法实现归并排序，作者：C语言技术网（www.freecplus.net）码农有道。
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

// 采用递归的方法实现归并排序函数。
// arr-待排序数组的首地址，arrtmp-用于排序的临时数组的首地址
// start-排序区间第一个元素的位置，end-排序区间最后一个元素的位置。
void _mergesort(int *arr,int *arrtmp,int start,int end) 
{
  // 如果start>=end，表示该区间的元素少于两个，递归终止。
  if (start>=end) return; 

  int mid=start+(end-start)/2;  // 计算排序区间中间的位置。

  int istart1=start,iend1=mid;  // 区间左边元素的第一和最后一个元素的位置。
  int istart2=mid+1,iend2=end;  // 区间右边元素的第一和最后一个元素的位置。

  _mergesort(arr,arrtmp,istart1,iend1);   // 对区间左边元素递归排序。
  _mergesort(arr,arrtmp,istart2,iend2);   // 对区间右边元素递归排序。

  int ii=start; // 已排序数组arrtmp的计数器。

  // 把区间左右两边数列合并到已排序数组arrtmp中。
  while (istart1<=iend1 && istart2<=iend2)
    arrtmp[ii++]=arr[istart1]<arr[istart2] ? arr[istart1++] : arr[istart2++];

  // 把左边数列其它的元素追加到已排序数组。
  while (istart1<=iend1) arrtmp[ii++]=arr[istart1++];

  // 把右边数列其它的元素追加到已排序数组。
  while (istart2<=iend2) arrtmp[ii++]=arr[istart2++];

  // 把已排序数组arrtmp中的元素复制到arr中。
  memcpy(arr+start,arrtmp+start,(end-start+1)*sizeof(int));
}

// 排序主函数，arr为待排序的数组的地址，len为数组的长度。
void mergesort(int *arr,unsigned int len) 
{
  if (len<2) return;  // 小于两个元素不需要排序。

  int arrtmp[len];  // 分配一个与待排序数组相同大小的数组。

  _mergesort(arr,arrtmp,0,len-1);  // 调用递归函数进行排序。
}

int main(int argc,char *argv[])
{
  int arr[]={44,3,38,5,47,15,36,26,27,2,46,4,19,50,48};
  int len=sizeof(arr)/sizeof(int);

  mergesort(arr,len);

  // 显示排序结果。
  int yy; for (yy=0;yy<len;yy++) printf("%2d ",arr[yy]); printf("\n");

  // system("pause");  // widnows下的C启用本行代码。

  return 0;
}

