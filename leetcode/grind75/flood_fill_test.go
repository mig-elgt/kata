package grind75

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	testCases := []struct {
		in            [][]int
		sr, sc, color int
		want          [][]int
	}{
		{
			in: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			sr:    1,
			sc:    1,
			color: 2,
			want: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		{
			in: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			sr:    0,
			sc:    0,
			color: 0,
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			in: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			sr:    1,
			sc:    0,
			color: 2,
			want: [][]int{
				{2, 2, 2},
				{2, 2, 2},
			},
		},
	}
	for _, tc := range testCases {
		if got := floodFill(tc.in, tc.sr, tc.sc, tc.color); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
