package main

import "fmt"

type Movie struct {
	Name    string
	Ranking int
}

type Node struct {
	Next  *Node
	Value Movie
}

type List struct {
	Head *Node
}

func (list *List) InsertSort(movie *Node) {
	if list.Head == nil {
		list.Head = movie
		return
	}
	tmp := list.Head
	prev := list.Head
	for tmp != nil {
		if movie.Value.Ranking <= tmp.Value.Ranking {
			break
		}
		prev = tmp
		tmp = tmp.Next
	}
	if tmp == list.Head {
		movie.Next = list.Head
		list.Head = movie
		return
	}
	if tmp == nil {
		prev.Next = movie
		return
	}
	movie.Next = tmp
	prev.Next = movie
}

func (list List) String() string {
	var str string
	tmp := list.Head
	for tmp != nil {
		str += fmt.Sprintf("%v: %v; ", tmp.Value.Name, tmp.Value.Ranking)
		tmp = tmp.Next
	}
	return str
}

func main() {
	mexicans := []Movie{{"A", 20}, {"B", 5}, {"C", 1}, {"D", 4}}
	europe := []Movie{{"E", 3}, {"F", 10}, {"G", 11}, {"H", 40}}
	us := []Movie{{"K", 2}, {"L", 3}, {"M", 13}, {"N", 6}}

	mxList := initList(mexicans)
	euList := initList(europe)
	usList := initList(us)

	top := combine(mxList, euList, usList)
	fmt.Println(top)

}

func initList(movies []Movie) *List {
	list := &List{}
	for _, m := range movies {
		list.InsertSort(&Node{
			Value: m,
		})
	}
	return list
}

func combine(lists ...*List) *List {
	topMovies := lists[0]
	for i := 1; i < len(lists); i++ {
		tmp := lists[i].Head
		for tmp != nil {
			topMovies.InsertSort(&Node{Value: tmp.Value})
			tmp = tmp.Next
		}
	}
	return topMovies
}
