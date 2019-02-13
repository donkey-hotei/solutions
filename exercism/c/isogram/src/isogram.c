#include "isogram.h"
#include<ctype.h>
#include<string.h>
#include<stdlib.h>

char tobit(char c) {
   if (isupper(c))
       return c - 'A';
   else
       return c - 'a';
}

bool is_isogram(const char phrase[]) {
    if (phrase == NULL) return false;

    int bit = 0;
    int len = strlen(phrase);

    for (int i = 0; i < len; i++) {
        char c = phrase[i];

        if (!isalpha(c)) continue;

        c = tobit(c);
        if ((bit & (1 << c)) > 0)
            return false;

        bit |= (1 << c);
    }

    return true;
}
