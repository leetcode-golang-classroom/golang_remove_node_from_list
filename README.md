# golang_remove_node_from_list

Given the `head`
 of a linked list, remove the `nth`
 node from the end of the list and return its head.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg](https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg)

```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

```

**Example 2:**

```
Input: head = [1], n = 1
Output: []

```

**Example 3:**

```
Input: head = [1,2], n = 1
Output: [1]

```

**Constraints:**

- The number of nodes in the list is `sz`.
- `1 <= sz <= 30`
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

## 解析

題目給定一個單向串列 head 與一個數字 n

要求要移除掉 head 中倒數第 n 個結點

並回傳更新後的 head

單向鏈結串列沒辦法紀錄走過的結點，因此必須要設計一個結構來記住

直覺來想有 3 種作法

作法 1:

  先 loop 一此找出總共有幾個 —> 算出 length of list = Len

  然後在從頭走到 第 Len - n 個

  移除掉這個結點再回傳 head

  Time Complexity: O(n)

  Space Complexity: O(1)

作法2:

   用一個指標陣列 arr 紀錄每次走過的結點

   要移除的結點就是 arr[len(arr)-n]

   if len(arr) - n == 0 { // 移出的原本的 head

        head = head.Next

   } 

   if len(arr) - n > 0 {

       arr[len(arr)-n-1].Next = arr[len(arr)-n].Next

   } 

  Time Complexity: O(n)

  Space Complexity: O(n)

作法3:

  用兩個指標 slow, fast

  還有一個 counter 用來紀錄 fast 走過幾個 node

  fast 先走，slow 等到 fast 走了 n 個點後再開始每次走一步

  當 fast 走到底時， slow 就是要移除的點

時間複雜度是 O(n)

空間複雜度是 O(1)
## 程式碼

```go
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
```

## 困難點

1. 單向鏈結串列沒辦法紀錄走過的結點，因此必須要設計一個結構來記住
2. 題目要求是要從最後往前數，所以必須計算的方式

## Solve Point

- [x]  Understand what problem this question would like to solve
- [x]  Analysis Complexity