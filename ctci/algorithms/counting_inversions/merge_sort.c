#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

/*
 * Claim: "split" inversions involving some element y of the second array
 *        C are precisely the numbers left in the first array B when y is
 *        copied to the output D.
 *
 * Proof: Let x be an element of the first array B:
 *   1.) If x is copied to the output D before y then x < y which means there's
 *       no inversion involving x & y.
 *   2.) If y is copied to output D before x then y < x which means there's x & y
 *       are a split inversion.
 */

long int merge_and_count_split_inversions(int * A, int n, int m) {
    int i, j, k;
    int * x = malloc(n * sizeof(int));
    long int inv_count = 0;  // number of split inversions

    // i is element of the first sub-array
    // j is element of the second sub-array
    for (i = 0, j = m, k = 0; k < n; k++) {
        if (j == n) {
            x[k] = A[i++];
        } else if (i == m) {
            x[k] = A[j++];
        } else if (A[j] < A[i]) {
            x[k] = A[j++];
        } else {
            inv_count += (m + 1) - i;
            x[k] = A[i++];
        }
    }

    for (i = 0; i < n; i++)
        A[i] = x[i];

    free(x);

    return inv_count;
}


int sort_and_count_inversions(int * A, int n) {
    if (n < 2) return 0;

    int m = n / 2;
    long int x, y, z;

    x = sort_and_count_inversions(A, m);
    y = sort_and_count_inversions(A + m, n - m);
    z = merge_and_count_split_inversions(A, n, m);

    return x + y + z;
}


int main() {
    int arr_one[6] = {1, 3, 5, 2, 4, 6};
    int arr_two[5] = {2, 1, 3, 1, 2};
    int inv_count_one, inv_count_two;
    inv_count_one = sort_and_count_inversions(arr_one, 6);
    inv_count_two = sort_and_count_inversions(arr_two, 5);

    printf("%d inversions found in array one.\n", inv_count_one);
    printf("%d inversions foudn in array two.\n", inv_count_two);
}
