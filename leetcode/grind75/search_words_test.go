package grind75

import "testing"

// ["WordDictionary","addWord","addWord","addWord","addWord","search","search","addWord","search","search","search","search","search","search"]
// [[],["at"],["and"],["an"],["add"],["a"],[".at"],["bat"],[".at"],["an."],["a.d."],["b."],["a.d"],["."]]
func TestWordDictionary(t *testing.T) {
	wd := NewWordDictionary()
	wd.AddWord("at")
	wd.AddWord("and")
	wd.AddWord("an")
	wd.AddWord("add")
	if got := wd.Search("a"); got != false {
		t.Fatalf("got %v; want false", got)
	}
	if got := wd.Search(".at"); got != false {
		t.Fatalf("got %v; want false", got)
	}
	wd.AddWord("bat")
	if got := wd.Search(".at"); got != true {
		t.Fatalf("got %v; want true", got)
	}
	if got := wd.Search("an."); got != true {
		t.Fatalf("got %v; want true", got)
	}
	if got := wd.Search("a.d."); got != false {
		t.Fatalf("got %v; want false", got)
	}
	if got := wd.Search("b."); got != false {
		t.Fatalf("got %v; want false", got)
	}
	if got := wd.Search("a.d"); got != true {
		t.Fatalf("got %v; want true", got)
	}
	if got := wd.Search("."); got != false {
		t.Fatalf("got %v; want false", got)
	}
}
