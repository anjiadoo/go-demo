package main

import (
	"fmt"
	"time"
)

func FibonacciSequence(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return FibonacciSequence(n-1) + FibonacciSequence(n-2)
}

func dpm(memo []int, n int) int {

	if n > 0 && memo[n] == 0 {
		memo[n] = dpm(memo, n-1) + dpm(memo, n-2)
	}
	return memo[n]
}

func dynamicProgrammingMemo(n int) int {
	if n < 1 {
		return 0
	}

	memo := make([]int, n+1)
	memo[1], memo[2] = 1, 1

	return dpm(memo, n)
}

func dynamicProgramming(n int) int {
	if n < 2 {
		return n
	}
	sum, prev, curr := 0, 0, 1
	for i := 0; i < n-1; i++ {
		sum = prev + curr
		prev = curr
		curr = sum
	}
	return sum
}

func main() {
	var num = 500

	fmt.Println("------------|---------------------------|-------------------------------")
	fmt.Println("     seq    |        fibonacci          |         running time          ")
	fmt.Println("------------|---------------------------|-------------------------------")

	//st1 := time.Now()
	//sum1 := dynamicProgrammingMemo(num)
	//et1 := time.Since(st1)
	//fmt.Println("\t", num, "\t|\t", sum1, "\t\t\t|\t\t", et1)

	st3 := time.Now()
	sum3 := dynamicProgramming(num)
	et3 := time.Since(st3)
	fmt.Println("\t", num, "\t|\t", sum3, "\t\t|\t\t", et3)

	//st2 := time.Now()
	//sum2 := FibonacciSequence(num)
	//et2 := time.Since(st2)
	//fmt.Println("\t", num, "\t|\t\t", sum2, "\t\t\t|\t\t", et2)

	fmt.Println("------------|---------------------------|-------------------------------")
}
