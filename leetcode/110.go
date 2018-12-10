package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func abs(n int) int {
    if n < 0 {
        return -n
    } else if n > 0 {
        return n
    } else {
        return 0
    }
}

func depth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    l_depth := depth(root.Left)
    if l_depth == -1 {
        return l_depth
    }

    r_depth := depth(root.Right)
    if r_depth == -1 {
        return r_depth
    }

    if abs(l_depth - r_depth) > 1 {
        return -1
    }

    return 1 + max(l_depth, r_depth)
}

func isBalanced(root *TreeNode) bool {
    return depth(root) != -1
}

func main() {
    tree := TreeNode {
        1,
        &TreeNode {
            2,
            &TreeNode {
                3,
                &TreeNode {
                    4, nil, nil,
                },
                &TreeNode {
                    4, nil, nil,
                },
            },
            &TreeNode {
                3, nil, nil,
            },
        },
        &TreeNode {
            2, nil, nil,
        },
    }

    fmt.Printf("%v", isBalanced(&tree))
}
