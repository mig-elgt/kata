package islands

import (
	"reflect"
	"testing"
)

func Test_RemoveIsland(t *testing.T) {
	testCases := map[string]struct {
		in  [][]int
		out [][]int
	}{
		"two black pixel removed": {
			in: [][]int{
				[]int{1, 0, 1, 0, 0},
				[]int{0, 1, 0, 1, 1},
				[]int{0, 0, 0, 1, 0},
				[]int{0, 1, 0, 1, 0},
				[]int{0, 0, 0, 0, 1},
			},
			out: [][]int{
				[]int{1, 0, 1, 0, 0},
				[]int{0, 0, 0, 1, 1},
				[]int{0, 0, 0, 1, 0},
				[]int{0, 0, 0, 1, 0},
				[]int{0, 0, 0, 0, 1},
			},
		},
		"three black pixel removed": {
			in: [][]int{
				[]int{1, 0, 0, 0, 0, 0},
				[]int{0, 1, 0, 1, 1, 1},
				[]int{0, 0, 1, 0, 1, 0},
				[]int{1, 1, 0, 0, 1, 0},
				[]int{1, 0, 1, 1, 0, 0},
				[]int{1, 0, 0, 0, 0, 1},
			},
			out: [][]int{
				[]int{1, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 1, 1, 1},
				[]int{0, 0, 0, 0, 1, 0},
				[]int{1, 1, 0, 0, 1, 0},
				[]int{1, 0, 0, 0, 0, 0},
				[]int{1, 0, 0, 0, 0, 1},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := removeIsland(tc.in); !reflect.DeepEqual(got, tc.out) {
				t.Fatalf("got %v; want %v", got, tc.out)
			}
		})
	}
}
