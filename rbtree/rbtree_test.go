package rbtree

import (
	"log"
	"testing"
)

func TestLeftRotate(t *testing.T) {
	var i10 Int = 10
	var i12 Int = 12

	rbtree := New()
	x := &node{rbtree.nilNode, rbtree.nilNode, rbtree.nilNode, BLACK, i10}
	rbtree.root = x
	y := &node{rbtree.root.rchild, rbtree.nilNode, rbtree.nilNode, RED, i12}
	rbtree.root.rchild = y

	log.Println("root : ", rbtree.root)
	log.Println("left : ", rbtree.root.lchild)
	log.Println("right : ", rbtree.root.rchild)

	rbtree.leftRotate(rbtree.root)

	log.Println("root : ", rbtree.root)
	log.Println("left : ", rbtree.root.lchild)
	log.Println("right : ", rbtree.root.rchild)

}

func TestRightRotate(t *testing.T) {
	var i10 Int = 10
	var i12 Int = 12

	rbtree := New()
	x := &node{rbtree.nilNode, rbtree.nilNode, rbtree.nilNode, BLACK, i10}
	rbtree.root = x
	y := &node{rbtree.root.rchild, rbtree.nilNode, rbtree.nilNode, RED, i12}
	rbtree.root.rchild = y

	log.Println("root : ", rbtree.root)
	log.Println("left : ", rbtree.root.lchild)
	log.Println("right : ", rbtree.root.rchild)

	rbtree.rightRotate(rbtree.root)

	log.Println("root : ", rbtree.root)
	log.Println("left : ", rbtree.root.lchild)
	log.Println("right : ", rbtree.root.rchild)

}

func TestReplaceOrInsertT(t *testing.T) {
	rbtree := New()

	rbtree.ReplaceOrInsert(&node{rbtree.nilNode, rbtree.nilNode, rbtree.nilNode, RED, Int(10)})
	rbtree.ReplaceOrInsert(&node{rbtree.nilNode, rbtree.nilNode, rbtree.nilNode, RED, Int(9)})
	rbtree.ReplaceOrInsert(&node{rbtree.nilNode, rbtree.nilNode, rbtree.nilNode, RED, Int(8)})
	rbtree.ReplaceOrInsert(&node{rbtree.nilNode, rbtree.nilNode, rbtree.nilNode, RED, Int(6)})
	rbtree.ReplaceOrInsert(&node{rbtree.nilNode, rbtree.nilNode, rbtree.nilNode, RED, Int(7)})

	log.Println("rbtree counts : ", rbtree.count)

	log.Println("------ ", rbtree.root.Item)
	log.Println("----", rbtree.root.lchild.Item, "---", rbtree.root.rchild.Item)
	log.Println("--", rbtree.root.lchild.lchild.Item, "-", rbtree.root.lchild.rchild.Item)

}
