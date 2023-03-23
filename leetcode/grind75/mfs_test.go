package grind75

import (
	"reflect"
	"testing"
)

func TestFileSystem(t *testing.T) {
	fs := NewFileSystem()
	if got := fs.Ls("/"); !reflect.DeepEqual(got, []string{}) {
		t.Fatalf("got %v; want empty list", got)
	}
	fs.Mkdir("/a/b/c")
	fs.AddContentToFile("/a/b/c/d", "hello")
	if got := fs.ReadContentFromFile("/a/b/c/d"); got != "hello" {
		t.Fatalf("got %v; want hello value", got)
	}
	if got := fs.Ls("/"); !reflect.DeepEqual(got, []string{"a"}) {
		t.Fatalf("got %v; want [a]", got)
	}
	if got := fs.Ls("/a/b"); !reflect.DeepEqual(got, []string{"c"}) {
		t.Fatalf("got %v; want [c]", got)
	}
	if got := fs.Ls("/a/b/c/d"); !reflect.DeepEqual(got, []string{"d"}) {
		t.Fatalf("got %v; want [d]", got)
	}
}
