package snakegame2

import (
	"container/list"
	"reflect"
	"testing"
)

func TestNewSnake(t *testing.T) {
	type args struct {
		board *Board
	}
	tests := []struct {
		name string
		args args
		want *Snake
	}{
		{
			name: "Constructor Test",
			args: args{board: NewBoard(3, 4, [][]int{})},
			want: &Snake{
				body: list.New(),
				board: &Board{
					width:  3,
					height: 4,
					food:   [][]int{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSnake(tt.args.board); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnake_Move(t *testing.T) {
	type fields struct {
		body  *list.List
		board *Board
	}
	type args struct {
		direction Direction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Snake{
				body:  tt.fields.body,
				board: tt.fields.board,
			}
			s.Move(tt.args.direction)
		})
	}
}
