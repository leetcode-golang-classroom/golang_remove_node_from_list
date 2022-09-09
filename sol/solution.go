package sol

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		head = head.Next
		return head
	}
	fast := head
	var slow *ListNode
	count := 1
	for fast.Next != nil {
		if count == n {
			slow = head
		}
		if count > n {
			slow = slow.Next
		}
		fast = fast.Next
		count++
	}
	if slow == nil {
		head = head.Next
	} else {
		slow.Next = slow.Next.Next
	}
	return head
}
