package main

import "fmt"

const Size = 10

func full(num, begin, size int, data [Size][Size]int) [Size][Size]int {
	if size == 0 {
		return data
	}
	if size == 1 {
		data[begin][begin] = num
		return data
	}

	i, j := begin, begin
	for k := 0; k < size-1; k++ {
		data[i][j] = num
		i++
		num++
	}
	for k := 0; k < size-1; k++ {
		data[i][j] = num
		j++
		num++
	}
	for k := 0; k < size-1; k++ {
		data[i][j] = num
		i--
		num++
	}
	for k := 0; k < size-1; k++ {
		data[i][j] = num
		j--
		num++
	}
	return full(num, begin+1, size-2, data)
}

/*
[1 36 35 34 33 32 31 30 29 28]
[2 37 64 63 62 61 60 59 58 27]
[3 38 65 84 83 82 81 80 57 26]
[4 39 66 85 96 95 94 79 56 25]
[5 40 67 86 97 100 93 78 55 24]
[6 41 68 87 98 99 92 77 54 23]
[7 42 69 88 89 90 91 76 53 22]
[8 43 70 71 72 73 74 75 52 21]
[9 44 45 46 47 48 49 50 51 20]
[10 11 12 13 14 15 16 17 18 19]
*/
func main() {
	var data = [Size][Size]int{}

	data = full(1, 0, Size, data)
	for _, row := range data {
		fmt.Println(row)
	}
}
