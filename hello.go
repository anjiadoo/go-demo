package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b uint32
	_, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		return
	}
	if a < 1 || b > 200 || a > b {
		return
	}

	var sum uint32
	for i := a; i <= b; i++ {
		sum += uint32(math.Pow(float64(i), 3))
	}
	fmt.Println(sum)
}
