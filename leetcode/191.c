#include<stdio.h>
#include<stdlib.h>

int hamming_weight(int n) {
    unsigned int count = 0;

    while (n) {
        n = n & (n - 1);

        count += 1;
    }

    return count;
}