package dynamic_programming

import "testing"

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "T1",
			args: args{matrix: [][]byte{
				{1, 0, 1, 0, 0},
				{1, 0, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 0, 0, 1, 0},
			}},
			want: 4,
		},
		{
			name: "T2",
			args: args{matrix: [][]byte{
				{0, 1},
				{1, 0},
			}},
			want: 1,
		},
		{
			name: "T3",
			args: args{matrix: [][]byte{
				{0},
			}},
			want: 0,
		},
		{
			name: "T4",
			args: args{matrix: [][]byte{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
