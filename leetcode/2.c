#include<stdio.h>
#include<stdlib.h>

/*
 *
 *
 */
struct list {
    int value;
    struct list * next;
};

/*
 * Appends a node with given value to list.
 */
struct list* append(struct list* l, int value) {
    if (l == NULL) return l;
    l->next = malloc(sizeof(struct list));
    l->next->value = value;
}

/*
 * Print a list starting from the head.
 */
void print_list(struct list* l) {
    struct list* curr = l;
    while (curr->next != NULL) {
        printf("%d -> ", curr->value);
        curr = curr->next;
    }
    printf("%d\n", curr->value);
}

/*
 * Add two numbers represented as linked-lists with the
 * least-significant digit living at the head (reversed):
 *
 *   1 -> 1 -> 9 = 911
 * + 0 -> 2 -> 4 = 420
 * -------------------
 *        
*/
struct list* add(struct list* l1, struct list* l2) {
    struct list* result = malloc(sizeof(struct list));
    int value = l1->value + l2->value;
    int carry = 0;

    if (value >= 10) {
        value -= 10;
        carry = 1;
    }

    result->value = value;
    struct list* head = result; // save pointer for later

    while (l1->next != NULL && l2->next != NULL) {
        l1 = l1->next;
        l2 = l2->next;

        value = l1->value + l2->value;

        if (carry > 1) {
            value += carry;
            carry = 0;
        }

        if (value >= 10) {
            value -= 10;
            carry = 1;
        }

        result = append(result, value);
    }

    return head;
}

int main() {
    struct list * l1 = malloc(sizeof(struct list));
    struct list * l2 = malloc(sizeof(struct list));

    l1->value = 2;
    l2->value = 5;

    append(l1, 1);
    append(l1->next, 3);
    append(l2, 4);
    append(l2->next, 8);

    print_list(l1);
    print_list(l2);

    struct list* result = add(l1, l2);
    print_list(result);
}