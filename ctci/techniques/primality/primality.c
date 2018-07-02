#include<math.h>
#include<stdio.h>
#include<string.h>
#include<stdlib.h>
#include<assert.h>
#include<limits.h>
#include<stdbool.h>


/*
 *
 */
bool sieve(int n) {
    bool table[n] = true;

    int i, j;
    for (i = 2; i < n + 1; i++) {
        for (j = 2; i * j < n; j++) {
            if (table[i * j])
                table[i * j] = false;
        }
    }
}

int main() {
    int p, i, n;
    scanf("%d", &p);
    for (i = 0; i < p; i++) {
        scanf("%d", &n);
    }
}