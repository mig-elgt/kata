package sliding_window

import (
	"reflect"
	"testing"
)

func TestWordsConcatenation(t *testing.T) {
	testCases := map[string]struct {
		str   string
		words []string
		want  []int
	}{
		"base case 1": {
			str:   "catfoxcat",
			words: []string{"cat", "fox"},
			want:  []int{0, 3},
		},
		"base case 2": {
			str:   "catcatfoxfox",
			words: []string{"cat", "fox"},
			want:  []int{3},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := WordsConcatenation(tc.str, tc.words); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
