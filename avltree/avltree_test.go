package avltree

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestAVLTree(t *testing.T) {
	var root *Tree
	vals := rand.Perm(10)
	fmt.Println(vals)
	for i := 0; i < len(vals); i++ {
		root = InsertNode(root, vals[i], struct{}{})
	}

	print("root avl树前序遍历(len=", len(vals), ")：\n")
	PrePrintTree(root)
}
