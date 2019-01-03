package answers

import "testing"

func TestReverseWords(t *testing.T) {
	msg := []string{"c", "a", "k", "e", " ", "p", "o", "u", "n", "d", " ", "s", "t", "e", "a", "l"}
	want := []string{"s", "t", "e", "a", "l", " ", "p", "o", "u", "n", "d", " ", "c", "a", "k", "e"}
	got := reverseWords(msg)
	if len(got) != len(want) {
		t.Errorf("length expected is %d, length got is %d", len(want), len(got))
	}
	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Errorf("expected %v, got %v", want, got)
		}
	}
}
