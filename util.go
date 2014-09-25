package buildver

import "unicode"

func min(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}

func notDigitOrDot(r rune) bool {
	return !(r == '.' || unicode.IsDigit(r))
}
