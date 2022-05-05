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

直覺來想有兩種作法

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
    cur := head
    arr := []*ListNode{}
    for cur != nil {
        arr = append(arr, cur)
        cur = cur.Next
    }
    aLen := len(arr)
    target := arr[aLen - n]
    if aLen - n > 0 {
        arr[aLen - n - 1].Next = target.Next
    } 
    if aLen - n == 0 {
        head = arr[aLen-n].Next
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