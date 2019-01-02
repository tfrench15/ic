package answers

import (
	"fmt"
	"testing"
)

func TestMergeRanges(t *testing.T) {
	tt := []struct {
		calendar []Meeting
		want     Meetings
	}{
		{[]Meeting{
			Meeting{0, 1},
			Meeting{3, 5},
			Meeting{4, 8},
			Meeting{9, 10},
			Meeting{10, 12},
		}, []Meeting{
			Meeting{0, 1},
			Meeting{3, 8},
			Meeting{9, 12},
		}},
		{[]Meeting{
			Meeting{1, 2},
			Meeting{2, 3},
		}, []Meeting{
			Meeting{1, 3},
		}},
		{[]Meeting{
			Meeting{1, 5},
			Meeting{2, 3},
		}, []Meeting{
			Meeting{1, 5},
		}},
		{[]Meeting{
			Meeting{1, 10},
			Meeting{2, 6},
			Meeting{3, 5},
			Meeting{7, 9},
		}, []Meeting{
			Meeting{1, 10},
		}},
	}

	for _, tc := range tt {
		got := MergeMeetings(tc.calendar)
		if len(tc.want) != len(got) {
			fmt.Println("want is ", tc.want)
			fmt.Println("got is ", got)
			t.Fatalf("length of want is %d, length of got is %d", len(tc.want), len(got))
		}
		for i := 0; i < len(tc.want); i++ {
			if tc.want[i] != got[i] {
				t.Errorf("expected %v, got %v", tc.want, got)
			}
		}
	}

}
