/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    tree1 := []int{}
    tree2 := []int{}
    treeTraversal(root1, &tree1)
    treeTraversal(root2, &tree2)
    if len(tree1) != len(tree2) {
        return false
    }
    for i:=0; i<len(tree1); i++ {
        if tree1[i] != tree2[i] {
            return false
        }
    }
    return true
}
func treeTraversal (root *TreeNode, leafArray *[]int) {
    if root != nil {
        if root.Left == nil && root.Right == nil {
            *leafArray = append(*leafArray, root.Val)
        }
        treeTraversal(root.Left, leafArray)
        treeTraversal(root.Right, leafArray)
    }
}
