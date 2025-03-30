package aign

/*有效的括号*/
func isValid(s string) bool {
	left := []rune{}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			left = append(left, v)
		} else {
			if len(left) > 0 && isTrue(v) == left[len(left)-1] {
				left = left[:len(left)-1]
			} else {
				return false
			}
		}
	}
	return len(left) == 0
}
func isTrue(c rune) rune {
	if c == '(' {
		return ')'
	}
	if c == '[' {
		return ']'
	}
	return '{'
}
