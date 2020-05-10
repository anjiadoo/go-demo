package rbtree

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
)

func TestLeverOrder(t *testing.T) {
	tree := New()
	vals := rand.Perm(20)

	fmt.Println(vals)
	for i := 0; i < len(vals); i++ {
		tree.ReplaceOrInsert(Int(vals[i]))
	}

	log.Println("rbtree counts : ", tree.GetNodeCnt())
	LeverOrder(tree.GetRootNode())
	//PreOrder(tree.GetRootNode())
}


//import (
//	"log"
//	"testing"
//)
//
//func TestLeftRotate(t *testing.T) {
//	var i10 Int = 10
//	var i12 Int = 12
//
//	rbtree := New()
//	x := &node{nilNode, nilNode, nilNode, BLACK, i10}
//	root = x
//	y := &node{root.rchild, nilNode, nilNode, RED, i12}
//	root.rchild = y
//
//	log.Println("root : ", root)
//	log.Println("left : ", root.lchild)
//	log.Println("right : ", root.rchild)
//
//	leftRotate(root)
//
//	log.Println("root : ", root)
//	log.Println("left : ", root.lchild)
//	log.Println("right : ", root.rchild)
//
//}
//
//func TestRightRotate(t *testing.T) {
//	var i10 Int = 10
//	var i12 Int = 12
//
//	rbtree := New()
//	x := &node{nilNode, nilNode, nilNode, BLACK, i10}
//	root = x
//	y := &node{root.rchild, nilNode, nilNode, RED, i12}
//	root.rchild = y
//
//	log.Println("root : ", root)
//	log.Println("left : ", root.lchild)
//	log.Println("right : ", root.rchild)
//
//	rightRotate(root)
//
//	log.Println("root : ", root)
//	log.Println("left : ", root.lchild)
//	log.Println("right : ", root.rchild)
//
//}
//
//func TestReplaceOrInsertT(t *testing.T) {
//	rbtree := New()
//
//	ReplaceOrInsert(&node{nilNode, nilNode, nilNode, RED, Int(10)})
//	ReplaceOrInsert(&node{nilNode, nilNode, nilNode, RED, Int(9)})
//	ReplaceOrInsert(&node{nilNode, nilNode, nilNode, RED, Int(8)})
//	ReplaceOrInsert(&node{nilNode, nilNode, nilNode, RED, Int(6)})
//	ReplaceOrInsert(&node{nilNode, nilNode, nilNode, RED, Int(7)})
//
//	log.Println("rbtree counts : ", count)
//
//	log.Println("------ ", root.Item)
//	log.Println("----", root.lchild.Item, "---", root.rchild.Item)
//	log.Println("--", root.lchild.lchild.Item, "-", root.lchild.rchild.Item)
//
//}
