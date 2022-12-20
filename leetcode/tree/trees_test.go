package trees

import (
	"testing"
)

func createTree(numbers ...int) *tree {
	t := tree{}
	for _, n := range numbers {
		t.Insert(n)
	}
	return &t
}

func Test_treeHeight(t *testing.T) {
	testCases := map[string]struct {
		in   *tree
		want int
	}{
		"zero": {
			in:   &tree{},
			want: 0,
		},
		"base case": {
			in:   createTree(5, 4, 3, 2, 1, 10, 6, 8, 9),
			want: 5,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			treeNumber := tc.in
			if got, want := treeNumber.Height(), tc.want; got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
