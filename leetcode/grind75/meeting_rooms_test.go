package grind75

import "testing"

func TestCanAttendMeetings(t *testing.T) {
	testCases := []struct {
		in   [][]int
		want bool
	}{
		{[][]int{{0, 30}, {5, 10}, {15, 20}}, false},
		{[][]int{{7, 10}, {2, 4}}, true},
		{[][]int{{5, 8}, {6, 8}}, false},
	}
	for _, tc := range testCases {
		t.Run("base cases: ", func(t *testing.T) {
			if got := canAttendMeetings(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
