package algorithms

import "testing"

//func TestProfiles_Len(t *testing.T) {
//	tests := []struct {
//		name string
//		p    Profiles
//		want int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.p.Len(); got != tt.want {
//				t.Errorf("Len() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestProfiles_Less(t *testing.T) {
//	type args struct {
//		i int
//		j int
//	}
//	tests := []struct {
//		name string
//		p    Profiles
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
//				t.Errorf("Less() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestProfiles_Swap(t *testing.T) {
//	type args struct {
//		i int
//		j int
//	}
//	tests := []struct {
//		name string
//		p    Profiles
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.p.Swap(tt.args.i, tt.args.j)
//		})
//	}
//}

func Test_maxProfitAssignment(t *testing.T) {
	type args struct {
		difficulty []int
		profit     []int
		worker     []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "T1",
			args: args{
				difficulty: []int{85, 47, 57},
				profit:     []int{24, 66, 99},
				worker:     []int{40, 25, 25},
			},
			want: 0,
		},
		{
			name: "T2",
			args: args{
				difficulty: []int{5, 50, 92, 21, 24, 70, 17, 63, 30, 53},
				profit:     []int{68, 100, 3, 99, 56, 43, 26, 93, 55, 25},
				worker:     []int{96, 3, 55, 30, 11, 58, 68, 36, 26, 1},
			},
			want: 765,
		},
		{
			name: "T3",
			args: args{
				difficulty: []int{66, 1, 28, 73, 53, 35, 45, 60, 100, 44, 59, 94, 27, 88, 7, 18, 83, 18, 72, 63},
				profit:     []int{66, 20, 84, 81, 56, 40, 37, 82, 53, 45, 43, 96, 67, 27, 12, 54, 98, 19, 47, 77},
				worker:     []int{61, 33, 68, 38, 63, 45, 1, 10, 53, 23, 66, 70, 14, 51, 94, 18, 28, 78, 100, 16},
			},
			want: 1392,
		},
		{
			name: "T4",
			args: args{
				difficulty: []int{49, 49, 76, 88, 100},
				profit:     []int{5, 8, 75, 89, 94},
				worker:     []int{98, 72, 16, 27, 76},
			},
			want: 172,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitAssignment(tt.args.difficulty, tt.args.profit, tt.args.worker); got != tt.want {
				t.Errorf("maxProfitAssignment() = %v, want %v", got, tt.want)
			}
		})
	}
}
