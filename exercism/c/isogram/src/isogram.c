#include "isogram.h"
#include<string.h>
#include<stdlib.h>


bool is_uppercase(char c) {
    return 'A' <= c && c <= 'Z';
}

char is_lowercase(char c) {
    return 'a' <= c && c <= 'z';
}

char lower(char c) {
    if (is_uppercase(c))
        return c - 'A';
    else if (is_lowercase(c))
        return c - 'a';
    else
        return c;
}

bool is_isogram(const char phrase[]) {
    if (phrase == NULL) return false;

    int bit = 0;
    int len = strlen(phrase);

    for (int i = 0; i < len; i++) {
        char c = lower(phrase[i]);

        if (phrase[i] == c)
            continue;
        if ((bit & (1 << c)) > 0)
            return false;

        bit |= (1 << c);
    }

    return true;
}
