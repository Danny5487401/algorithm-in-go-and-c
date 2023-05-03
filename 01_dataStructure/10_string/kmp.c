/*
 * 程序名：kmp.c，此程序演示了字符串的匹配模式算法，包括：
 * 1）普通的模式匹配算法，即BF算法；
 * 2）KMP算法的求next数组；
 * 3）KMP算法的求nextval数组；
 * 4）KMP算法的实现。
 * 作者：C语言技术网(www.freecplus.net) 日期：20210125
*/
#include <stdio.h>
#include <string.h>

// 采用BF算法，查找在目标串sstr中模式串tstr出现的位置，字符串的起始位置从0开始。
// 只要在目标串sstr中找到了第一个模式串tstr，函数就返回。
// 成功返回模式串tstr在目标串sstr中第一次出现的数组下标，失败返回-1。
int index_bf(char sstr[],char tstr[])
// int index_bf(char *sstr,char *tstr)
{
  if ( (sstr==0) || (tstr==0) ) return -1;  // 判断空指针。

  // 目标串sstr和模式串tstr都从数组下标0开始。
  int ii=0,jj=0;  // 循环从0开始。
  int slen=strlen(sstr),tlen=strlen(tstr);
  
  while ( (ii<slen) && (jj<tlen) )
  {
    if (sstr[ii] == tstr[jj]) 
    {
      ii++; jj++;  // 继续比较后面的字符。
    }
    else
    {
      ii=ii-jj+1; jj=0; // 目标串和模式串的指针都回溯。
    }
  }

  // 匹配成功。
  if (jj==tlen) return ii-tlen;

  return -1;
}

// 根据模式串tstr，计算next数组，字符串的起始位置从0开始。
// void getnext(char *tstr,int *next)
void getnext(char tstr[],int next[])
{
  if ( (tstr==0) || (next==0) ) return;  // 判断空指针。

  int tlen=strlen(tstr);  // 模式串的长度。

  // 如果模式串的长度小于3，无需计算next数组，直接返回。
  if (tlen==0) return;
  if (tlen==1) { next[0]=-1; return; }
  if (tlen==2) { next[0]=-1; next[1]=0; return; }

  next[0]=-1;  // 位置0直接填入-1。
  next[1]=0;   // 位置1直接填入0。

  // 如果想调试本函数，放开// xxx就可以了。

  int ii;

  // 循环从2开始，位置0和1不必计算。
  for (ii=2;ii<tlen;ii++)
  {
    // 显示子串。
    // xxx char p0[tlen]; 
    // xxx memset(p0,0,sizeof(p0));
    // xxx strncpy(p0,tstr,ii); 
    // xxx printf("p0=%s  ",p0);

    int maxlen=0;  // 最长公共前后缀的长度。

    // 用一个循环，得到全部的前缀和后缀，并判断最长公共前后缀。
    int jj=1;
    for (jj=1;jj<ii;jj++)
    { 
      char p1[tlen],p2[tlen]; 
      memset(p1,0,sizeof(p1));
      memset(p2,0,sizeof(p2));

      strncpy(p1,tstr,jj);       // 取前缀。
      strncpy(p2,tstr+ii-jj,jj); // 取后缀。

      // xxx printf("%s %s    ",p1,p2);

      // 判断前后缀是否相同，相同就记下jj的值。
      if (strcmp(p1,p2)==0) maxlen=jj;
    }

    // xxx printf(" maxlen=%d\n",maxlen);  
    next[ii]=maxlen;
  }

  // xxx printf("next "); int kk;for (kk=0;kk<strlen(tstr);kk++) printf("%d ",next[kk]); printf("\n");

  return ;
}

// 根据模式串tstr和next数组计算nextval数组，字符串的起始位置从0开始。
// void getnextval(char *tstr,int *next,int *nextval)
void getnextval(char tstr[],int next[],int nextval[])
{
  if ( (tstr==0) || (next==0) || (nextval==0) ) return;  // 判断空指针。

  int tlen=strlen(tstr);  // 模式串的长度。
  
  nextval[0]=-1;   // 位置0直接填入-1。

  int ii;

  // 从第1个位置开始扫描。
  for (ii=1;ii<tlen;ii++)
  {
    if (tstr[ii]==tstr[next[ii]]) 
      nextval[ii]=nextval[next[ii]];
    else
      nextval[ii]=next[ii];
  }
}

// 采用kmp算法，查找在目标串sstr中模式串tstr出现的位置，字符串的起始位置从0开始。
// 只要在目标串sstr中找到了第一个模式串tstr，函数就返回。
// 成功返回模式串tstr在目标串sstr中第一次出现的数组下标，失败返回-1。
// int index_kmp(char *sstr,char *tstr)
int index_kmp(char sstr[],char tstr[])
{
  if ( (sstr==0) || (tstr==0) ) return -1;  // 判断空指针。

  // 目标串sstr和模式串tstr都从数组下标0开始。
  int ii=0,jj=0;  
  int slen=strlen(sstr),tlen=strlen(tstr);
  
  // 获取next数组。
  int next[tlen]; 
  getnext(tstr,next);

  // 根据模式串tstr和next数组计算nextval数组，字符串的起始位置从0开始。
  int nextval[tlen]; 
  getnextval(tstr,next,nextval);

  while ( (ii<slen) && (jj<tlen) )
  {
    if ( (jj==-1) || (sstr[ii] == tstr[jj]) )
    {
      ii++; jj++;  // 继续比较后面的字符。
    }
    else
    {
      // jj=next[jj]; // 根据next数组模式串指针回溯。
      jj=nextval[jj]; // 根据nextval数组模式串指针回溯。
    }
  }

  // 匹配成功。
  if (jj==tlen) return ii-tlen;

  return -1;
}

int main()
{
  //////////////////////////////////
  // 获取模式串的next数组，字符串的位置从0开始。
  char tmp[100];
  memset(tmp,0,sizeof(tmp));
  strcpy(tmp,"baabaa");
  int next[100];
  getnext(tmp,next);
  printf("next    "); int kk;for (kk=0;kk<strlen(tmp);kk++) printf("%d ",next[kk]); printf("\n");
  //////////////////////////////////

  //////////////////////////////////
  // 根据模式串tstr和next数组计算nextval数组，字符串的起始位置从0开始。
  int nextval[100];
  getnextval(tmp,next,nextval);
  printf("nextval "); int ll;for (ll=0;ll<strlen(tmp);ll++) printf("%d ",nextval[ll]); printf("\n");
  //////////////////////////////////

  //////////////////////////////////
  // 判断模式串和目标串是否匹配。
  char sstr[100]; memset(sstr,0,sizeof(sstr));
  char tstr[100]; memset(tstr,0,sizeof(tstr));
  strcpy(sstr,"aabaabcbcabaabc"); 
  strcpy(tstr,"abaabc");
  printf("index_bf =%d\n", index_bf(sstr,tstr));
  printf("index_kmp=%d\n",index_kmp(sstr,tstr));
  //////////////////////////////////
}
