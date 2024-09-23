package heap

import "testing"

func TestKthLargest_Add(t *testing.T) {
	type fields struct {
		minHeap MinHeap
		k       int
	}
	type args struct {
		val []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "T1",
			fields: fields{
				minHeap: []*HeapNode{},
				k:       0,
			},
			args: args{val: []int{3, 5, 10, 9, 4}},
			want: []int{4, 5, 5, 8, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(3, []int{4, 5, 8, 2})
			for idx, i := range tt.args.val {
				if got := this.Add(i); got != tt.want[idx] {
					t.Errorf("Add() = %v, want %v", got, tt.want)
				}
			}

		})
	}
}
