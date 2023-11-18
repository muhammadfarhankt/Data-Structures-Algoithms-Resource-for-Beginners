/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    length := 0
    curr := head
    for curr != nil {
        length++
        curr = curr.Next
    }
    length = length/2
    curr = head
    for i:=0; i<length; i++ {
        curr = curr.Next
    }
    return curr
}
