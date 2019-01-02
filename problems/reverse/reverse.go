package answers

func reverseInPlace(chars []string) []string {
	if len(chars) == 0 || len(chars) == 1 {
		return chars
	}

	for i := 0; i < len(chars)/2; i++ {
		chars[i], chars[len(chars)-1-i] = chars[len(chars)-1-i], chars[i]
	}

	return chars
}
