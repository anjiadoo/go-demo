package main

import (
	"fmt"
	"go-demo/bptree"
	"time"
)

func main() {
	var tree bptree.BPlusTree
	(&tree).Initialize()
	var i int = 1
	for i <= 10000000 {
		_, result := tree.Insert(i, i*i)
		//fmt.Println(i)
		if result == false {
			fmt.Println("数据已存在")
		}
		i++
	}

	tree.Remove(7)
	tree.Remove(6)
	tree.Remove(5)
	find := 50

	start := time.Now()
	resultDate, success := tree.FindData(find)
	fmt.Println("dd:", time.Since(start))

	if success == true {
		fmt.Println("found key=", find, " val=", resultDate)
	} else {
		fmt.Println("Not Found Key=", find)
	}

	//遍历结点元素
	fmt.Printf("1 %+v\n", tree.Root)
	fmt.Printf("2 %+v\n", tree.Root.Children[0])
	fmt.Printf("3 %+v\n", tree.Root.Children[0].Children[0])
	fmt.Printf("3 %+v\n", tree.Root.Children[0].Children[0].LeafNode)
}
