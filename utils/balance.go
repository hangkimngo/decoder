package utils

func BalanceBracket(s string) bool {
	if len(s) == 0 {
		return false
	}
	brackets := []rune{}
	count_left := 0
	count_right := 0
	for _, c := range s {
		switch c {
		case '[':
			brackets = append(brackets, c)
			count_left++
		case ']':
			brackets = append(brackets, c)
			count_right++
		}
	}

	if count_left != count_right {
		return false
	}

	for i := 0; i < len(brackets); i = i + 2 {
		if brackets[i] == '[' && brackets[i+1] == ']' {
			continue
		} else if brackets[i] == ']' && brackets[i+1] == '[' {
			continue
		} else {
			return false
		}
	}
	return true
}
