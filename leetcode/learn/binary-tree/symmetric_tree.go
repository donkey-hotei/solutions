package main

type TreeNode struct {
    Right *TreeNode
    Val int
    Left *TreeNode
}

func isSymmetric(root *TreeNode) {
    if root == nil {
        return true
    }

    return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) {
    if left == nil || right == nil {
        return left == nil && right == nil
    }

    return left.Val == right.Val &&
           isMirror(left.Right, right.Left) &&
           isMirror(left.Left, right.Right)
}