package _151_reverse_words_in_a_string

func reverseWords(s string) string {
	word := []rune{}
	result := make([]rune, len(s))
	i := len(result)
	for _, c := range s {
		if c == ' ' {
			if len(word) > 0 {
				i -= len(word)
				copy(result[i:], word)
				word = []rune{}
				i--
				result[i] = ' '
			}
			continue
		}
		word = append(word, c)
	}
	if len(word) > 0 {
		i -= len(word)
		copy(result[i:], word)
	}
	if i < len(result) && result[i] == ' ' {
		i++
	}
	return string(result[i:])
}
