package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

func Init() {
	startTime = time.Now()
}

// 接收数据之后排序完再发送数据
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		// 读取数据到内存Read into memoty
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		fmt.Println("Read done:", time.Now().Sub(startTime))

		// 排序Sort
		sort.Ints(a)
		fmt.Println("InMenSort done:", time.Now().Sub(startTime))

		// 输出Output
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

// 归并
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 < v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("Merge done:", time.Now().Sub(startTime))
	}()
	return out
}

// 读取数据源
func ReaderSource(reader io.Reader, chunSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer)) // byte类型转换成int类型
				out <- v
			}
			if err != nil || (chunSize != -1 && bytesRead >= chunSize) { // chunSize=-1表示一直读
				break
			}
		}
		close(out)
	}()
	return out
}

// 写入数据文件
func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v)) // 写入byte类型的数据
		writer.Write(buffer)
	}
}

// 生成count个随机数作为数据源
func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

// 两两归并
func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	// merge inputs[0..m) and inputs [m..end)
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))
}
