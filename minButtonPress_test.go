package algorithms

import "testing"

func Test_countButtonPress(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "T1",
			args: args{word: "aabbccddeeffgghhiiiiii"},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countButtonPress(tt.args.word); got != tt.want {
				t.Errorf("countButtonPress() = %v, want %v", got, tt.want)
			}
		})
	}
}
