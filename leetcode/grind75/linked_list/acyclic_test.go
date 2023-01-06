package linked_list

import "testing"

func createAcylicList() *List {
	list := &List{}
	for i := 0; i < 10; i++ {
		list.Append(&Node{Value: i})
	}
	return list
}

func createCylicList() *List {
	list := &List{}
	var head, aux *Node
	for i := 0; i < 10; i++ {
		node := &Node{Value: i}
		if i == 2 {
			head = node
		}
		aux = node
		list.Append(node)
	}
	aux.Next = head
	return list
}

func TestIsAcyclic(t *testing.T) {
	testCases := map[string]struct {
		in   *List
		want bool
	}{
		"nil list": {
			in:   nil,
			want: true,
		},
		"empty list": {
			in:   &List{},
			want: true,
		},
		"acyclic list": {
			in:   createAcylicList(),
			want: true,
		},
		"cyclic list": {
			in:   createCylicList(),
			want: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got, want := IsAcyclic(tc.in), tc.want; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
		})
	}

}
