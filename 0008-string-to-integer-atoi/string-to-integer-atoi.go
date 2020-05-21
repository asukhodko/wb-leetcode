package _008_string_to_integer_atoi

type parseState int

const (
	seekSign parseState = iota
	seekDigit
)

func myAtoi(str string) int {
	sign := 1
	result := 0
	var state parseState
Loop:
	for _, r := range str {
		switch state {
		case seekSign:
			switch {
			case r == ' ':
			case r == '+':
				state = seekDigit
			case r == '-':
				state = seekDigit
				sign = -1
			case r >= '0' && r <= '9':
				state = seekDigit
				v := int(r - '0')
				result = result*10 + v
			default:
				return 0
			}
		case seekDigit:
			if r >= '0' && r <= '9' {
				v := int(r - '0')
				result = result*10 + v
				if result > 0x7fffffff {
					if sign == 1 {
						return 0x7fffffff
					} else {
						return -0x80000000
					}
				}
			} else {
				break Loop
			}
		}
	}
	return sign * result
}
