package algorithms

import "testing"

func Test_replaceWords(t *testing.T) {
	type args struct {
		dictionary []string
		sentence   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				dictionary: []string{"cat", "bat", "rat"},
				sentence:   "the cattle was rattled by the battery",
			},
			want: "the cat was rat by the bat",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := replaceWords(tt.args.dictionary, tt.args.sentence)
			if got != tt.want {
				t.Errorf("got %v, want: %v", got, tt.want)
			}
		})
	}
}
