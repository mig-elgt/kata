package grind75

import "testing"

func TestGetFood(t *testing.T) {
	testCases := []struct {
		in   [][]byte
		want int
	}{
		{[][]byte{
			{'X', 'X', 'X', 'X', 'X', 'X'},
			{'X', '*', 'O', 'O', 'O', 'X'},
			{'X', 'O', 'O', '#', 'O', 'X'},
			{'X', 'X', 'X', 'X', 'X', 'X'},
		}, 3},
		{[][]byte{
			{'X', 'X', 'X', 'X', 'X'},
			{'X', '*', 'X', 'O', 'X'},
			{'X', 'O', 'X', '#', 'X'},
			{'X', 'X', 'X', 'X', 'X'},
		}, -1},
		{[][]byte{
			{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
			{'X', '*', 'O', 'X', 'O', '#', 'O', 'X'},
			{'X', 'O', 'O', 'X', 'O', 'O', 'X', 'X'},
			{'X', 'O', 'O', 'O', 'O', '#', 'O', 'X'},
			{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		}, 6},
		{[][]byte{
			{'O', '*'},
			{'#', 'O'},
		}, 2},
	}
	for _, tc := range testCases {
		t.Run("base cases: ", func(t *testing.T) {
			if got := getFood(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
