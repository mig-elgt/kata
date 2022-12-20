package mergeintervals

import "testing"

func TestCanAttendAllAppointments(t *testing.T) {
	testCases := []struct {
		in   [][]int
		want bool
	}{
		{[][]int{{1, 4}, {2, 5}, {7, 9}}, false},
		{[][]int{{6, 7}, {2, 4}, {8, 12}}, true},
		{[][]int{{4, 5}, {2, 3}, {3, 6}}, false},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := CanAttendAllAppointments(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
