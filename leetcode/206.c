#include<stdio.h>
#include<stdlib.h>

struct ListNode {
    int val;
    struct ListNode * next;
};

struct ListNode * reverse(struct ListNode * head) {
    struct ListNode * curr, *prev, *next;

    curr = head;
    prev = NULL;
    next = NULL;

    while (curr) {
        next = curr->next;
        curr->next = prev;

        prev = curr;
        curr = next;
    }

    return prev;
}

int main() { }