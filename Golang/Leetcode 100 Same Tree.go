/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func inOrder(n *TreeNode) []int {
//     nums := []int{}
//     if n == nil {
//         return nums
//     }
//     nums = append(nums, inOrder(n.Left)...)
//     nums = append(nums, n.Val)
//     nums = append(nums, inOrder(n.Right)...)
//     return nums
// }
func isSameTree(p *TreeNode, q *TreeNode) bool {
    // pArray := inOrder(p)
    // qArray := inOrder(q)
    // if (reflect.DeepEqual(pArray,qArray)) {
    //     return true
    // } else {
    //     return false
    // }
    if p == nil && q == nil {
        return true
    } else if p == nil || q == nil {
        return false
    }
    if p.Val != q.Val {
        return false
    }
    return ( isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right) )
    return true
}
