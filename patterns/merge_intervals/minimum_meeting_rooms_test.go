package mergeintervals

import "testing"

func TestFindMinimumMeetingRooms(t *testing.T) {
	testCases := []struct {
		in   [][]int
		want int
	}{
		{[][]int{{1, 4}, {2, 5}, {7, 9}}, 2},
		{[][]int{{6, 7}, {2, 4}, {8, 12}}, 1},
		{[][]int{{1, 4}, {2, 3}, {3, 6}}, 2},
		{[][]int{{4, 5}, {2, 3}, {2, 4}, {3, 5}}, 2},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := FindMinimumMeetingRooms(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
