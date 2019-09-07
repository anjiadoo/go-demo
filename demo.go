package main

import(
	"fmt"
)

type RouteGuideServer interface{
	Do()
}

type server struct{
	addr *Addr
}

func (s *server) Do(){
	fmt.Println("------")
}

type Addr struct{}

func newServer() RouteGuideServer{
	return &server{
		addr: (*Addr)(nil),
	}
}

func main(){
	s := newServer()
	s.(RouteGuideServer).Do()
}