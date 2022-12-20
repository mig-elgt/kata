package arrays

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		in   []string
		want [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{
			{"eat", "tea", "ate"},
			{"tan", "nat"},
			{"bat"},
		}},
		{[]string{""}, [][]string{
			{""},
		}},
		{[]string{"a"}, [][]string{
			{"a"},
		}},
		{[]string{"ac", "c"}, [][]string{
			{"ac"},
			{"c"},
		}},
	}
	for _, tc := range testCases {
		t.Run("base cases: ", func(t *testing.T) {
			if got := groupAnagrams(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
