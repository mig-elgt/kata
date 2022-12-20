package sliding_window

import (
	"reflect"
	"testing"
)

func TestStringAnagrams(t *testing.T) {
	testCases := map[string]struct {
		in      string
		pattern string
		want    []int
	}{
		"two anagrams": {
			in:      "ppqp",
			pattern: "pq",
			want:    []int{1, 2},
		},
		"three anagrams": {
			in:      "abbcabc",
			pattern: "abc",
			want:    []int{2, 3, 4},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := StringAnagrams(tc.in, tc.pattern); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
