/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    dummy := head
    link := []int{}
    n := 0
    
    for dummy != nil {
        n++
        link = append(link, dummy.Val)
        dummy = dummy.Next
    }

    ans := 0

    for i := 0; i < n/2; i++ {
        ans = max(ans, link[i] + link[n-i-1])
    }

    return ans
}