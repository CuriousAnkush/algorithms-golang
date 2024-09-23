package blind75

import "testing"

func Test_shortestPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "T1",
			input: "aacecaaa",
			want:  "aaacecaaa",
		},
		{
			name:  "T2",
			input: "abcd",
			want:  "dcbabcd",
		},
		{
			name:  "T3",
			input: "aabcaa",
			want:  "aacbaabcaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPalindrome(tt.input); got != tt.want {
				t.Errorf("shortestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
