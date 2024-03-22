/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    values := []int{}
    for head != nil {
        values = append(values, head.Val)
        head = head.Next
    }
    size := len(values)
    //fmt.Println("Values : ", values, "\nSize : ", size)
    for i:=0; i<size/2; i++ {
        if values[i] != values[size-1-i] {
            return false
        }
    }
    return true
}
