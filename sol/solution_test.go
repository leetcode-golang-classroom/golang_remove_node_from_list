package sol

import (
	"reflect"
	"testing"
)

func CreateList(array *[]int) *ListNode {
	arr := *array
	var head *ListNode
	var cur *ListNode
	for idx, val := range arr {
		if idx == 0 {
			head = &ListNode{Val: val}
			cur = head
		} else {
			cur.Next = &ListNode{Val: val}
			cur = cur.Next
		}
	}
	return head
}
func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "head = [1,2,3,4,5], n = 2",
			args: args{head: CreateList(&[]int{1, 2, 3, 4, 5}), n: 2},
			want: CreateList(&[]int{1, 2, 3, 5}),
		},
		{
			name: "head = [1], n = 1",
			args: args{head: CreateList(&[]int{1}), n: 1},
			want: CreateList(&[]int{}),
		},
		{
			name: "head = [1,2], n = 1",
			args: args{head: CreateList(&[]int{1, 2}), n: 1},
			want: CreateList(&[]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
