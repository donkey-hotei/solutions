#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

/*
 * Finds integer in array s.t. it is the only
 * integer w/o a duplicate.
 */
int lonely_integer(int * a, int n) {
    int r, i = 0;

    for (i = 0; i <= n; i++) {
        r ^= a[i];
    }
    return r;
}

int main(){
    int n;
    scanf("%d",&n);

    if (n < 1)
        return 0;

    int *a = malloc(sizeof(int) * n);

    for(int a_i = 0; a_i < n; a_i++)
        scanf("%d", &a[a_i]);

    printf("%d\n", lonely_integer(a, n));

    return 0;
}

