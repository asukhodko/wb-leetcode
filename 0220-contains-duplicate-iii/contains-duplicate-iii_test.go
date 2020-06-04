package _220_contains_duplicate_iii

import "testing"

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
		t    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3, 1},
				k:    3,
				t:    0,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1, 0, 1, 1},
				k:    1,
				t:    2,
			},
			want: true,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{1, 5, 9, 1, 5, 9},
				k:    2,
				t:    3,
			},
			want: false,
		},
		{
			name: "Example 4",
			args: args{
				nums: []int{1, 2},
				k:    0,
				t:    1,
			},
			want: false,
		},
		{
			name: "Example 5",
			args: args{
				nums: []int{7, 1, 3},
				k:    2,
				t:    3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyAlmostDuplicate(tt.args.nums, tt.args.k, tt.args.t); got != tt.want {
				t.Errorf("containsNearbyAlmostDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
