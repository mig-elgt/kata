package two_pointers

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testCases := map[string]struct {
		in   []int
		want int
	}{
		"base case 1": {
			in:   []int{2, 3, 3, 3, 6, 9, 9},
			want: 4,
		},
		"base case 2": {
			in:   []int{2, 2, 2, 11},
			want: 2,
		},
		// "not duplicates": {
		// 	in:   []int{1, 2, 3, 4},
		// 	want: 0,
		// },
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := RemoveDuplicates(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
