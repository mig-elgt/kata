package grind75

import "testing"

func TestFindMedianStream(t *testing.T) {
	testCases := []struct {
		stream  []int
		medians []float64
	}{
		{[]int{6, 10, 2, 6, 5, 0, 6, 3, 1, 0, 0}, []float64{6, 8, 6, 6, 6, 5.5, 6, 5.5, 5, 4, 3}},
	}
	for _, tc := range testCases {
		mf := NewMedianFinder()
		for i := 0; i < len(tc.stream); i++ {
			mf.AddNum(tc.stream[i])
			if got, want := mf.FindMedian(), tc.medians[i]; got != want {
				t.Fatalf("num(%v); got %v; want %v", tc.stream[i], got, want)
			}
		}
	}
}
