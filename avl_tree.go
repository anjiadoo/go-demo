package main

import (
	"math"
)

type Tree struct {
	parent *Tree
	lchild *Tree
	rchild *Tree
	depth  int
	val    int
	data   interface{}
}

// LL型，在node节点的左子树根节点的左子树上插入节点而破坏平衡->右旋
func LL_rotate(node *Tree) *Tree {
	var parent, son *Tree
	parent = node.parent
	son = node.lchild

	if son.rchild != nil {
		son.rchild.parent = node
	}
	node.lchild = son.rchild
	update_depth(node)

	son.rchild = node
	son.parent = parent

	if parent != nil {
		if parent.lchild == node {
			parent.lchild = son
		} else {
			parent.rchild = son
		}
	}
	node.parent = son
	update_depth(son)
	return son
}

// RR型，在node节点的右子树根节点的右子树上插入节点而破坏平衡->左旋
func RR_rotate(node *Tree) *Tree {
	var parent, son *Tree
	parent = node.parent
	son = node.rchild

	if son.lchild != nil {
		son.lchild.parent = node
	}
	node.rchild = son.lchild
	update_depth(node)

	son.lchild = node
	son.parent = parent
	if parent != nil {
		if parent.lchild == node {
			parent.lchild = son
		} else {
			parent.rchild = son
		}
	}
	node.parent = son
	update_depth(son)
	return son
}

//LR型，在node节点的左子树根节点的右子树上插入节点而破坏平衡->先左旋再右旋
func LR_rotate(node *Tree) *Tree {
	RR_rotate(node.lchild)
	return LL_rotate(node)
}

//LR型，在node节点的右子树根节点的左子树上插入节点而破坏平衡->先右旋再左旋
func RL_rotate(node *Tree) *Tree {
	LL_rotate(node.rchild)
	return RR_rotate(node)
}

func update_depth(node *Tree) {
	if node == nil {
		return
	}
	depth_Lchild := get_balance(node.lchild)
	depth_Rchild := get_balance(node.rchild)

	node.depth = int(math.Max(float64(depth_Lchild), float64(depth_Rchild))) + 1
}

func get_balance(node *Tree) int {
	if node == nil {
		return 0
	}
	return node.depth
}

func is_balance(node *Tree) int {
	if node == nil {
		return 0
	}
	return get_balance(node.lchild) - get_balance(node.rchild)
}

func insertNode(node *Tree, val int, data interface{}) *Tree {
	if node == nil {
		node = &Tree{
			parent: nil,
			lchild: nil,
			rchild: nil,
			depth:  0,
			val:    val,
			data:   data,
		}
	} else if val < node.val {
		node.lchild = insertNode(node.lchild, val, data)
		if is_balance(node) == 2 {
			if val < node.lchild.val {
				node = LL_rotate(node)
			} else {
				node = LR_rotate(node)
			}
		}
	} else if val > node.val {
		node.rchild = insertNode(node.rchild, val, data)
		if is_balance(node) == -2 {
			if val > node.rchild.val {
				node = RR_rotate(node)
			} else {
				node = RL_rotate(node)
			}
		}
	}
	node.depth = int(math.Max(float64(get_balance(node.lchild)), float64(get_balance(node.rchild)))) + 1
	return node
}

func prePrintTree(node *Tree) {
	if node == nil {
		return
	}
	print("node.val=", node.val)
	if node.lchild != nil {
		print(" lchild.val=", node.lchild.val)
	}
	if node.rchild != nil {
		print(" rchild.val=", node.rchild.val)
	}
	print("\n")
	prePrintTree(node.lchild)
	prePrintTree(node.rchild)
}

func main() {
	var root1, root2 *Tree
	var val1 = []int{20, 35, 23, 40, 15, 30, 25, 50}
	for i := 0; i < len(val1); i++ {
		root1 = insertNode(root1, val1[i], struct{}{})
	}

	print("root1 avl树前序遍历(len=", len(val1), ")：\n")
	prePrintTree(root1)
	print("root1 没有23这个节点\n")

	print("--------------------分割线---------------------\n")

	var val2 = []int{20, 35, 40, 23, 15, 30, 25, 50}
	for i := 0; i < len(val2); i++ {
		root2 = insertNode(root2, val2[i], struct{}{})
	}
	print("root2 avl树前序遍历(len=", len(val2), ")：\n")
	prePrintTree(root2)
	print("root2 有23这个节点")
}
