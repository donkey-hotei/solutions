#include<stdio.h>
#include<stdlib.h>
#include<assert.h>

#define true  1
#define false 0

int is_palindrome(int n) {
    int r = n;
    int m = 0;

    if (n < 0) {
        return false;
    }

    while (r > 0) {
        m *= 10;
        m += r % 10;
        r /= 10;
    }

    return m == n;
}

int main() {
    assert(is_palindrome(121) == true);
    assert(is_palindrome(1221) == true);
    assert(is_palindrome(123) == false);
}