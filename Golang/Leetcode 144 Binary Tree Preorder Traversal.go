/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    nums := []int{}
    if root == nil {
        return nums
    }
    nums = append(nums, root.Val)
    nums = append(nums, preorderTraversal(root.Left)...)
    nums = append(nums, preorderTraversal(root.Right)...)
    return nums
}
