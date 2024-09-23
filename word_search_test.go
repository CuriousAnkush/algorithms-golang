package algorithms

//
//import (
//	"algorithms/blind75"
//	"testing"
//)
//
//func Test_exist(t *testing.T) {
//	type args struct {
//		board [][]string
//		word  []string
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		{
//			name: "TestCase 1",
//			args: args{
//				board: [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}},
//				word:  []string{"ABCCED"},
//			},
//			want: true,
//		},
//		{
//			name: "TestCase 2",
//			args: args{
//				board: [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}},
//				word:  []string{"SEE"},
//			},
//			want: true,
//		},
//		{
//			name: "TestCase 3",
//			args: args{
//				board: [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}},
//				word:  []string{"ABCB"},
//			},
//			want: false,
//		},
//		{
//			name: "TestCase 3",
//			args: args{
//				board: [][]string{{"A", "B", "C", "E"}, {"S", "F", "E", "S"}, {"A", "D", "E", "E"}},
//				word:  []string{"ABCESEEEFS"},
//			},
//			want: true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := blind75.FindWords(tt.args.board, tt.args.word); got != tt.want {
//				t.Errorf("exist() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
