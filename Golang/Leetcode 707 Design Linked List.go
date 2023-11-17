type Node struct {
    Val int
    next *Node
}
type MyLinkedList struct {
    head *Node
    length int
}
func Constructor() MyLinkedList {
    return MyLinkedList{head : nil, length : 0}
}
func (this *MyLinkedList) Get(index int) int {
        if index < 0 || index >= this.length {
            return -1
        } else {
        curr := this.head
        for i := 0; i < index; i++ {
            curr = curr.next
        }
        return curr.Val
        }
}
func (this *MyLinkedList) AddAtHead(val int)  {
    this.head = &Node{Val: val, next: this.head}
    this.length++
}
func (this *MyLinkedList) AddAtTail(val int)  {
    if this.length == 0 {
        this.AddAtHead(val)
        return
    }
    curr := this.head
    for i := 0; i < this.length-1; i++ {
        curr = curr.next
    }
    curr.next = &Node{Val: val}
    this.length++
    return
}
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index < 0 || index > this.length {
            return
    }
    if index == 0 {
        this.AddAtHead(val)
        return
    } else if index == this.length {
        this.AddAtTail(val)
        return
    } else {
        curr := this.head
        for i := 0; i < index-1; i++ {
            curr = curr.next
        }
        curr.next = &Node{Val: val, next: curr.next}
        this.length++
        return
    }
}
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index >= this.length {
        return
    } else if index == 0 {
        this.head = this.head.next
        this.length--
        return
    } else {
        curr := this.head
        for i := 0; i < index-1; i++ {
            curr = curr.next
        }
        curr.next = curr.next.next
        this.length--
        return
    }
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
