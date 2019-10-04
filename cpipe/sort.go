package cpipe

import (
	"./pipeline"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/******************生成测试数据*****************
	const filename = "large512.in"
	const n = 67108864

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(n) // 获取随机数据源

	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush() // 倒出bufio里面的全部数据
	************************************************/
	p := createPipeline("small01.in", 1024, 4)
	writeToFile(p, "small01.out")
	printFile("small01.out")

	// p := createNetworkPipeline("large512.in", 536870912, 4)
	// writeToFile(p, "large512.out")
	// printFile("large512.out")
}

// 创建一个Pipeline
func createPipeline(filename string, fileSize, chunCount int) <-chan int {
	chunkSize := fileSize / chunCount
	pipeline.Init()

	sortResults := []<-chan int{}
	for i := 0; i < chunCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0) // 从每一块的开始位置读取

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, pipeline.InMemSort(source))
	}
	return pipeline.MergeN(sortResults...)
}

// 写文件
func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}

// 打印文件
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

// 创建一个网络版Pipeline
func createNetworkPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	pipeline.Init()

	sortAddr := []string{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(9000+i)
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))

		sortAddr = append(sortAddr, addr)
	}

	sortResults := []<-chan int{}
	for _, addr := range sortAddr {
		sortResults = append(sortResults, pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(sortResults...)
}
