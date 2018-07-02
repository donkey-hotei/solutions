package main

import (
    "fmt"
    "./tree"
)

/*
 * 4.10: Check Subtree
 *
 * T1 and T2 are two very large binary trees, with T1 much bigger than T2.
 * Create an algorithm to determine if T2 is a subtree of T1.
 *
 * A tree T2 is a subtree of T1 if there exists a node n in T1 s.t. the subtree
 * of n is identical to T2. That is, if you cut off the tree at node n, the two
 * trees would be identical.
*/

/* Checks if t2 is a subtree of t1. */
func isSubtree(t2, t1 *tree.Tree) bool {
    q := make([]*tree.Tree, 0)
    q = append(q, t1)

    for len(q) != 0 {
        curr_t := q[0]
        q = q[1:]

        if curr_t.Value == t2.Value {
            if isEqual(curr_t, t2) {
                return true
            }
        } else {
            q = append(q, curr_t.Left)
            q = append(q, curr_t.Right)
        }
    }

    return false
}


/* Returns true if t1 equals t2 */
func isEqual(t1, t2 *tree.Tree) bool {
    if t1 == nil {
        return t2 == nil
    }

    if t2 == nil {
        return t1 == nil
    }

    if t1.Value == t2.Value {
        l_equal := isEqual(t1.Left, t2.Left)
        r_equal := isEqual(t1.Right, t2.Right)

        return l_equal && r_equal
    }

    return false
}


func main() {
    t1 := &tree.Tree{
        &tree.Tree{
            &tree.Tree{
                nil,
                8,
                nil,
            },
            3,
            &tree.Tree{
                nil,
                13,
                nil,
            },
        },
        1,
        &tree.Tree{
            nil,
            5,
            nil,
        },
    }

    t2 := &tree.Tree{
        &tree.Tree{
            nil,
            8,
            nil,
        },
        3,
        &tree.Tree{
            nil,
            13,
            nil,
        },
    }

    if isSubtree(t2, t1) {
        fmt.Println("T2 is a subtree of T1.")
    } else {
        fmt.Println("T2 is not a subtree of T1")
    }
}