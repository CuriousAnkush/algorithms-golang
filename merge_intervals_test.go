package algorithms

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "T1 Already Sorted",
			args: args{intervals: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			}},
			want: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},

		{
			name: "T1 Not Sorted",
			args: args{intervals: [][]int{
				{8, 10},
				{2, 6},
				{1, 3},
				{15, 18},
			}},
			want: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},

		{
			name: "T1 One element",
			args: args{intervals: [][]int{
				{8, 10},
			}},
			want: [][]int{
				{8, 10},
			},
		},

		{
			name: "T1 One element",
			args: args{intervals: [][]int{
				{4, 5},
				{1, 5},
			}},
			want: [][]int{
				{1, 5},
			},
		},

		{
			name: "All non overlapping",
			args: args{intervals: [][]int{
				{9, 11},
				{1, 8},
				{90, 100},
				{50, 60},
			}},
			want: [][]int{
				{1, 8},
				{9, 11},
				{50, 60},
				{90, 100},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
