package _066_plus_one

//https://leetcode.com/problems/plus-one/

//Given a non-empty array of digits representing a non-negative integer, plus one to the integer.
//
//The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.
//
//You may assume the integer does not contain any leading zero, except the number 0 itself.
//
//Example 1:
//
//Input: [1,2,3]
//Output: [1,2,4]
//Explanation: The array represents the integer 123.
//Example 2:
//
//Input: [4,3,2,1]
//Output: [4,3,2,2]
//Explanation: The array represents the integer 4321.

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Plus One.
//Memory Usage: 2 MB, less than 100.00% of Go online submissions for Plus One.
func plusOne(digits []int) []int {
	i := len(digits) - 1
	digits[i]++
	for digits[i] > 9 {
		digits[i] = 0
		if i == 0 {
			digits = append([]int{0}, digits...)
			i++
		}
		digits[i-1]++
		i--
	}
	return digits
}
