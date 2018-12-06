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

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func depth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return 1 + max(depth(root.Left), depth(root.Right))
}

func isBalanced(root *TreeNode) bool {
    if root == nil || (root.Left == nil && root.Right == nil) {
        return true
    }

    min_depth := min(depth(root.Left), depth(root.Right))
    max_depth := max(depth(root.Left), depth(root.Right))

    fmt.Printf("min_depth=%d, max_depth=%d\n", min_depth, max_depth)

    return max_depth - min_depth <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
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
