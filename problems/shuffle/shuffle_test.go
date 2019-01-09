package answers

import "testing"

func TestIsRiffle(t *testing.T) {
	tt := []struct {
		h1           []int
		h2           []int
		shuffledDeck []int
		want         bool
	}{
		{[]int{1, 4}, []int{2, 5}, []int{1, 5, 2, 4, 7}, false},
		{[]int{1, 4}, []int{2, 5}, []int{1, 4, 2, 5}, true},
		{[]int{1, 5, 8, 9, 10}, []int{2, 4, 7, 11, 22}, []int{1, 2, 4, 5, 7, 8, 9, 10, 11, 22}, true},
	}
	for _, tc := range tt {
		ok := isRiffle(tc.h1, tc.h2, tc.shuffledDeck)
		if ok != tc.want {
			t.Errorf("expected %v, got %v", tc.want, ok)
		}
	}
}

func TestGetCounts(t *testing.T) {
	tt := []struct {
		slc  []int
		want map[int]int
	}{
		{
			[]int{1, 2, 1, 5, 8, 1, 5},
			map[int]int{
				1: 3,
				2: 1,
				5: 2,
				8: 1,
			},
		},
	}
	for _, tc := range tt {
		got := getCounts(tc.slc)
		if len(got) != len(tc.want) {
			t.Errorf("maps different sizes: got map of length %d, expected map of length %d", len(got), len(tc.want))
		}
		for k, v := range got {
			w, ok := tc.want[k]
			if !ok {
				t.Errorf("maps not equivalent: got %v, expected %v", got, tc.want)
			}
			if v != w {
				t.Errorf("maps not equivalent: got %v, expected %v", got, tc.want)
			}
		}
	}
}
