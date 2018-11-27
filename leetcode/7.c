#include<stdio.h>
#include<stdlib.h>
#include<limits.h>

int reverse_number(int n) {
    int m = 0;
    int t = 0;;

    while (n != 0) {
        t = n % 10;

        if (m > INT_MAX / 10 || (m == INT_MAX / 10 && t > 7))
            return 0;

        if (m < INT_MAX / 10 || (m == INT_MAX / 10 && t < -8))
            return 0;

        m *= 10;
        m += t;
        n /= 10;
    }

    return m;
}