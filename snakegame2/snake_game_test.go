package snakegame2

import (
	"testing"
)

func TestSnakeGame_Move(t *testing.T) {
	type args struct {
		width  int
		height int
		food   [][]int
	}

	tests := []struct {
		name string
		args args
		want []map[string]int
	}{
		{
			name: "TestCase 1",
			args: args{
				width:  3,
				height: 2,
				food:   [][]int{{1, 2}, {0, 1}},
			},
			want: []map[string]int{
				{"R": 0},
				{"D": 0},
				{"R": 1},
				{"U": 1},
				{"L": 2},
				{"U": -1},
			},
		},
		{
			name: "TestCase 2",
			args: args{
				width:  3,
				height: 3,
				food:   [][]int{{2, 0}, {0, 0}, {0, 2}, {0, 1}, {2, 2}, {0, 1}},
			},
			want: []map[string]int{
				{"D": 0},
				{"D": 1},
				{"R": 1},
				{"U": 1},
				{"U": 1},
				{"L": 2},
				{"D": 2},
				{"R": 2},
				{"R": 2},
				{"U": 3},
				{"L": 4},
				{"L": 4},
				{"D": 4},
				{"R": 4},
				{"U": -1},
			},
		},
		{
			name: "TestCase 3",
			args: args{
				width:  3,
				height: 3,
				food:   [][]int{{2, 0}, {0, 0}, {0, 2}, {2, 2}},
			},
			want: []map[string]int{
				{"D": 0},
				{"D": 1},
				{"R": 1},
				{"U": 1},
				{"U": 1},
				{"L": 2},
				{"D": 2},
				{"R": 2},
				{"R": 2},
				{"U": 3},
				{"L": 3},
				{"D": 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Constructor(tt.args.width, tt.args.height, tt.args.food)
			for _, dirScoreMap := range tt.want {
				for dir, score := range dirScoreMap {
					var got int
					if score == 3 {
						got = g.Move(dir)
					} else {
						got = g.Move(dir)
					}
					if got != score {
						t.Errorf("CurrentScore(): %d, expected: %d", got, score)
					}
				}
			}
		})
	}
}
