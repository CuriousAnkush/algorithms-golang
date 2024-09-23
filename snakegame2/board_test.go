package snakegame2

import (
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	type args struct {
		width  int
		height int
		food   [][]int
	}
	tests := []struct {
		name string
		args args
		want *Board
	}{
		{
			name: "Test Constructor",
			args: args{
				width:  3,
				height: 2,
				food:   [][]int{},
			},
			want: &Board{
				width:  3,
				height: 2,
				food:   [][]int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoard(tt.args.width, tt.args.height, tt.args.food); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_FoodAvailable(t *testing.T) {
	type fields struct {
		width      int
		height     int
		food       [][]int
		foodCursor int
	}
	type args struct {
		rowNu int
		colNu int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				width:      tt.fields.width,
				height:     tt.fields.height,
				food:       tt.fields.food,
				foodCursor: tt.fields.foodCursor,
			}
			if got := b.FoodAvailable(tt.args.rowNu, tt.args.colNu); got != tt.want {
				t.Errorf("FoodAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}
