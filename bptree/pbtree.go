package bptree

const M = 4
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX
const LIMIT_M_2 = (M + 1) / 2

type Position *BPlusFullNode

type BPlusLeafNode struct {
	Next  *BPlusFullNode
	Datas []interface{}
}

//叶子节点应该为Children为空，但leafNode中datas不为空 Next一般不为空
type BPlusFullNode struct {
	KeyNum   int
	Key      []int
	isLeaf   bool
	Children []*BPlusFullNode
	LeafNode *BPlusLeafNode
}

type BPlusTree struct {
	Root *BPlusFullNode
	Ptr  *BPlusFullNode
}

func mallocNewNode(isLeaf bool) *BPlusFullNode {
	var NewNode *BPlusFullNode
	if isLeaf == true {
		NewLeaf := mallocNewLeaf()
		NewNode = &BPlusFullNode{
			KeyNum:   0,
			Key:      make([]int, M+1), //申请M + 1是因为插入时可能暂时出现节点key大于M 的情况,待后期再分裂处理
			isLeaf:   isLeaf,
			Children: nil,
			LeafNode: NewLeaf,
		}
	} else {
		NewNode = &BPlusFullNode{
			KeyNum:   0,
			Key:      make([]int, M+1),
			isLeaf:   isLeaf,
			Children: make([]*BPlusFullNode, M+1),
			LeafNode: nil,
		}
	}
	for i, _ := range NewNode.Key {
		NewNode.Key[i] = INT_MIN
	}
	return NewNode
}

func mallocNewLeaf() *BPlusLeafNode {
	NewLeaf := BPlusLeafNode{
		Next:  nil,
		Datas: make([]interface{}, M+1),
	}
	for i, _ := range NewLeaf.Datas {
		NewLeaf.Datas[i] = i
	}
	return &NewLeaf
}

/* 初始化根结点 */
func (tree *BPlusTree) Initialize() {
	T := mallocNewNode(true)
	tree.Ptr = T
	tree.Root = T
}

func findMostLeft(P Position) Position {
	var Tmp Position
	Tmp = P
	if Tmp.isLeaf == true || Tmp == nil {
		return Tmp
	} else if Tmp.Children[0].isLeaf == true {
		return Tmp.Children[0]
	} else {
		for (Tmp != nil && Tmp.Children[0].isLeaf != true) {
			Tmp = Tmp.Children[0]
		}
	}
	return Tmp.Children[0]
}

func findMostRight(P Position) Position {
	var Tmp Position
	Tmp = P

	if Tmp.isLeaf == true || Tmp == nil {
		return Tmp
	} else if Tmp.Children[Tmp.KeyNum-1].isLeaf == true {
		return Tmp.Children[Tmp.KeyNum-1]
	} else {
		for (Tmp != nil && Tmp.Children[Tmp.KeyNum-1].isLeaf != true) {
			Tmp = Tmp.Children[Tmp.KeyNum-1]
		}
	}

	return Tmp.Children[Tmp.KeyNum-1]
}

/* 寻找一个兄弟节点，其存储的关键字未满，若左右都满返回nil */
func findSibling(Parent Position, posAtParent int) Position {
	var Sibling Position
	var upperLimit int
	upperLimit = M
	Sibling = nil
	if posAtParent == 0 {
		if Parent.Children[1].KeyNum < upperLimit {
			Sibling = Parent.Children[1]
		}
	} else if (Parent.Children[posAtParent-1].KeyNum < upperLimit) {
		Sibling = Parent.Children[posAtParent-1]
	} else if (posAtParent+1 < Parent.KeyNum && Parent.Children[posAtParent+1].KeyNum < upperLimit) {
		Sibling = Parent.Children[posAtParent+1]
	}
	return Sibling
}

/* 查找兄弟节点，其关键字数大于M/2 ;没有返回nil j用来标识是左兄还是右兄*/
func findSiblingKeyNum_M_2(Parent Position, i int, j *int) Position {
	var lowerLimit int
	var Sibling Position
	Sibling = nil

	lowerLimit = LIMIT_M_2

	if (i == 0) {
		if (Parent.Children[1].KeyNum > lowerLimit) {
			Sibling = Parent.Children[1]
			*j = 1
		}
	} else {
		if (Parent.Children[i-1].KeyNum > lowerLimit) {
			Sibling = Parent.Children[i-1]
			*j = i - 1
		} else if (i+1 < Parent.KeyNum && Parent.Children[i+1].KeyNum > lowerLimit) {
			Sibling = Parent.Children[i+1]
			*j = i + 1
		}

	}
	return Sibling
}

/**
 * 当要对beInsertedElement插入data的时候，posAtParent是beInsertedElement在Parent的位置，insertIndex是data要插入的位置，
 * j可由查找得到,当要对Parent插入beInsertedElement节点的时候，posAtParent是要插入的位置，Key和j的值没有用
 */
func (tree *BPlusTree) insertElement(isData bool, Parent Position, beInsertedElement Position, Key int, posAtParent int, insertIndex int, data interface{}) Position {

	var k int
	if (isData) {
		/* 插入data*/
		k = beInsertedElement.KeyNum - 1 //lastIdx
		for (k >= insertIndex) {
			beInsertedElement.Key[k+1] = beInsertedElement.Key[k]
			beInsertedElement.LeafNode.Datas[k+1] = beInsertedElement.LeafNode.Datas[k]
			k--
		}

		beInsertedElement.Key[insertIndex] = Key
		beInsertedElement.LeafNode.Datas[insertIndex] = data
		if (Parent != nil) {
			Parent.Key[posAtParent] = beInsertedElement.Key[0] //可能min_key已发生改变
		}

		beInsertedElement.KeyNum++

	} else {
		/* 插入节点，并对树的叶子节点进行链接 */
		if (beInsertedElement.isLeaf == true) {
			if (posAtParent > 0) {
				/* 链接前驱节点 */
				Parent.Children[posAtParent-1].LeafNode.Next = beInsertedElement
			}
			/* 链接后继节点 */
			beInsertedElement.LeafNode.Next = Parent.Children[posAtParent]

			if beInsertedElement.Key[0] <= tree.Ptr.Key[0] { //更新叶子指针
				tree.Ptr = beInsertedElement
			}
		}

		k = Parent.KeyNum - 1
		for (k >= posAtParent) { //插入节点时key也要对应的插入
			Parent.Children[k+1] = Parent.Children[k]
			Parent.Key[k+1] = Parent.Key[k]
			k--
		}
		Parent.Key[posAtParent] = beInsertedElement.Key[0]
		Parent.Children[posAtParent] = beInsertedElement
		Parent.KeyNum++
	}

	return beInsertedElement
}

/**
 * deleteIndex@需要删除的元素的下标
 * 两个参数X posAtParent 有些重复 posAtParent可以通过X的最小关键字查找得到
 */
func (tree *BPlusTree) removeElement(isData bool, Parent Position, beRemovedElement Position, posAtParent int, deleteIndex int) Position {
	var k, keyNum int

	if (isData) {
		keyNum = beRemovedElement.KeyNum
		/* 删除key */
		k = deleteIndex + 1
		for (k < keyNum) {
			beRemovedElement.Key[k-1] = beRemovedElement.Key[k]
			beRemovedElement.LeafNode.Datas[k-1] = beRemovedElement.LeafNode.Datas[k]
			k++
		}
		/* reset */
		beRemovedElement.Key[keyNum-1] = INT_MIN
		beRemovedElement.LeafNode.Datas[keyNum-1] = INT_MIN
		Parent.Key[posAtParent] = beRemovedElement.Key[0]
		beRemovedElement.KeyNum--
	} else {
		/* 删除节点，修改树叶节点的链接 */
		if (beRemovedElement.isLeaf == true && posAtParent > 0) {
			Parent.Children[posAtParent-1].LeafNode.Next = Parent.Children[posAtParent+1]
		}

		keyNum = Parent.KeyNum
		k = posAtParent + 1
		for (k < keyNum) {
			Parent.Children[k-1] = Parent.Children[k]
			Parent.Key[k-1] = Parent.Key[k]
			k++
		}

		if beRemovedElement.Key[0] == tree.Ptr.Key[0] { // refresh Ptr
			tree.Ptr = Parent.Children[0]
		}
		Parent.Children[Parent.KeyNum-1] = nil
		Parent.Key[Parent.KeyNum-1] = INT_MIN

		Parent.KeyNum--
	}
	return beRemovedElement
}

/**
 * Src和Dst是两个相邻的节点，posAtParent是Src在Parent中的位置；
 * 将Src的元素移动到Dst中 ,eNum是移动元素的个数
 */
func (tree *BPlusTree) moveElement(src Position, dst Position, parent Position, posAtParent int, eNum int) Position {
	var TmpKey int
	var data interface{}
	var Child Position
	var j int
	var srcInFront bool

	srcInFront = false

	if (src.Key[0] < dst.Key[0]) {
		srcInFront = true
	}
	j = 0
	/* 节点Src在Dst前面 */
	if (srcInFront) {
		if (src.isLeaf == false) {
			for (j < eNum) {
				Child = src.Children[src.KeyNum-1]
				tree.removeElement(false, src, Child, src.KeyNum-1, INT_MIN)        //每删除一个节点keyNum也自动减少1，队尾删
				tree.insertElement(false, dst, Child, INT_MIN, 0, INT_MIN, INT_MIN) //队头加
				j++
			}
		} else {
			for (j < eNum) {
				TmpKey = src.Key[src.KeyNum-1]
				data = src.LeafNode.Datas[src.KeyNum-1]
				tree.removeElement(true, parent, src, posAtParent, src.KeyNum-1)      //队尾删
				tree.insertElement(true, parent, dst, TmpKey, posAtParent+1, 0, data) //队头加
				j++
			}
		}

		parent.Key[posAtParent+1] = dst.Key[0]
		/* 将树叶节点重新连接 */
		if (src.KeyNum > 0) {
			findMostRight(src).LeafNode.Next = findMostLeft(dst) //似乎不需要重连，src的最右本身就是dst最左的上一元素
		} else {
			if src.isLeaf == true {
				if posAtParent == 0 {
					posAtParent = 1
				}
				parent.Children[posAtParent-1 ].LeafNode.Next = dst //todo:index out of range
			}
			//  此种情况肯定是merge merge中有实现先移动再删除操作
			//tree.removeElement(false ,parent.parent，parent ,parentIndex,INT_MIN )
		}
	} else {
		if (src.isLeaf == false) {
			for (j < eNum) {
				Child = src.Children[0]
				tree.removeElement(false, src, Child, 0, INT_MIN) //从src的队头删
				tree.insertElement(false, dst, Child, INT_MIN, dst.KeyNum, INT_MIN, INT_MIN)
				j++
			}

		} else {
			for (j < eNum) {
				TmpKey = src.Key[0]
				data = src.LeafNode.Datas[0]
				tree.removeElement(true, parent, src, posAtParent, 0)
				tree.insertElement(true, parent, dst, TmpKey, posAtParent-1, dst.KeyNum, data)
				j++
			}

		}

		parent.Key[posAtParent] = src.Key[0]
		if (src.KeyNum > 0) {
			findMostRight(dst).LeafNode.Next = findMostLeft(src)
		} else {
			if src.isLeaf == true {
				dst.LeafNode.Next = src.LeafNode.Next
			}
			//tree.removeElement(false ,parent.parent，parent ,parentIndex,INT_MIN )
		}
	}

	return parent
}

//posAtParent为节点beSplitedNode在Parent中的位置
func (tree *BPlusTree) splitNode(Parent Position, beSplitedNode Position, posAtParent int) Position {
	var j, k, keyNum int
	var NewNode Position

	if beSplitedNode.isLeaf == true {
		NewNode = mallocNewNode(true)
	} else {
		NewNode = mallocNewNode(false)
	}

	k = 0                        //刚申请节点从下标0开始接收data
	j = beSplitedNode.KeyNum / 2 //从分割节点这个位置开始导出data
	keyNum = beSplitedNode.KeyNum
	for (j < keyNum) {
		if (beSplitedNode.isLeaf == false) { //Internal node
			NewNode.Children[k] = beSplitedNode.Children[j]
			beSplitedNode.Children[j] = nil //reset
		} else {
			NewNode.LeafNode.Datas[k] = beSplitedNode.LeafNode.Datas[j]
			beSplitedNode.LeafNode.Datas[j] = INT_MIN
		}
		NewNode.Key[k] = beSplitedNode.Key[j]
		beSplitedNode.Key[j] = INT_MIN
		NewNode.KeyNum++
		beSplitedNode.KeyNum--
		j++
		k++
	}

	if (Parent != nil) {
		tree.insertElement(false, Parent, NewNode, INT_MIN, posAtParent+1, INT_MIN, INT_MIN)
	} else {
		/* 如果X是根，那么创建新的根并返回 */
		Parent = mallocNewNode(false)
		tree.insertElement(false, Parent, beSplitedNode, INT_MIN, 0, INT_MIN, INT_MIN)
		tree.insertElement(false, Parent, NewNode, INT_MIN, 1, INT_MIN, INT_MIN)
		tree.Root = Parent

		/**
		 * 为什么返回一个beSplitedNode一个Parent?
		 * 在方法外层，split后需要更新Parent的min_key，当树只有一层的时候，如果返回beSplitedNode节点更新根节点的min_key会出错
		 */
		return Parent
	}
	return beSplitedNode
}

/**
 * posAtParent@节点X在Parent中的位置
 * 合并节点,X少于M/2关键字，S有大于或等于M/2个关键字
 */
func (tree *BPlusTree) mergeNode(Parent Position, X Position, S Position, posAtParent int) Position {
	var Limit int

	/* S的关键字数目大于M/2 */
	if (S.KeyNum > LIMIT_M_2) {
		/* 从S中移动一个元素到X中 */
		tree.moveElement(S, X, Parent, posAtParent, 1)
	} else {
		/* 将X全部元素移动到S中，并把X删除 */
		Limit = X.KeyNum
		tree.moveElement(X, S, Parent, posAtParent, Limit) //最多时S恰好MAX MoveElement已考虑了parent.key的索引更新
		tree.removeElement(false, Parent, X, posAtParent, INT_MIN)
	}
	return Parent
}

/**
 * beInsertedElement@插入节点
 * Key@插入的key
 * posAtParent@父节点所在的位置
 * data@存储的数据
 */
func (tree *BPlusTree) recursiveInsert(beInsertedElement Position, Key int, posAtParent int, Parent Position, data interface{}) (Position, bool) {
	var InsertIndex, upperLimit int
	var Sibling Position
	var result = true

	/* 查找分支 */
	InsertIndex = 0
	for (InsertIndex < beInsertedElement.KeyNum && Key >= beInsertedElement.Key[InsertIndex]) {
		/* 重复值不插入 */
		if (Key == beInsertedElement.Key[InsertIndex]) {
			return beInsertedElement, false
		}
		InsertIndex++
	}
	//key必须大于被插入节点的最小元素且小于后继节点的最小元素，才能插入到此节点，故需回退一步
	if (InsertIndex != 0 && beInsertedElement.isLeaf == false) {
		InsertIndex--
	}

	if (beInsertedElement.isLeaf == true) {
		beInsertedElement = tree.insertElement(true, Parent, beInsertedElement, Key, posAtParent, InsertIndex, data) //返回叶子节点
	} else {
		//更新parent发生在split时
		beInsertedElement.Children[InsertIndex], result = tree.recursiveInsert(beInsertedElement.Children[InsertIndex], Key, InsertIndex, beInsertedElement, data)
	}

	/* 调整节点 */
	upperLimit = M
	if (beInsertedElement.KeyNum > upperLimit) {

		if (Parent == nil) { //根节点
			/* 分裂节点 */
			beInsertedElement = tree.splitNode(Parent, beInsertedElement, posAtParent)
		} else {
			Sibling = findSibling(Parent, posAtParent)
			if (Sibling != nil) {
				/* 将beInsertedElement的一个元素（Key或者Child）移动的Sibling中 */
				tree.moveElement(beInsertedElement, Sibling, Parent, posAtParent, 1)
			} else {
				/* 分裂节点 */
				beInsertedElement = tree.splitNode(Parent, beInsertedElement, posAtParent)
			}
		}
	}

	if (Parent != nil) {
		/* 插入节点的最小值可能已经变化，需要重新设置 */
		Parent.Key[posAtParent] = beInsertedElement.Key[0]
	}

	return beInsertedElement, result
}

/* 插入 */
func (tree *BPlusTree) Insert(Key int, data interface{}) (Position, bool) {
	//从根节点开始插入
	return tree.recursiveInsert(tree.Root, Key, 0, nil, data)
}

func (tree *BPlusTree) recursiveRemove(beRemovedElement Position, Key int, posAtParent int, Parent Position) (Position, bool) {
	var deleteIndex int
	var Sibling Position
	var NeedAdjust bool
	var result bool
	Sibling = nil

	/* 查找分支   TODO查找函数可以在参考这里的代码 或者实现一个递归遍历*/
	deleteIndex = 0
	for (deleteIndex < beRemovedElement.KeyNum && Key >= beRemovedElement.Key[deleteIndex]) {
		if (Key == beRemovedElement.Key[deleteIndex]) {
			break
		}
		deleteIndex++
	}

	if (beRemovedElement.isLeaf == true) {
		/* 没找到 */
		if (Key != beRemovedElement.Key[deleteIndex] || deleteIndex == beRemovedElement.KeyNum) {
			return beRemovedElement, false
		}
	} else {
		if (deleteIndex == beRemovedElement.KeyNum || Key < beRemovedElement.Key[deleteIndex]) {
			deleteIndex-- //准备到下层节点查找
		}
	}

	/* 树叶 */
	if (beRemovedElement.isLeaf == true) {
		beRemovedElement = tree.removeElement(true, Parent, beRemovedElement, posAtParent, deleteIndex)
	} else {
		beRemovedElement.Children[deleteIndex], result = tree.recursiveRemove(beRemovedElement.Children[deleteIndex], Key, deleteIndex, beRemovedElement)
	}

	NeedAdjust = false
	//有子节点的root节点，当keyNum小于2时
	if (Parent == nil && beRemovedElement.isLeaf == false && beRemovedElement.KeyNum < 2) {
		NeedAdjust = true
	} else if (Parent != nil && beRemovedElement.isLeaf == false && beRemovedElement.KeyNum < LIMIT_M_2) {
		/* 除根外，所有中间节点的儿子数不在[M/2]到M之间时。(符号[]表示向上取整) */
		NeedAdjust = true
	} else if (Parent != nil && beRemovedElement.isLeaf == true && beRemovedElement.KeyNum < LIMIT_M_2) {
		/* （非根）树叶中关键字的个数不在[M/2]到M之间时 */
		NeedAdjust = true
	}

	/* 调整节点 */
	if (NeedAdjust) {
		/* 根 */
		if (Parent == nil) {
			if (beRemovedElement.isLeaf == false && beRemovedElement.KeyNum < 2) {
				//树根的更新操作 树高度减一
				beRemovedElement = beRemovedElement.Children[0]
				tree.Root = beRemovedElement.Children[0]
				return beRemovedElement, true
			}

		} else {
			/* 查找兄弟节点，其关键字数目大于M/2 */
			Sibling = findSiblingKeyNum_M_2(Parent, posAtParent, &deleteIndex)
			if (Sibling != nil) {
				tree.moveElement(Sibling, beRemovedElement, Parent, deleteIndex, 1)
			} else {
				if (posAtParent == 0) {
					Sibling = Parent.Children[1]
				} else {
					Sibling = Parent.Children[posAtParent-1]
				}

				Parent = tree.mergeNode(Parent, beRemovedElement, Sibling, posAtParent)
				//Merge中已考虑空节点的删除
				beRemovedElement = Parent.Children[posAtParent]
			}
		}

	}

	return beRemovedElement, result
}

/* 删除 */
func (tree *BPlusTree) Remove(Key int) (Position, bool) {
	return tree.recursiveRemove(tree.Root, Key, 0, nil)
}

func (tree *BPlusTree) FindData(key int) (interface{}, bool) {
	var currentNode *BPlusFullNode
	var index int
	currentNode = tree.Root
	for index < currentNode.KeyNum {
		index = 0
		for key >= currentNode.Key[index] && index < currentNode.KeyNum {
			index++
		}
		if index == 0 {
			return INT_MIN, false
		} else {
			index--
			if currentNode.isLeaf == false {
				currentNode = currentNode.Children[index]
			} else {
				if key == currentNode.Key[index] {
					return currentNode.LeafNode.Datas[index], true
				} else {
					return INT_MIN, false
				}
			}
		}

	}
	return INT_MIN, false
}
