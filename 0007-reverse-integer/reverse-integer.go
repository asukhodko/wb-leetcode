package _007_reverse_integer

// https://leetcode.com/problems/reverse-integer/
// Given a 32-bit signed integer, reverse digits of an integer.
//
//Example 1:
//
//Input: 123
//Output: 321
//Example 2:
//
//Input: -123
//Output: -321
//Example 3:
//
//Input: 120
//Output: 21
//Note:
//Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
//For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Integer.
// Memory Usage: 2.2 MB, less than 80.00% of Go online submissions for Reverse Integer.
func reverse(x int) (result int) {
	for x != 0 {
		mod := x % 10
		result = result*10 + mod
		if result > 0x7fffffff || result < -0x80000000 {
			return 0
		}
		x /= 10
	}
	return
}
