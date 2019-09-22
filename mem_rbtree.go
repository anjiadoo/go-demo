package main

import (
	"fmt"
	"go-demo/rbtree"
	"log"
	"math/rand"
)

func main() {
	tree := rbtree.New()
	vals := rand.Perm(20)

	fmt.Println(vals)
	for i := 0; i < len(vals); i++ {
		tree.ReplaceOrInsert(rbtree.Int(vals[i]))
	}

	log.Println("rbtree counts : ", tree.GetNodeCnt())
	rbtree.LeverOrder(tree.GetRootNode())
	//rbtree.PreOrder(tree.GetRootNode())
}
