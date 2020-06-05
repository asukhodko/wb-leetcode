package _151_reverse_words_in_a_string

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{"the sky is blue"}, "blue is sky the"},
		{"Example 2", args{"  hello world!  "}, "world! hello"},   // Your reversed string should not contain leading or trailing spaces.
		{"Example 3", args{"a good   example"}, "example good a"}, // You need to reduce multiple spaces between two words to a single space in the reversed string.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
