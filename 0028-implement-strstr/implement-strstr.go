package _028_implement_strstr

//https://leetcode.com/problems/implement-strstr/

//Implement strStr().
//
//Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
//Example 1:
//
//Input: haystack = "hello", needle = "ll"
//Output: 2
//Example 2:
//
//Input: haystack = "aaaaa", needle = "bba"
//Output: -1
//Clarification:
//
//What should we return when needle is an empty string? This is a great question to ask during an interview.
//
//For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement strStr().
//Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Implement strStr().
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i, j := 0, 0; i < len(haystack); i++ {
		if haystack[i] == needle[j] {
			j++
			if j == len(needle) {
				return i - j + 1
			}
		} else {
			i -= j
			j = 0
		}
	}
	return -1
}
