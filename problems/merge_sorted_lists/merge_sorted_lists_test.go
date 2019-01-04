package answers

import "testing"

func TestMergeLists(t *testing.T) {
	tt := []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{[]int{3, 4, 6, 10, 11, 15}, []int{1, 5, 8, 12, 14, 19}, []int{1, 3, 4, 5, 6, 8, 10, 11, 12, 14, 15, 19}},
		{[]int{1, 2, 10}, []int{3}, []int{1, 2, 3, 10}},
	}

	for _, tc := range tt {
		got := mergeLists(tc.l1, tc.l2)
		if len(got) != len(tc.want) {
			t.Errorf("length expected is %d, length got is %d", len(tc.want), len(got))
		}
		for i := 0; i < len(got); i++ {
			if got[i] != tc.want[i] {
				t.Errorf("expected %v, got %v", tc.want, got)
			}
		}
	}
}
