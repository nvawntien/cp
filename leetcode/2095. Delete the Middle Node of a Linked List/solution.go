/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    n := 0
    dummy := head

    for dummy != nil {
        n++
        dummy = dummy.Next
    }

    if n == 1 {
        return nil
    }
    
    dummy = head

    for i := 0; i < n/2-1; i++ {
        dummy = dummy.Next
    }

    dummy.Next = dummy.Next.Next

    return head
}