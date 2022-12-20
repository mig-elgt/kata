package subsets

import (
	"reflect"
	"testing"
)

func TestGetStringPermutationsByChangingCase(t *testing.T) {
	testCases := []struct {
		in   string
		want []string
	}{
		{"ad52", []string{"ad52", "aD52", "Ad52", "AD52"}},
		{"ab7c", []string{"ab7c", "ab7C", "aB7c", "aB7C", "Ab7c", "Ab7C", "AB7c", "AB7C"}},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := GetStringPermutationsByChangingCase(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
