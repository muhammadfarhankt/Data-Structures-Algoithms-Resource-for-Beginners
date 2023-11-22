/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    countA, countB := 0,0
    listA, listB := headA, headB
    for listA != nil {
        countA++
        listA = listA.Next
    }
    for listB != nil {
        countB++
        listB = listB.Next
    }
    listA, listB = headA, headB
    if countA > countB {
        for i:=0; i<(countA-countB); i++ {
            listA = listA.Next
        }
    } else {
        for i:=0; i<(countB-countA); i++ {
            listB = listB.Next
        }
    }
    for listA != nil && listB != nil {
        if listB == listA {
            return listB
        }
        listA = listA.Next
        listB = listB.Next
    }
    return nil
}
