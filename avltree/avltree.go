package avltree

type Tree struct {
	parent *Tree
	lchild *Tree
	rchild *Tree
	depth  int
	val    int
	data   interface{}
}

// LL型，在node节点的左子树根节点的左子树上插入节点而破坏平衡->右旋
func llRotate(node *Tree) *Tree {
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
func rrRotate(node *Tree) *Tree {
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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//LR型，在node节点的左子树根节点的右子树上插入节点而破坏平衡->先左旋再右旋
func lrRotate(node *Tree) *Tree {
	node.lchild = rrRotate(node.lchild)
	return llRotate(node)
}

//RL型，在node节点的右子树根节点的左子树上插入节点而破坏平衡->先右旋再左旋
func rlRotate(node *Tree) *Tree {
	node.rchild = llRotate(node.rchild)
	return rrRotate(node)
}

func update_depth(node *Tree) {
	if node == nil {
		return
	}
	node.depth = max(get_balance(node.lchild), get_balance(node.rchild)) + 1
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

func InsertNode(node *Tree, val int, data interface{}) *Tree {
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
		node.lchild = InsertNode(node.lchild, val, data)
		if is_balance(node) == 2 {
			if val < node.lchild.val {
				node = llRotate(node)
			} else {
				node = lrRotate(node)
			}
		}
	} else if val > node.val {
		node.rchild = InsertNode(node.rchild, val, data)
		if is_balance(node) == -2 {
			if val > node.rchild.val {
				node = rrRotate(node)
			} else {
				node = rlRotate(node)
			}
		}
	}
	// 更新depth, 可能插入没有旋转
	update_depth(node)
	return node
}

func PrePrintTree(node *Tree) {
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
	PrePrintTree(node.lchild)
	PrePrintTree(node.rchild)
}

//------------------------------------
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
