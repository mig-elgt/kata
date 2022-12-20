package slidingwindow

import "testing"

func TestCharacterReplacement(t *testing.T) {
	testCases := []struct {
		in   string
		k    int
		want int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
	}
	for _, tc := range testCases {
		t.Run("base cases: ", func(t *testing.T) {
			if got := characterReplacement(tc.in, tc.k); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
