package arrays

import "testing"

func TestMinMeetingRooms(t *testing.T) {
	testCases := []struct {
		in   [][]int
		want int
	}{
		{[][]int{{0, 30}, {5, 10}, {15, 20}}, 2},
		{[][]int{{7, 10}, {2, 4}}, 1},
		{[][]int{{0, 8}, {3, 5}, {4, 8}, {10, 12}, {13, 18}, {14, 17}}, 3},
	}
	for _, tc := range testCases {
		t.Run("base cases: ", func(t *testing.T) {
			if got := minMeetingRooms(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
