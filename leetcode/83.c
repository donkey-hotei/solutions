#include<stdlib.h>

struct Node {
    int val;
    struct Node * next;
}

struct Node * dedupe(struct Node * head) {
    if (head == NULL || head->next == NULL) return head;

    struct Node * curr, next;

    curr = head;
    next = NULL;

    while(curr->next) {
	if (curr->val == curr->next->val) {
	    next = curr->next->next;
	    free(curr->next);
	    curr->next = next;
	} else {
	    curr = curr->next;
	}
    }

    return head;
}
