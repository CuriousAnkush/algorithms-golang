package algorithms

import "testing"

func Test_minMovesToSeat(t *testing.T) {
	type args struct {
		seats    []int
		students []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestCase 1",
			args: args{
				seats:    []int{2, 2, 6, 6},
				students: []int{1, 3, 2, 6},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMovesToSeat(tt.args.seats, tt.args.students); got != tt.want {
				t.Errorf("minMovesToSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
