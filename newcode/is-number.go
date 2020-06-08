package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		str := scanner.Text()
		if len(str) == 0 {
			break
		}
		strArr := []rune(str)

		var rec [26][]rune
		for _, v := range strArr {
			if unicode.IsLetter(v) {
				rec[unicode.ToLower(v)-'a'] = append(rec[unicode.ToLower(v)-'a'], v)
			}
		}

		var tmpRuneArr []rune
		for i := 0; i < 26; i++ {
			if len(rec[i]) != 0 {
				for _, v := range rec[i] {
					tmpRuneArr = append(tmpRuneArr, v)
				}
			}
		}

		index := 0
		if tmpRuneArr != nil {
			for i, v := range str {
				if unicode.IsLetter(v) {
					strArr[i] = tmpRuneArr[index]
					index++
					if index > len(tmpRuneArr) {
						break
					}
				}
			}
		}

		fmt.Println(string(strArr))
	}
}
