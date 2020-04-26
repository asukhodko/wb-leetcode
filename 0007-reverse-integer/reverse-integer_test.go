package _007_reverse_integer

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "123", args: args{x: 123}, want: 321},
		{name: "-123", args: args{x: -123}, want: -321},
		{name: "321", args: args{x: 321}, want: 123},
		{name: "-321", args: args{x: -321}, want: -123},
		{name: "1534236469", args: args{x: 1534236469}, want: 0},
		{name: "1563847412", args: args{x: 1563847412}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
