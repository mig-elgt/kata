package grind75

import "testing"

func TestNumIslands(t *testing.T) {
	testCases := []struct {
		in   [][]byte
		want int
	}{
		{
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			3,
		},
		{
			[][]byte{
				{'1'},
				{'1'},
			},
			1,
		},
	}
	for _, tc := range testCases {
		if got := numIslands(tc.in); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
