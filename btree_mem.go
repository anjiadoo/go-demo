package main

import (
	"flag"
	"fmt"
	"go-demo/btree"
	"sync"
	"time"
)

var (
	size   = flag.Int("size", 1000000, "size of the tree to build")
	degree = flag.Int("degree", 8, "degree of btree")
	gollrb = flag.Bool("llrb", false, "use llrb instead of btree")
)

//func main() {
//	flag.Parse()
//	vals := rand.Perm(*size)
//	var t, v interface{}
//	v = vals
//	var stats runtime.MemStats
//	for i := 0; i < 10; i++ {
//		runtime.GC()
//	}
//	fmt.Println("-------- BEFORE ----------")
//	runtime.ReadMemStats(&stats)
//	fmt.Printf("MemStats: %+v\n", stats)
//	start := time.Now()
//	/*if *gollrb {
//		tr := llrb.New()
//		for _, v := range vals {
//			tr.ReplaceOrInsert(llrb.Int(v))
//		}
//		t = tr // keep it around
//	} else */{
//		tr := btree.New(*degree)
//		for _, v := range vals {
//			tr.ReplaceOrInsert(btree.Int(v))
//		}
//		t = tr // keep it around
//	}
//	fmt.Printf("%v inserts in %v\n", *size, time.Since(start))
//	fmt.Println("-------- AFTER ----------")
//	runtime.ReadMemStats(&stats)
//	fmt.Printf("MemStats: %+v\n", stats)
//	for i := 0; i < 10; i++ {
//		runtime.GC()
//	}
//	fmt.Println("-------- AFTER GC ----------")
//	runtime.ReadMemStats(&stats)
//	fmt.Printf("MemStats: %+v\n", stats)
//	if t == v {
//		fmt.Println("to make sure vals and tree aren't GC'd")
//	}
//}

func main() {
	start := time.Now()
	// 并发写是不安全的
	tr := btree.New(*degree)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		for i := btree.Int(0); i < 1000; i++ {
			tr.ReplaceOrInsert(i)
		}
		wg.Done()
	}()

	go func() {
		for i := btree.Int(1000); i < 2000; i++ {
			tr.ReplaceOrInsert(i)
		}
		wg.Done()
	}()

	go func() {
		for i := btree.Int(2000); i < 3000; i++ {
			tr.ReplaceOrInsert(i)
		}
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("len:       ", tr.Len())
	fmt.Println("get3:      ", tr.Get(btree.Int(3)))
	fmt.Println("get100:    ", tr.Get(btree.Int(100)))
	fmt.Println("del4:      ", tr.Delete(btree.Int(4)))
	fmt.Println("del100:    ", tr.Delete(btree.Int(100)))
	fmt.Println("replace5:  ", tr.ReplaceOrInsert(btree.Int(5)))
	fmt.Println("replace100:", tr.ReplaceOrInsert(btree.Int(100)))
	fmt.Println("min:       ", tr.Min())
	fmt.Println("delmin:    ", tr.DeleteMin())
	fmt.Println("max:       ", tr.Max())
	fmt.Println("delmax:    ", tr.DeleteMax())
	fmt.Println("len:       ", tr.Len())
	fmt.Println("耗时:", time.Since(start))
	// Output:
	// len:        10
	// get3:       3
	// get100:     <nil>
	// del4:       4
	// del100:     <nil>
	// replace5:   5
	// replace100: <nil>
	// min:        0
	// delmin:     0
	// max:        100
	// delmax:     100
	// len:        8
}
