package answers

func isRiffle(h1, h2, fullDeck []int) bool {
	for _, card := range h2 {
		h1 = append(h1, card)
	}
	m1 := getCounts(h1)
	m2 := getCounts(fullDeck)
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		w, ok := m2[k]
		if !ok {
			return false
		}
		if v != w {
			return false
		}
	}
	return true
}

func getCounts(slc []int) map[int]int {
	m := make(map[int]int)
	for _, card := range slc {
		_, ok := m[card]
		if ok {
			m[card]++
		} else {
			m[card] = 1
		}
	}
	return m
}
