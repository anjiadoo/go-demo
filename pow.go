package main

import "fmt"

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	t := pow(x, n/2)
	if n%2 == 1 {
		return x * t * t
	}
	return t * t
}

func fn(n int) int {
	if n == 0 {
		return 1
	}
	return fn(n-1) + fn(n-1)
}

func main() {
	var x float64 = 2
	var n int = 10

	fmt.Println(pow(x, n))
	fmt.Println(fn(n))
}
