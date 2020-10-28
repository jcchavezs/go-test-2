package leftpad

func LeftPad(s string, padChar string, length int) string {
	if len(s) <= length {
		return s
	}

	for i := 0; i < length-len(s); i++ {
		s = padChar + s
	}

	return s
}
