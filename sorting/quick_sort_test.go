package sorting

import (
	"reflect"
	"testing"
)

//func Test_placePivot(t *testing.T) {
//	type args struct {
//		list     []int
//		pivotIdx int
//	}
//	tests := []struct {
//		name     string
//		args     args
//		want     int
//		wantList []int
//	}{
//		{
//			name: "T1",
//			args: args{
//				list:     []int{1, 2, 3, 4, 8, 7},
//				pivotIdx: 5,
//			},
//			want:     4,
//			wantList: []int{1, 2, 3, 4, 7, 8},
//		},
//
//		{
//			name: "T2",
//			args: args{
//				list:     []int{1, 2, 3, 4, 5, 6},
//				pivotIdx: 5,
//			},
//			want:     5,
//			wantList: []int{1, 2, 3, 4, 5, 6},
//		},
//
//		{
//			name: "T3",
//			args: args{
//				list:     []int{1, 9, 3, 13, 5, 7, 11},
//				pivotIdx: 6,
//			},
//			want:     5,
//			wantList: []int{1, 9, 3, 7, 5, 11, 13},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := placePivot(tt.args.list, tt.args.pivotIdx); got != tt.want {
//				t.Errorf("placePivot() = %v, want %v", got, tt.want)
//			}
//			if !reflect.DeepEqual(tt.args.list, tt.wantList) {
//				t.Errorf("placePivot() = %v, want %v", tt.args.list, tt.wantList)
//			}
//		})
//	}
//}

func Test_swap(t *testing.T) {
	type args struct {
		list []int
		fIdx int
		Lidx int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swap(tt.args.list, tt.args.fIdx, tt.args.Lidx)
		})
	}
}

func Test_quickSortRec(t *testing.T) {
	type args struct {
		list  []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "T1",
			args: args{
				list:  []int{1, 9, 3, 13, 5, 7, 11},
				start: 0,
				end:   6,
			},
			want: []int{1, 3, 5, 7, 9, 11, 13},
		},

		{
			name: "T2",
			args: args{
				list:  []int{6, 5, 4, 3, 2, 1},
				start: 0,
				end:   5,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},

		{
			name: "T3",
			args: args{
				list:  []int{1, 9, 3, 13, 5, 7, 11},
				start: 0,
				end:   6,
			},
			want: []int{1, 3, 5, 7, 9, 11, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSortRec(tt.args.list, tt.args.start, tt.args.end)
			if !reflect.DeepEqual(tt.args.list, tt.want) {
				t.Errorf("placePivot() = %v, want %v", tt.args.list, tt.want)
			}

		})
	}
}

func Test_quickSort(t *testing.T) {
	type args struct {
		list  []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "T1",
			args: args{
				list:  []int{1, 9, 3, 13, 5, 7, 11},
				start: 0,
				end:   6,
			},
			want: []int{1, 3, 5, 7, 9, 11, 13},
		},

		{
			name: "T2",
			args: args{
				list:  []int{6, 5, 4, 3, 2, 1},
				start: 0,
				end:   5,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},

		{
			name: "T3",
			args: args{
				list:  []int{1, 9, 3, 13, 5, 7, 11},
				start: 0,
				end:   6,
			},
			want: []int{1, 3, 5, 7, 9, 11, 13},
		},

		{
			name: "T3",
			args: args{
				list:  []int{5, 1, 1, 2, 0, 0},
				start: 0,
				end:   5,
			},
			want: []int{0, 0, 1, 1, 2, 5},
		},

		{
			name: "T4",
			args: args{
				list:  []int{5, 2, 2, 1},
				start: 0,
				end:   3,
			},
			want: []int{1, 2, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortArray(tt.args.list)
			if !reflect.DeepEqual(tt.args.list, tt.want) {
				t.Errorf("placePivot() = %v, want %v", tt.args.list, tt.want)
			}

		})
	}
}
