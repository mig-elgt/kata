package grind75

import (
	"reflect"
	"testing"
)

func TestCodecEncode(t *testing.T) {
	testCases := []struct {
		in   []string
		want string
	}{
		{[]string{"foo", "bar"}, "13foo13bar"},
		{[]string{"hello world foo", "hello foo"}, "215hello world foo19hello foo"},
	}
	for _, tc := range testCases {
		codecer := Codec{}
		if got := codecer.Encode(tc.in); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}

func TestCodecDecode(t *testing.T) {
	testCases := []struct {
		in   string
		want []string
	}{
		{"13foo13bar", []string{"foo", "bar"}},
		{"215hello world foo19hello foo", []string{"hello world foo", "hello foo"}},
	}
	for _, tc := range testCases {
		codecer := Codec{}
		if got := codecer.Decode(tc.in); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
