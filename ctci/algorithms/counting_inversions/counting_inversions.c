#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

/*
 * count(array A, length n)
 *   if n = 1 return 0
 *   else
 *     x = count(1st half of A, n / 2)
 *     y = count(2nd half of A, n / 2)
 *     z = count_split_inv(A, n)
 */

/*
 * Merges two sub-arrays of A, arr[p..q] and arr[q + 1..r].
 */
void merge(int A[], int p, int q, int r) {
    int i, j, k;
    int n1 = q - p + 1;
    int n2 = r - q;

    int L[n1 + 1], R[n2 + 1];

    // copy over left & right subarrays
    for (i = 0; i <= n1; i++)
        L[i] = A[p + i - 1];

    for (j = 0; j <= n2; j++)
        R[i] = A[q + j];

    i, j = 0;
    for (k = p; k <= r; k++) {
        if (L[i] <= R[j]) {
            A[k] = L[i];
            i += 1;
        } else {
            A[k] = R[j];
            j += 1;
        }
    }
}

/*
 * Simple implementation of merge sort
 */
void merge_sort(int A[], int p, int r) {
    if (p < r) {
        int q = (int)floor((double)((p + r) / 2));
        merge_sort(A, p, q);
        merge_sort(A, q + 1, r);
        merge(A, p, q, r);
    }
}

/*
 * A modification of the merge sort algorithm used to count the number
 * of inversions made whilist sorting the given array.
 */
long int count_inversions(int * arr) {
}

int main() {
    int t, n, arr_i;
    long int result;
    int * arr;
    scanf("%i", &t);

    for(int a0 = 0; a0 < t; a0++){
        scanf("%i", &n);
        arr = malloc(sizeof(int) * n);

        for (arr_i = 0; arr_i < n; arr_i++) {
           scanf("%i",&arr[arr_i]);
        }

        result = count_inversions(arr);

        printf("%ld\n", result);
    }

    return 0;
}

