/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    dummyHead := &ListNode{0, head}
    prev := dummyHead

    for prev.Next != nil && prev.Next.Next != nil {
        first := prev.Next
        second := first.Next

        first.Next = second.Next
        second.Next = first
        prev.Next = second

        prev = first
    }

    return dummyHead.Next
}