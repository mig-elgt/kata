package grind75

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	testCases := []struct {
		in   []string
		k    int
		want []string
	}{
		{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2, []string{"i", "love"}},
		{[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4, []string{"the", "is", "sunny", "day"}},
	}
	for _, tc := range testCases {
		t.Run("base cases: ", func(t *testing.T) {
			if got := topKFrequent(tc.in, tc.k); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
