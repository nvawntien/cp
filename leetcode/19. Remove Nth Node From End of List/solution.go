/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummyHead := &ListNode{0, head}
    first := dummyHead
    second := dummyHead

    for i := 0; i <= n; i++ {
        first = first.Next
    }

    for first != nil {
        first = first.Next
        second = second.Next
    }

    second.Next = second.Next.Next
    
    return dummyHead.Next
}