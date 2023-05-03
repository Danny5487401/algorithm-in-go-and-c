// 采用循环的方法实现归并排序函数。，作者：C语言技术网（www.freecplus.net）码农有道。
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

int min(int x,int y) { return x<y ? x : y; }  // 取x和y中的较小者的值。

// 采用循环实现归并排序，arr-待排序数组的首地址，len--待排序数组的长度。
void mergesort(int *arr,unsigned int len) 
{
  if (len<2) return;  // 小于两个元素不需要排序。

  int *aa=arr;     // aa指向待排序的数组。
  int *bb=(int *)malloc(len*sizeof(int)); // bb指向已排序的数组。

  int iseg;    // 区间分段的计数器，1,2,4,8,16,...
  int istart;  // 区间起始位置的计数器。

  // 排序的趟数的循环。
  for (iseg=1;iseg<len;iseg=iseg*2) 
  {
    // 每趟排序选取区间的循环。
    for (istart=0;istart<len;istart=istart+iseg*2) 
    {
      // 把每个区间分成两部分，ilow是起始位置，imid是中间位置，imax是结束位置。
      int ilow=istart;
      int imid=min(istart+iseg,len);     // 考虑分段不均的情况，imid不能超出len。
      int imax=min(istart+iseg*2,len);   // 考虑分段不均的情况，imax也不能超出len。

      int ii=ilow; // 已排序数组的计数器。
      int istart1=ilow,iend1=imid;  // 待排序左边数列的起始和结束位置。
      int istart2=imid,iend2=imax;  // 待排序右边数列的起始和结束位置。

      // 把待排序左右两边数列合并到已排序数组。
      while ((istart1<iend1) && (istart2<iend2))
        bb[ii++]=aa[istart1]<aa[istart2] ? aa[istart1++] : aa[istart2++];

      // 把左边数列其它的元素追加到已排序数组。
      while (istart1<iend1) bb[ii++]=aa[istart1++];

      // 把右边数列其它的元素追加到已排序数组。
      while (istart2<iend2) bb[ii++]=aa[istart2++];
    }

    // 交换一下两个数组的指针，准备下一趟的排序。
    int *ptmp=aa; aa=bb; bb=ptmp;
  }

  // 如果aa指向的不是原始数组的指针，把aa中的内容复制到arr中。
  if (aa != arr) 
  {
    memcpy(arr,aa,len*sizeof(int));

    bb = aa;
  }

  free(bb);
}

int main(int argc,char *argv[])
{
  int arr[]={44,3,38,5,47,15,36,26,27,2,46,4,19,50,48,10};
  int len=sizeof(arr)/sizeof(int);

  mergesort(arr,len);

  // 显示排序结果。
  int yy; for (yy=0;yy<len;yy++) printf("%2d ",arr[yy]); printf("\n");

  // system("pause");  // widnows下的C启用本行代码。

  return 0;
}

