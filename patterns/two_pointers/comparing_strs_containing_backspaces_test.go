package two_pointers

import "testing"

func TestComparingStrsContainingBackspaces(t *testing.T) {
	testCases := []struct {
		str1 string
		str2 string
		want bool
	}{
		{"xy#z", "xzz#", true},
		{"xy#z", "xyz#", false},
		{"xp#", "xyz##", true},
		{"xywrrmp", "xywrrmu#p", true},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := ComparingStrsContainingBackspaces(tc.str1, tc.str2); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
