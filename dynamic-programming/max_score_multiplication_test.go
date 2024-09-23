package dynamic_programming

import "testing"

func Test_key(t *testing.T) {
	type args struct {
		mIdx int
		lIdx int
		rIdx int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test1",
			args: args{
				mIdx: 0,
				lIdx: 1,
				rIdx: 2,
			},
			want: "m_0_l_1_r_2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := key(tt.args.mIdx, tt.args.lIdx, tt.args.rIdx); got != tt.want {
				t.Errorf("key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumScore(t *testing.T) {
	type args struct {
		nums        []int
		multipliers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "T1",
			args: args{
				nums:        []int{1, 2, 3},
				multipliers: []int{3, 2, 1},
			},
			want: 14,
		},
		{
			name: "T2",
			args: args{
				nums:        []int{-5, -3, -3, -2, 7, 1},
				multipliers: []int{-10, -5, 3, 4, 6},
			},
			want: 102,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.args.nums, tt.args.multipliers); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
