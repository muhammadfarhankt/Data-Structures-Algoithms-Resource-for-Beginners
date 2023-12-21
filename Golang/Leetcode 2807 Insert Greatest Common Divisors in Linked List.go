/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    curr := head
    for curr != nil && curr.Next != nil {
        num := gcd(curr.Val, curr.Next.Val)
        newNode := &ListNode{
            Val: num,
            Next: curr.Next,
        }
        curr.Next = newNode
        curr = curr.Next.Next
    }
    return head
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	tmp := a
	a = b
	b = tmp % a
	return gcd(a, b)
}
