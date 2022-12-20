package graph

import "testing"

func TestGetAlienDictionaryOrder(t *testing.T) {
	testCases := []struct {
		in   []string
		want string
	}{
		{[]string{"ba", "bc", "ac", "cab"}, "bac"},
		{[]string{"cab", "aaa", "aab"}, "cab"},
		{[]string{"ywx", "wz", "xww", "xz", "zyy", "zwz"}, "ywxz"},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := GetAlienDictionaryOrder(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
