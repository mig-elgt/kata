package bfs

import "testing"

func TestLadderLength(t *testing.T) {
	testCases := []struct {
		beginWord, endWord string
		wordList           []string
		want               int
	}{
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0},
	}
	for _, tc := range testCases {
		if got := ladderLength(tc.beginWord, tc.endWord, tc.wordList); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
