package _028_implement_strstr

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"hello", args{"hello", "ll"}, 2},
		{"aaaaa", args{"aaaaa", "bba"}, -1},
		{"awfkasjflaskj", args{"awfkasjflaskj", "kj"}, 11},
		{"any", args{"any", ""}, 0},
		{"TheFirst", args{"TheFirst", "The"}, 0},
		{"TheNotFound", args{"TheNotFound", "TheF"}, -1},
		{"mississippi", args{"mississippi", "issip"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
