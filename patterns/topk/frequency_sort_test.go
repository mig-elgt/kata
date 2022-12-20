package topk

import "testing"

func TestSortCharacterByFrequency(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{"Programming", "rrggmmPiano"},
		{"abcbab", "bbbaac"},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := SortCharacterByFrequency(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
