package answers

import (
	"strings"
)

func reverseWords(msg []string) []string {
	s := strings.Join(msg, "")
	msg2 := strings.SplitAfter(s, " ")
	dec := []string{}
	for i := len(msg2) - 1; i >= 0; i-- {
		dec = append(dec, msg2[i])
	}

	trimmed := []string{}
	for _, item := range dec {
		trimmed = append(trimmed, strings.TrimSpace(item))
	}

	r := strings.Join(trimmed, " ")
	ret := strings.Split(r, "")
	return ret
}
