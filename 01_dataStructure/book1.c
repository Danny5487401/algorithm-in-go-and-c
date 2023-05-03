#include <stdio.h>
 
int main(int argc, char *argv[])
{
    int a = 0, i = 0, j = 0;
     
    printf("请输入一个数：\n");
    scanf("%d", &a);
    if (a <= 1)
    {
        printf("素数要大于1\n");
        return 0;
    }
     
    for (j = a; j > 1; j--)  /*每次从本身开始每次发现不是素数就j--取离他最近的值*/
    {
        for (i = 2; i < j; i++) /*查找是否有除了1和他本身之外的因子*/
        {
            /*确定为素数则j--再选最近的一个*/
            if (0 == j%i)
            {
                break;
            }
        }
        if (i == j)
        {
            printf("不大于%d的最大素数是%d\n", a, j);
            break;
        }
    }
     
    return 0;
}
