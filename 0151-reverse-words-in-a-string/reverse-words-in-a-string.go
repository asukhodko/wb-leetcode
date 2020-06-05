package _151_reverse_words_in_a_string

import "strings"

func reverseWords(s string) string {
	word := []rune{}
	result := []string{}
	for _, c := range s {
		if c == ' ' {
			if len(word) > 0 {
				result = append([]string{string(word)}, result...)
				word = []rune{}
			}
			continue
		}
		word = append(word, c)
	}
	if len(word) > 0 {
		result = append([]string{string(word)}, result...)
	}
	return strings.Join(result, " ")
}
