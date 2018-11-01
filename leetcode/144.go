/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorderTraversal(root *TreeNode) []int {
    path := make([]int, 0)

    if root != nil {
        path = append(path, root.Val)
        path = append(path, preorderTraversal(root.Left)...)
        path = append(path, preorderTraversal(root.Right)...)
    }

    return path
}
