package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

type MyStruct struct {
	i int
	j int
}

func myFunction1(ms *MyStruct) {
	ptr := unsafe.Pointer(ms)
	for i := 0; i < 2; i++ {
		c := (*int)(unsafe.Pointer((uintptr(ptr) + uintptr(8*i))))
		*c += i + 1
		fmt.Printf("[%p] %d\n", c, *c)
	}
}

func main121() {
	a := &MyStruct{i: 40, j: 50}
	myFunction1(a)
	fmt.Printf("[%p] %v\n", a, a)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type user struct {
	name string
	addr *addr
}

type addr struct {
	province string
	city     string
	num      int
}

func (a *addr) setAttr() {
	a.num += a.num
}

func foo(a1 *[]int) {
	//fmt.Printf("%p \n", aa)
	*a1 = append(*a1, 6)
	fmt.Printf("%p \n", a1)
}

func main() {
	aa := []int{1, 2, 3, 4, 5}
	bb := &aa
	cc := 1

	fmt.Printf("%p \n", aa)
	fmt.Printf("%p \n", &bb)
	foo(&aa)
	fmt.Printf("%p", &cc)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func IsSameDay1(timestamp1, timestamp2 int64) bool {
	t1 := time.Unix(timestamp1, 0).Format("2006-01-02 00:00:00")
	t2 := time.Unix(timestamp2, 0).Format("2006-01-02 00:00:00")
	return t1 == t2
}

func main1111111() {
	now := time.Now().Unix()
	fmt.Println(IsSameDay1(0, now))
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
const (
	TaskCycle_WEEK  = 1
	TaskCycle_MONTH = 2
	TaskCycle_DAY   = 3
)

//CalcCycleCursor 计算当前任务周期游标，游标号段/命名空间方式合成
//空间分布: 类型，年， 值
func CalcCycleCursor(TaskCycle int32) int {
	switch TaskCycle {
	case TaskCycle_WEEK:
		return calcWeekCursor()
	case TaskCycle_MONTH:
		return calcMonthlyCursor()
	case TaskCycle_DAY:
		return calcDayCursor()
	default:
		return 0
	}
}

func calcWeekCursor() int {
	now := time.Now()
	year, week := now.ISOWeek()
	strCursor := fmt.Sprintf("%d%04d%03d", TaskCycle_WEEK, year, week)
	cursor, err := strconv.Atoi(strCursor)
	if err != nil {
		log.Fatal(err.Error())
	}
	return cursor
}

func calcMonthlyCursor() int {
	now := time.Now()
	strCursor := fmt.Sprintf("%d%04d%03d", TaskCycle_MONTH, now.Year(), now.Month())
	cursor, err := strconv.Atoi(strCursor)
	if err != nil {
		log.Fatal(err.Error())
	}
	return cursor
}

func calcDayCursor() int {
	now := time.Now()
	strCursor := fmt.Sprintf("%d%04d%03d", TaskCycle_DAY, now.Year(), 11)
	cursor, err := strconv.Atoi(strCursor)
	if err != nil {
		log.Fatal(err.Error())
	}
	return cursor
}

func main111111() {
	//fmt.Println(CalcCycleCursor(TaskCycle_WEEK))
	//fmt.Println(CalcCycleCursor(TaskCycle_MONTH))
	fmt.Println(CalcCycleCursor(TaskCycle_DAY))
}

func IsSameDay(timestamp1, timestamp2 int64) bool {
	return getDayStr(timestamp1) == getDayStr(timestamp2)
}

func getDayStr(timestamp int64) string {
	fmt.Println(time.Unix(timestamp, 0).Format("2006-01-02 00:00:00"))
	return time.Unix(timestamp, 0).Format("2006-01-02 00:00:00")
}

func main1() {
	fmt.Println(IsSameDay(1577775700, 1577787700))
}

func byteToString(b []byte) string {
	return string(b)
}

func byteToStringNoAlloc(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	sh := reflect.StringHeader{uintptr(unsafe.Pointer(&b[0])), len(b)}
	return *(*string)(unsafe.Pointer(&sh))
}

func main11111() {
	b := []byte("anjiadoo")
	fmt.Println("切片第一个元素: ", spew.Sdump(&b[0]))

	str := byteToString(b)
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	fmt.Println("分配内存的方式: ", spew.Sdump(sh))

	strNoAlloc := byteToStringNoAlloc(b)
	shNoAlloc := (*reflect.StringHeader)(unsafe.Pointer(&strNoAlloc))
	fmt.Println("不分配内存的方式: ", spew.Sdump(shNoAlloc))
}

func main111() {
	var bb = []byte("anjiadooo")

	tt := time.Now()
	println(string(bb))
	println(time.Since(tt))

	ti := time.Now()
	sh := reflect.StringHeader{
		uintptr(unsafe.Pointer(&bb[0])),
		len(bb),
	}
	println(*(*string)(unsafe.Pointer(&sh)))
	println(time.Since(ti))
}
