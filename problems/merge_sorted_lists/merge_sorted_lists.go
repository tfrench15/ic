package answers

import "fmt"

func mergeLists(l1, l2 []int) []int {
	i, j := 0, 0
	var ret []int
	for i < len(l1) || j < len(l2) {
		// case: l1 is done.
		if i == len(l1) {
			ret = append(ret, l2[j])
			j++
			if j == len(l2) {
				return ret
			}
			continue
		}
		// case: l2 is done.
		if j == len(l2) {
			ret = append(ret, l1[i])
			i++
			if i == len(l1) {
				return ret
			}
			continue
		}
		if l1[i] <= l2[j] {
			ret = append(ret, l1[i])
			i++
			continue
		}
		if l1[i] > l2[j] {
			ret = append(ret, l2[j])
			j++
			continue
		}
	}
	fmt.Println(ret)
	return ret
}
