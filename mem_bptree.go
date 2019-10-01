package main

import (
	"go-demo/bptree"

	"fmt"
	"time"
)

/**
 * ----------------------------------------------------------------------------
 * -----------------------------[1,9,0,0,0]------------------------------------
 * ---------------------------------/-\----------------------------------------
 * -------------------------------/-----\--------------------------------------
 * ---------------------[1,5,0,0,]-----[9,13,17,0]-----------------------------
 * ----------------------/\---------------/|\----------------------------------
 * --------[1,2,3,4]->[5,6,7,8]---->[9,10,11,12]->[13,14,15,16]->[17,18,19,20]-
 * ----------------------------------------------------------------------------
 * ----------------------------------------------------------------------------
 */

func main() {
	var tree bptree.BPlusTree
	(&tree).Initialize()
	var i int = 1
	fmt.Println("st=", time.Now())
	for i <= 20 {
		_, result := tree.Insert(i, i*i)
		if result == false {
			fmt.Println("数据已存在")
		}
		i++
	}

	fmt.Println("et=", time.Now())
	//tree.Remove(9)
	//tree.Remove(10)
	//tree.Remove(11)
	//tree.Remove(12)
	//tree.Remove(13)

	//ptr := tree.Ptr
	//for ptr != nil {
	//	fmt.Println(ptr.KeyNum, ptr.Key /*, ptr.LeafNode.Datas*/)
	//	ptr = ptr.LeafNode.Next
	//}

	//find := 50
	//
	//start := time.Now()
	//resultDate, success := tree.FindData(find)
	//fmt.Println("dd:", time.Since(start))
	//
	//if success == true {
	//	fmt.Println("found key=", find, " val=", resultDate)
	//} else {
	//	fmt.Println("Not Found Key=", find)
	//}

	//遍历结点元素
	fmt.Printf("%+v\n", tree.Root)
	fmt.Println("---------------------------------------")
	fmt.Printf("%+v\n", tree.Root.Children[0])
	fmt.Printf("%+v\n", tree.Root.Children[0].Children[0])
	fmt.Printf("%+v\n", tree.Root.Children[0].Children[1])
	fmt.Println("---------------------------------------")
	fmt.Printf("%+v\n", tree.Root.Children[1])
	fmt.Printf("%+v\n", tree.Root.Children[1].Children[0])
	fmt.Printf("%+v\n", tree.Root.Children[1].Children[1])
	//fmt.Printf("%+v\n", tree.Root.Children[1].Children[2])
}
