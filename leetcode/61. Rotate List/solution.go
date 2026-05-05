/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {
        return head
    }

    dummy := head
    length := 1

    for dummy.Next != nil {
        length++
        dummy = dummy.Next
    }

    if k % length == 0 {
        return head
    }
    
    newHead := head
    newTail := head

    n := length - (k % length)

    for i := 0; i < n; i++ {
        if i < n-1 {
            newTail = newTail.Next
        }  
        newHead = newHead.Next
    }

    dummy.Next = head
    newTail.Next = nil

    return newHead
}