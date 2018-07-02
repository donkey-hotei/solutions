package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) {
    if root == nil {
        return 0
    }

    l_depth := maxDepth(root.Left)
    r_depth := maxDepth(root.Right)
    return max(l_depth, r_depth) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}