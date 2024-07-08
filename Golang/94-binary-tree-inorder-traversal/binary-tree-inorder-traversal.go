/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    nums := []int{}
    if root == nil {
        return nums
    }
    nums = append(nums, inorderTraversal(root.Left)...)
    nums = append(nums, root.Val)
    nums = append(nums, inorderTraversal(root.Right)...)
    return nums
}