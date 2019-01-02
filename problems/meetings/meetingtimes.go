package answers

import (
	"sort"
)

// Meeting contains the start time and end time of a meeting.
type Meeting struct {
	Start int
	End   int
}

// Meetings implements sort.Interface for []Meeting based on the
// Start field.
type Meetings []Meeting

func (m Meetings) Len() int {
	return len(m)
}

func (m Meetings) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m Meetings) Less(i, j int) bool {
	return m[i].Start < m[j].Start
}

// MergeMeetings returns a merged set of overlapping
// meetings.
func MergeMeetings(m []Meeting) Meetings {
	var (
		merged  Meetings
		current Meeting
	)
	sort.Sort(Meetings(m))
	current.Start, current.End = m[0].Start, m[0].End
	for i := 1; i < len(m); i++ {
		if m[i].Start <= current.End {
			if m[i].End > current.End {
				current.End = m[i].End
				continue
			}
		} else {
			merged = append(merged, current)
			current.Start, current.End = m[i].Start, m[i].End
		}
	}
	merged = append(merged, current)
	return merged
}
