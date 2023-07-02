package slidingwindow

import "testing"

func TestFruiBaskets(t *testing.T) {
	testCases := []struct {
		fruits []int
		want   int
	}{
		{[]int{1, 2, 1}, 3},
		{[]int{0, 1, 2, 2}, 3},
		{[]int{1, 2, 3, 2, 2}, 4},
	}
	for _, tc := range testCases {
		if got := totalFruit(tc.fruits); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
