/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    path := make([]int, 0)
    if root != nil {
        path = append(path, postorderTraversal(root.Left)...)
        path = append(path, postorderTraversal(root.Right)...)
        path = append(path, root.Val)
    }
    return path
}

