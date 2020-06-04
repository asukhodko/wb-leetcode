package _220_contains_duplicate_iii

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	abs := func(a int) int {
		if a >= 0 {
			return a
		} else {
			return -a
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}
