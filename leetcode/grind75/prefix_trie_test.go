package grind75

import "testing"

func TestPrefixTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("apple")
	if got := trie.Search("apple"); got != true {
		t.Fatalf("got %v; want true", got)
	}
	if got := trie.Search("app"); got != false {
		t.Fatalf("got %v; want false", got)
	}
	if got := trie.StartsWith("app"); got != true {
		t.Fatalf("got %v; want false", got)
	}
	trie.Insert("app")
	if got := trie.Search("app"); got != true {
		t.Fatalf("got %v; want true", got)
	}
}
