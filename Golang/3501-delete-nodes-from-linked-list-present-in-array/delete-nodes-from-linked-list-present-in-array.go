/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
    numSet := make(map[int]bool)
    for _, num := range nums {
        numSet[num] = true
    }
    
    temp := &ListNode{Next: head}
    prev := temp
    curr := head
    
    for curr != nil {
        if _, exists := numSet[curr.Val]; exists {
            prev.Next = curr.Next
        } else {
            prev = curr
        }
        curr = curr.Next
    }
    
    return temp.Next
}
