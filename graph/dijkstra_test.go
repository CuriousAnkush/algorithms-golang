package graph

import "testing"

func Test_networkDelayTime(t *testing.T) {
	type args struct {
		times [][]int
		n     int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "T1",
			args: args{
				times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
				n:     4,
				k:     2,
			},
			want: 2,
		},
		{
			name: "T2",
			args: args{
				times: [][]int{{2, 5, 1}, {4, 5, 2}, {3, 4, 1}, {1, 2, 1}, {1, 3, 1}, {1, 4, 3}, {2, 4, 2}},
				n:     4,
				k:     1,
			},
			want: 2,
		},
		{
			name: "T3",
			args: args{
				times: [][]int{{1, 2, 1}},
				n:     2,
				k:     1,
			},
			want: 1,
		},
		{
			name: "T4",
			args: args{
				times: [][]int{{1, 2, 1}},
				n:     2,
				k:     2,
			},
			want: -1,
		},
		{
			name: "T5",
			args: args{
				times: [][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 1}},
				n:     2,
				k:     2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := networkDelayTime(tt.args.times, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("networkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
