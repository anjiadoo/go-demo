package rbtree

import "fmt"

const (
	RED   = true
	BLACK = false
)

var NIL = &node{nil, nil, nil, BLACK, nil}

type node struct {
	parent *node
	lchild *node
	rchild *node
	color  bool
	Item
}

type rbTree struct {
	nilNode *node
	root    *node
	count   uint64
}

func New() *rbTree {
	node := node{nil, nil, nil, BLACK, nil}
	return &rbTree{
		nilNode: &node,
		root:    &node,
		count:   0,
	}
}

// 左旋
func (rbt *rbTree) leftRotate(node *node) {
	if node.rchild == rbt.nilNode {
		return
	}

	rchild := node.rchild
	node.rchild = rchild.lchild

	if rchild.lchild != rbt.nilNode {
		node.lchild.parent = node
	}
	rchild.parent = node.parent

	if node.parent == rbt.nilNode {
		rbt.root = rchild
	} else if node == node.parent.lchild {
		node.parent.lchild = rchild
	} else {
		node.parent.rchild = rchild
	}

	rchild.lchild = node
	node.parent = rchild
}

// 右旋
func (rbt *rbTree) rightRotate(node *node) {
	if node.lchild == rbt.nilNode {
		return
	}

	lchild := node.lchild
	node.lchild = lchild.rchild

	if lchild.rchild != rbt.nilNode {
		lchild.rchild.parent = node
	}
	lchild.parent = node.parent

	if node.parent == rbt.nilNode {
		rbt.root = lchild
	} else if node == node.parent.lchild {
		node.parent.lchild = lchild
	} else {
		node.parent.rchild = lchild
	}

	lchild.rchild = node
	node.parent = lchild
}

// 替换或插入
func (rbt *rbTree) ReplaceOrInsert(val Item) {
	node := &node{
		Item:   val,
		color:  RED,
		parent: rbt.nilNode,
		rchild: rbt.nilNode,
		lchild: rbt.nilNode,
	}

	root := rbt.root
	tmp := rbt.nilNode
	for root != rbt.nilNode {
		tmp = root
		if less(node.Item, root.Item) {
			root = root.lchild
		} else if less(root.Item, node.Item) {
			root = root.rchild
		} else {
			root.Item = node.Item
		}
	}

	node.parent = tmp
	if tmp == rbt.nilNode {
		rbt.root = node
	} else if less(node.Item, tmp.Item) {
		tmp.lchild = node
	} else {
		tmp.rchild = node
	}

	rbt.count++
	rbt.insertFixup(node)
}

func (rbt *rbTree) insertFixup(node *node) {
	// 父节点为红色，如果为黑色直接插入不影响树的平衡
	for node.parent.color == RED {
		// node节点的父节点为祖父节点的左孩子的情况
		if node.parent == node.parent.parent.lchild {
			uncle := node.parent.parent.rchild
			if uncle.color == RED {
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent //循环向上自平衡
			} else { // 叔叔节点为黑色需要先左旋
				if node == node.parent.rchild {
					node = node.parent
					rbt.leftRotate(node)
				}
				node.parent.color = BLACK
				node.parent.parent.color = RED
				rbt.rightRotate(node.parent.parent)
			}
		} else { // node节点的父节点为祖父节点的右孩子的情况
			uncle := node.parent.parent.lchild
			if uncle.color == RED {
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent //循环向上自平衡
			} else {
				if node == node.parent.lchild {
					node = node.parent
					rbt.rightRotate(node)
				}
				node.parent.color = BLACK
				node.parent.parent.color = RED
				rbt.leftRotate(node.parent.parent)
			}
		}
	}
	rbt.root.color = BLACK
}

//----------------------------------------------
// 前序遍历
func PreOrder(node *node) {
	if node == nil {
		return
	}
	var RBM = map[bool]string{true: "Red", false: "Black"}

	if *node != *NIL {
		fmt.Printf("node.val=%v color:%v\n", node.Item, RBM[node.color])
	}
	if node.lchild != nil {
		fmt.Printf("\tlchild.val=%v color:%v\n", node.lchild.Item, RBM[node.lchild.color])
	}
	if node.rchild != nil {
		fmt.Printf("\trchild.val=%v color:%v\n", node.rchild.Item, RBM[node.rchild.color])
	}

	PreOrder(node.lchild)
	PreOrder(node.rchild)
}

// 层序遍历
func LeverOrder(root *node) {
	var front, rear = -1, -1
	var queue []*node
	var RBM = map[bool]string{true: "Red", false: "Black"}

	if *root == *NIL {
		return
	}

	rear++
	queue = append(queue, root)
	for front != rear {
		front++
		ele := queue[front]

		if *ele == *NIL {
			continue
		}
		fmt.Printf("node.val=%v, %v\t", ele.Item, RBM[ele.color], )

		if ele.lchild != nil {
			fmt.Printf("lchild=%v, %v\t", ele.lchild.Item, RBM[ele.lchild.color])
		}

		if ele.rchild != nil {
			fmt.Printf("rchild=%v, %v\t", ele.rchild.Item, RBM[ele.rchild.color])
		}

		if ele.lchild != nil {
			rear++
			queue = append(queue, ele.lchild)
		}
		if ele.rchild != nil {
			rear++
			queue = append(queue, ele.rchild)
		}
		fmt.Println()
	}
}

func (rbt *rbTree) GetNodeCnt() uint64 {
	return rbt.count
}

func (rbt *rbTree) GetRootNode() *node {
	return rbt.root
}

//----------------------------------------------
type Item interface {
	Less(than Item) bool
}

type Int int

func (x Int) Less(than Item) bool {
	return x < than.(Int)
}

type Uint32 uint32

func (x Uint32) Less(than Item) bool {
	return x < than.(Uint32)
}

func less(root, tmp Item) bool {
	return root.Less(tmp)
}
