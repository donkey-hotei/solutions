#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#include<strings.h>

int lengthOfLastWord(char* s) {
    int len = 0, tail = strlen(s) -1;

    while (tail >= 0 && s[tail] == ' ')
        tail--;
    while (tail >= 0 && s[tail] != ' ') {
        len++;
        tail--;
    }

    return len;
}

int main() {
    char * str = " a ";

    printf("Length of last word in %s is %d", str, lengthOfLastWord(str));
}
