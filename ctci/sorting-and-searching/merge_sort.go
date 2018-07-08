package main

import "fmt"

/*
 * Top-down merge sort works by recursively splitting the array into subarrays
 * until the subarray is of length one, it then merges those sublists to produce
 * a sorted list.
 */
func merge_sort(a []int) []int {
    n := len(a)

    if n <= 1 {
        return a
    }

    // to keep track of inversions
    inversions := 0

    // divide the array into two halves
    l := a[:n / 2]
    r := a[n / 2:]

    // and proceed to recursively split
    l = merge_sort(l)
    r = merge_sort(r)

    // before finally merging
    return merge(l, r, inversions)
}


func merge(left, right []int, inversions int) []int {
    l_len := len(left)
    r_len := len(right)
    n := l_len + r_len

    /*

        There are at least three types of inversions one can
        detect when performing merge sort. They are as follows:


    */

    // O(n) in space for working array
    array := make([]int, n)

    l, r := 0, 0
    for k := 0; k < n; k++ {
        if l >= l_len {
            array[k] = right[r]
            r++
            continue
        } else if r >= r_len {
            array[k] = left[l]
            l++
            continue
        }

        if left[l] >= right[r] {
            array[k] = right[r]
            r++
            inversions++
        } else {
            array[k] = left[l]
            l++
            inversions++
        }
    }

    return array
}


func main() {
    a := []int{14, 33, 27, 10, 35, 19, 42, 44}
    a = merge_sort(a)
    fmt.Printf("%v\n", a)
}
