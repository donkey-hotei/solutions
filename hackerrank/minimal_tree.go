package main

import (
    "fmt"
    "math"
    "./tree"
)

/*
 * 4.2: Minimal Tree
 *
 * Given a sorted array of numbers, from it build a binary tree with minimal height.
*/

func minimalTree(arr []int) *tree.Tree {
    n := len(arr)
    t := &tree.Tree{}

    if n == 0 {
        return t
    }

    q := make([]*tree.Tree, 0)
    q = append(q, t)

    k := int(math.Floor(float64(n) / 2.0))

    // O(floor(n / 2)) where n is len(arr)
    for i := 0; i < k; i++ {
        l := 2 * i + 1
        r := 2 * i + 2

        if len(q)  == 0 {
            break
        }

        curr_t := q[0]
        q = q[1:]

        curr_t.Value = arr[i]

        if l < n {
            curr_t.Left = &tree.Tree{
                nil,
                arr[l],
                nil,
            }
            q = append(q, curr_t.Left)
        }

        if r < n {
            curr_t.Right = &tree.Tree{
                nil,
                arr[r],
                nil,
            }
            q = append(q, curr_t.Right)
        }
    }

    return t
}

func main() {
    /*
             1
           /   \
          /     \
         3        5
        / \      / \
       /   \    /   \
      8    13  21   42

    */
    a := []int{1, 3, 5, 8, 13, 21, 42}
    t := minimalTree(a)
    println(t.String())
}