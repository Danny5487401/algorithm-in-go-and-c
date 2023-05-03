// 优化后的选择排序，作者：C语言技术网（www.freecplus.net）码农有道。
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

  int ileft,iright;   // 每趟排序的最左和最右的位置。
  int ii;             // 每趟排序的元素位置计数器。
  int iminpos;        // 每趟循环选出的最小值的位置（数组的下标）。
  int imaxpos;        // 每趟循环选出的最大值的位置（数组的下标）。

  ileft=0; iright=len-1;  // ileft从0开始，iright从len-1开始。

  while (ileft<iright)
  {
    iminpos=imaxpos=ileft;

    for (ii=ileft;ii<=iright;ii++)  // 每趟循环从ileft和iright之间选取元素。
    {
      // 找出值更小的元素，记下它的位置。
      if (arr[ii]<arr[iminpos])  iminpos=ii;

      // 找出值更大的元素，记下它的位置。
      if (arr[ii]>arr[imaxpos])  imaxpos=ii;
    }

    // 如果本趟循环的最小的元素不是最左边的元素，则交换它们的位置。
    if (iminpos!=ileft) swap(&arr[ileft],&arr[iminpos]);

    // 如果imaxpos的位置是ileft，在上面的代码中ileft已被交换到了iminpos的位置。
    // 所以imaxpos的值要修改为iminpos。
    if (imaxpos==ileft) imaxpos=iminpos;

    // 如果本趟循环的最大的元素不是最右边的元素，则交换它们的位置。
    if (imaxpos!=iright) swap(&arr[iright],&arr[imaxpos]);

    ileft++;
    iright--;
  }
}



// 采用递归实现的方法。
// 参数arr是待排序数组的首地址，len是数组元素的个数。
void selectsort2(int *arr,unsigned int len)
{
  if (len<2) return;

  int ii;             // 每趟排序的元素位置计数器。
  int iminpos=0;      // 每趟循环选出的最小值的位置（数组的下标）。
  int imaxpos=0;      // 每趟循环选出的最大值的位置（数组的下标）。
  int ileft=0;
  int iright=len-1;

  for (ii=ileft;ii<=iright;ii++)  // 循环从ileft和iright之间选取元素。
  {
    // 找出值更小的元素，记下它的位置。
    if (arr[ii]<arr[iminpos])  iminpos=ii;

    // 找出值更大的元素，记下它的位置。
    if (arr[ii]>arr[imaxpos])  imaxpos=ii;
  }

  // 如果本趟循环的最小的元素不是最左边的元素，则交换它们的位置。
  if (iminpos!=ileft) swap(&arr[ileft],&arr[iminpos]);

  // 如果imaxpos的位置是ileft，在上面的代码中ileft已被交换到了iminpos的位置。
  // 所以imaxpos的值要修改为iminpos。
  if (imaxpos==ileft) imaxpos=iminpos;

  // 如果本趟循环的最大的元素不是最右边的元素，则交换它们的位置。
  if (imaxpos!=iright) swap(&arr[iright],&arr[imaxpos]);

  len=len-2;

  selectsort2(++arr,len);
}

int main(int argc,char *argv[])
{
  int arr[]={44,3,38,5,47,15,36,26,27,2,46,4,19,50,48};
  int len=sizeof(arr)/sizeof(int);

  // selectsort1(arr,len);  // 采用两层循环排序的方法。
  selectsort2(arr,len);  // 采用递归排序的方法。

  // 显示排序结果。
  int jj; for (jj=0;jj<len;jj++) printf("%2d ",arr[jj]); printf("\n");

  // system("pause");  // widnows下的C启用本行代码。

  return 0;
}
