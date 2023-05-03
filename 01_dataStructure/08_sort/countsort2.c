// 全省高考成绩排名算法，作者：C语言技术网（www.freecplus.net）码农有道。
#include <stdio.h>

int main() 
{
  // 模拟全省考生的成绩。
  int score[]={581,575,632,581,750,522,576,581,580,582,630,710,508,749,250};

  int len=sizeof(score)/sizeof(int); // 考生总人数。

  int buckets[751]={0};        // 准备751个桶。

  int ii;    // 循环的计数器。

  // 把每名学生的成绩计数。
  for (ii=0;ii<len;ii++) buckets[score[ii]]++;  

  // 计算出每个分数的全省排名。
  for (ii=750-1;ii>=0;ii--) buckets[ii]=buckets[ii]+buckets[ii+1];

  int itmp=0;
  printf("请输入考生的成绩：");
  scanf("%d",&itmp);
  printf("该考生全省排名是：%d\n",buckets[itmp]);

  // system("pause");  // widnows下的C启用本行代码。

  return 0;
}
