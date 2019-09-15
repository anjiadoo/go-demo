package main

import (
	"fmt"
	"go-demo/avltree"
	"math/rand"
)

func main() {
	var root *avltree.Tree
	vals := rand.Perm(10)
	fmt.Println(vals)
	for i := 0; i < len(vals); i++ {
		root = avltree.InsertNode(root, vals[i], struct{}{})
	}

	print("root avl树前序遍历(len=", len(vals), ")：\n")
	avltree.PrePrintTree(root)
}
