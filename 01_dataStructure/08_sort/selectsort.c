// 选择排序，作者：C语言技术网（www.freecplus.net）码农有道。
#include <stdlib.h>
#include <stdio.h>

// 交换两个变量的值。
void swap(int *x,int *y)
{
  int itmp=*x;
  *x=*y;
  *y=itmp;
}

// 采用两层循环实现的方法。
// 参数arr是待排序数组的首地址，len是数组元素的个数。
void selectsort1(int *arr,unsigned int len)
{
  if (len<2) return; // 数组小于2个元素不需要排序。

  int ii;      // 排序的趟数的计数器。
  int jj;      // 每趟排序的元素位置计数器。
  int iminpos; // 每趟循环选出的最小值的位置（数组的下标）。

  // 44,3,38,5,47,15,36,26,27,2,46,4,19,50,48  
  for (ii=0;ii<len-1;ii++)  // 一共进行len-1趟比较。
  {
    iminpos=ii;

    for (jj=ii+1;jj<len;jj++)  // 每趟只需要比较ii+1......len-1之间的元素，ii之前的元素是已经排序好的。
    {
      // 找出值更小的元素，记下它的位置。
      if (arr[jj]<arr[iminpos])  iminpos=jj;
    }

    // 如果本趟循环的最小的元素不是起始位置的元素，则交换它们的位置。
    if (iminpos!=ii) swap(&arr[ii],&arr[iminpos]);
  }
}

// 采用递归实现的方法。
// 参数arr是待排序数组的首地址，len是数组元素的个数。
void selectsort2(int *arr,unsigned int len)
{
  if (len<2) return; // 数组小于2个元素不需要排序。

  int ii;        // 排序的趟数的计数器。
  int iminpos=0; // 每趟循环选出的最小值的位置（数组的下标）。

  for (ii=1;ii<len;ii++)  
  {
    // 找出值更小的元素，记下它的位置。
    if (arr[ii]<arr[iminpos])  iminpos=ii;
  }

  // 如果本趟循环的最小的元素不是起始位置的元素，则交换它们的位置。
  if (iminpos!=0) swap(&arr[0],&arr[iminpos]);

  selectsort2(arr+1,--len);
}

int main(int argc,char *argv[])
{
  int arr[]={44,3,38,5,47,15,36,26,27,2,46,4,19,50,48};
  //int arr[]={3,4,5,1};
  int len=sizeof(arr)/sizeof(int);

  // selectsort1(arr,len);  // 采用两层循环排序的方法。
  selectsort2(arr,len);  // 采用递归排序的方法。

  // 显示排序结果。
  int ii; for (ii=0;ii<len;ii++) printf("%2d ",arr[ii]);

  printf("\n");

  // system("pause");  // widnows下的C启用本行代码。

  return 0;
}
