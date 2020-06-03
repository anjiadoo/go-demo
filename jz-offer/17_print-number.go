package jz_offer

import "fmt"

func printNumbers(n int) {
	if n == 0 {
		return
	}
	arr := make([]byte, n)
	for i := range arr {
		arr[i] = '0'
	}

	for !isFinished(arr) {
		print1(arr)
	}
}

func isFinished(arr []byte) bool {
	var overNum int
	n := len(arr)

	// 循环进位处理
	for i := n - 1; i >= 0; i-- {
		num := int(arr[i]-'0') + overNum
		if i == n-1 {
			num++
		}

		if num >= 10 {
			if i == 0 {
				return true
			} else {
				overNum = 1
				arr[i] = '0'
			}
		} else {
			arr[i] = byte(num + '0')
			break
		}
	}
	return false
}

func print1(arr []byte) {
	var flag bool
	for i := 0; i < len(arr); i++ {
		if arr[i] != '0' {
			flag = true
		}
		if flag {
			fmt.Printf("%s\n", arr[i:])
			break
		}
	}
}
