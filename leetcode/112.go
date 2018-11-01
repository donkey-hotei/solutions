package main

/*
 * Returns true iff there is a root to leaf path that sums to a given sum.
*/
func hasPathSum(t *TreeNode, sum int) bool {
    if t == nil {
        return false
    }
    if (sum - t.Val == 0) && t.Left == nil && t.Right == nil {
        return true
    }
    return hasPathSum(t.Left, sum - node.Val) || hasPathSum(t.Right, sum - node.Val)
}