package sliding_window

import "testing"

func TestFruitsIntoBaskets(t *testing.T) {
	testCases := map[string]struct {
		in   []rune
		want int
	}{
		"base case: 3 max": {
			in:   []rune{'A', 'B', 'C', 'A', 'C'},
			want: 3,
		},
		"base case: 5 max": {
			in:   []rune{'A', 'B', 'C', 'B', 'B', 'C'},
			want: 5,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := FruitsIntoBaskets(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
