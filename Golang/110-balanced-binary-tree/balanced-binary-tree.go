/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    // left := height(root.Left)
    // right := height(root.Right)
    if abs(height(root.Left), height(root.Right)) > 1 {
        return false
    }
    return isBalanced(root.Left) && isBalanced(root.Right)
}
func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(height(root.Left), height(root.Right)) + 1
}
func abs(x, y int) int {
    if x >= y {
        return x - y
    }
    return y - x
}
func max(x, y int) int {
    if x >= y {
        return x
    }
    return y
}