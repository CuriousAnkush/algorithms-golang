package algorithms

import (
	"reflect"
	"testing"
)

func Test_commonChars(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "Testcase-1",
			args: []string{"bella", "label", "roller"},
			want: []string{"e", "l", "l"},
		},
		{
			name: "Testcase-2",
			args: []string{"cool", "lock", "cook"},
			want: []string{"c", "o"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonChars(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
