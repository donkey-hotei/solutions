package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    if root.Left == nil && root.Right == nil {
        return 1
    }

    min_depth := int(^uint(0) >> 1)  // max signed int

    if root.Left != nil {
        min_depth = min(minDepth(root.Left), min_depth)
    }

    if root.Right != nil {
        min_depth = min(minDepth(root.Right), min_depth)
    }


    return min_depth + 1
}

func main() {
    tree := TreeNode {
        3,
        &TreeNode {
            9,
            nil,
            nil,
        },
        &TreeNode {
            20,
            &TreeNode {
                15,
                nil,
                nil,
            },
            &TreeNode {
                7,
                nil,
                nil,
            },
        },
    }
    fmt.Printf("%v", minDepth(&tree))
}
