package tree

import (
    "fmt"
    "math/rand"
)

/*
 * Tree is a binary tree with integer values.
 */
type Tree struct {
    Left *Tree
    Value int
    Right *Tree
}

/*
 * Returns a new binary tree holding values k, 2k, ..., 10k.
 */
func New(k int) *Tree {
    var t *Tree

    for _, v := range rand.Perm(10) {
        t = insert(t, (1 + v) * k)
    }

    return t
}

/*
 * Inserts a new value into a Binary Tree.
*/

func (t *Tree) Insert(v int) * Tree {
    return insert(t, v)
}

func insert(t *Tree, v int) *Tree {
    if t == nil {
        return &Tree{nil, v, nil}
    }

    if v < t.Value {
        t.Left = insert(t.Left, v)
    } else {
        t.Right = insert(t.Right, v)
    }

    return t
}

func (t *Tree) String() string {
    if t == nil {
        return "()"
    }

    s := ""
    if t.Left != nil {
        s += t.Left.String() + " "
    }
    s += fmt.Sprint(t.Value)
    if t.Right != nil {
        s += " " + t.Right.String()
    }

    return "(" + s + ")"
}
