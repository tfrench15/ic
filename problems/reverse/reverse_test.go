package answers

import "testing"

func TestReverseInPlace(t *testing.T) {
	tt := []struct {
		chars []string
		want  []string
	}{
		{[]string{"a", "b", "c", "d", "e"}, []string{"e", "d", "c", "b", "a"}},
		{[]string{"a", "b", "c", "d", "e", "f"}, []string{"f", "e", "d", "c", "b", "a"}},
	}

	for _, tc := range tt {
		got := reverseInPlace(tc.chars)
		if len(got) != len(tc.want) {
			t.Errorf("got length %d, wanted length %d", len(got), len(tc.want))
		}
		for i := 0; i < len(tc.want); i++ {
			if got[i] != tc.want[i] {
				t.Errorf("got %v, expected %v", got, tc.want)
			}
		}
	}
}
