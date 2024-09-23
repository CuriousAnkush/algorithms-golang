package algorithms

import (
	"testing"
)

func TestTopVotedCandidate_Q(t *testing.T) {
	type fields struct {
		persons []int
		times   []int
	}
	type args struct {
		t int
	}
	tests := []struct {
		name   string
		fields fields
		args   []int
		want   []int
	}{
		{
			name: "T1",
			fields: fields{
				persons: []int{0, 1, 1, 0, 0, 1, 0},
				times:   []int{0, 5, 10, 15, 20, 25, 30},
			},
			args: []int{3, 12, 25, 15, 24, 8},
			want: []int{0, 1, 1, 0, 0, 1},
		},
		{
			name: "T2",
			fields: fields{
				persons: []int{0, 1, 2, 2, 1},
				times:   []int{20, 28, 29, 54, 56},
			},
			args: []int{28, 53, 57, 29, 29, 28, 30, 20, 56, 55},
			want: []int{1, 2, 1, 2, 2, 1, 2, 0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.fields.persons, tt.fields.times)
			for idx := range tt.args {
				got := this.Q(tt.args[idx])
				want := tt.want[idx]
				if got != want {
					t.Errorf("Q() = %v, want %v", got, want)
				}
			}
		})
	}
}
