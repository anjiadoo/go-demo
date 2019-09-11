package main

import (
	"fmt"
	"sort"
)

type RouteGuideServer interface {
	Do()
}

type server struct {
	addr *Addr
}

func (s *server) Do() {
	fmt.Println("------")
}

type Addr struct{}

func newServer() RouteGuideServer {
	return &server{
		addr: (*Addr)(nil),
	}
}

func main() {
	s := newServer()
	s.(RouteGuideServer).Do()

	var src = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//index := 4
	//dd := copy(src[index:], src[index+1:])
	//src = src[:len(src)-1]
	//fmt.Println("-", src)

	x := 5
	i := sort.Search(len(src), func(i int) bool { return src[i] >= x })
	print(i)
	if i < len(src) && src[i] == x {
		fmt.Println("x is present at data[i]")
	} else {
		fmt.Println("x not is present at data[i]")
		// x is not present in data,
		// but i is the index where it would be inserted.
	}
}
