/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head.Next == nil {
        head = nil
        return head
    }
    dummyHead := &ListNode{0, head}
    Lptr := dummyHead
    Rptr := dummyHead
    for i:=0 ; i <= n; i++ {
        Rptr = Rptr.Next
    }
    for Rptr != nil {
        Rptr = Rptr.Next
        Lptr = Lptr.Next
    }
    Lptr.Next = Lptr.Next.Next
    return dummyHead.Next
}

//Method 2
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head.Next == nil {
        head = nil
        return head
    }
    Lptr := head
    Rptr := head
    for i:=0 ; i < n; i++ {
        Rptr = Rptr.Next
    }
    if Rptr == nil {
        return head.Next
    }
    for Rptr.Next != nil {
        Rptr = Rptr.Next
        Lptr = Lptr.Next
    }
    Lptr.Next = Lptr.Next.Next
    return head
}
