/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    depth := maxDepth(root)
    return sum(root, 1, depth)
}
func maxDepth(root *TreeNode) int {
    if root == nil { return 0 }
    left := maxDepth(root.Left)
    right := maxDepth(root.Right)
    if left > right { return left + 1 }
    return right + 1
}
func sum(root *TreeNode, current int, depth int) int {
    if root == nil { return 0 }
    if current == depth { return root.Val }
    left := sum(root.Left, current+1, depth)
    right := sum(root.Right, current+1, depth)
    return left + right
}
