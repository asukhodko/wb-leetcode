package _066_plus_one

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{"[1,2,3]", args{[]int{1, 2, 3}}, []int{1, 2, 4}},
		{"[1,2,9]", args{[]int{1, 2, 9}}, []int{1, 3, 0}},
		{"[2,2,9]", args{[]int{2, 2, 9}}, []int{2, 3, 0}},
		{"[4,3,2,1]", args{[]int{4, 3, 2, 1}}, []int{4, 3, 2, 2}},
		{"[1,2,9,9,9,9]", args{[]int{1, 2, 9, 9, 9, 9}}, []int{1, 3, 0, 0, 0, 0}},
		{"[9,9,9,9]", args{[]int{9, 9, 9, 9}}, []int{1, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := plusOne(tt.args.digits); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("plusOne() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
