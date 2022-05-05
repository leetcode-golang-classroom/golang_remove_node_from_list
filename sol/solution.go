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
	cur := head
	arr := []*ListNode{}
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	aLen := len(arr)
	target := arr[aLen-n]
	if aLen-n > 0 {
		arr[aLen-n-1].Next = target.Next
	}
	if aLen-n == 0 {
		head = arr[aLen-n].Next
	}
	return head
}
