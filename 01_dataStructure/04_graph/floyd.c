/*
 * 程序名：floyd.c，此程序演示弗洛伊德算法（floyd）。
 * 作者：C语言技术网(www.freecplus.net) 日期：20200202
*/
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define VEXNUM   5      // 图的顶点数。

void print(int AA[][VEXNUM]) // 显示矩阵。
{
  int ii,jj;
  for (ii=0;ii<VEXNUM;ii++)
  {
    for (jj=0;jj<VEXNUM;jj++)
      printf("%5d",AA[ii][jj]);

    printf("\n");
  }
}

int main()
{
  // 初始化矩阵。
  int AA[VEXNUM][VEXNUM]={\
        0  ,999,  1,999, 10,\
        999,  0,999,  1,  5,\
        999,  1,  0,999,  7,\
        999,999,999,  0,  1,\
        999,999,999,999,  0 };

  int path[VEXNUM][VEXNUM]={\
        -1,-1,-1,-1,-1,\
        -1,-1,-1,-1,-1,\
        -1,-1,-1,-1,-1,\
        -1,-1,-1,-1,-1,\
        -1,-1,-1,-1,-1};

  printf("init\n"); print(AA); print(path);  printf("\n");

  int kk,ii,jj;

  for (kk=0; kk<VEXNUM; kk++) {   // 遍历全部的顶点。
    printf("kk=%d\n",kk); 

    // 遍历整个矩阵，ii是行，jj是列。
    for (ii=0; ii<VEXNUM; ii++) {
      for (jj=0; jj<VEXNUM; jj++) {
        // 如果经过kk中转的路径更短，则更新矩阵。
        if ( AA[ii][kk] + AA[kk][jj] < AA[ii][jj]) {
          // 更新最短路径，记下中转顶点。
          AA[ii][jj] = AA[ii][kk] + AA[kk][jj]; 
          path[ii][jj] = kk;  
          printf("AA[%d][%d] updated.\n",ii,jj);
        }
      }
    }

    print(AA); print(path);  printf("\n");
  }
  
  // print(AA); print(path);

  return 0;
}

