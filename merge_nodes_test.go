package algorithms

import (
	"reflect"
	"testing"
)

func Test_mergeNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "T1",
			args: args{head: createList([]int{0, 1, 0, 3, 0, 2, 2, 0})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createList(list []int) *ListNode {
	head := &ListNode{
		Val:  list[0],
		Next: nil,
	}
	current := head

	for i := range list[1:] {
		current.Next = &ListNode{
			Val:  list[1:][i],
			Next: nil,
		}
		current = current.Next
	}
	return head
}
