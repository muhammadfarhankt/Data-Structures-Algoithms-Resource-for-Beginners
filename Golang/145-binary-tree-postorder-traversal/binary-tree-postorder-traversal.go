/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    nums := []int{}
    if root == nil {
        return nums
    }
    nums = append(nums, postorderTraversal(root.Left)...)
    nums = append(nums, postorderTraversal(root.Right)...)
    nums = append(nums, root.Val)
    return nums
}