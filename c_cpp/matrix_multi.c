#include <stdio.h>
int main()
{
    int a[2][4], b[4][3], c[2][3];
    int i, j, k, sum;
    printf("Please input a 2×4 matrix：\n");
    for (i = 0; i < 2; i++)
        for (j = 0; j < 4; j++)
            scanf("%d", &a[i][j]);
    printf("Please input a 4×3 matrix;\n");
    for (i = 0; i < 4; i++)
        for (j = 0; j < 3; j++)
            scanf("%d", &b[i][j]);
    for (i = 0; i < 2; i++)
    {
        for (j = 0; j < 3; j++)
        {
            sum = 0;
            for (k = 0; k < 4; k++) 
            {
                sum += a[i][k] * b[k][j]; 
            }
            c[i][j] = sum; 
        }
    }
    printf("Get 2×3 matrix：\n");
    for (i = 0; i < 2; i++) 
    {
        for (j = 0; j < 3; j++)
            printf("%5d", c[i][j]);
        printf("\n");
    }
}
