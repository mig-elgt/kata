package grind75

import "testing"

func TestExistWord(t *testing.T) {
	testCases := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCCED", true},
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "SEE", true},
	}
	for _, tc := range testCases {
		if got := existWord(tc.board, tc.word); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
