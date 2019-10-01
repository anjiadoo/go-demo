package main

import "fmt"

func getNext(tr string) []int {
	var i, j, leg int
	T := []rune(tr)
	var next = make([]int, len(T))

	next[0] = -1
	for j = 1; j < len(T); j++ {
		for leg = j - 1; leg >= 1; leg-- {
			for i = 0; i < leg; i++ {
				if T[i] != T[j-leg+i] {
					break
				}
			}
			if i == leg {
				next[j] = leg
				break
			}
		}
		if leg < 1 {
			next[j] = 0
		}
	}
	return next
}

func kmp(str, tr string) []int {
	var i, j int
	var ret []int

	s := []rune(str)
	t := []rune(tr)
	next := getNext(tr)
	for i = 0; i < len(s); {
		for j = 0; j < len(t) && i < len(s); {
			if s[i] == t[j] {
				i++
				j++
			} else {
				j = next[j]
				if j == -1 {
					i++
					j++
				}
			}
		}
		if j == len(t) {
			ret = append(ret, i-len(t))
		}
	}
	return ret
}

func main() {
	var T = "abc"
	var S = "abceeeabckkkabcsssabclllabc"

	fmt.Println(kmp(S, T))

	idx := kmp(S, T)

	fmt.Println(string([]rune(S)[idx[0] : idx[0]+3]))
	fmt.Println(string([]rune(S)[idx[1] : idx[1]+3]))
	fmt.Println(string([]rune(S)[idx[2] : idx[2]+3]))
	fmt.Println(string([]rune(S)[idx[3] : idx[3]+3]))
	fmt.Println(string([]rune(S)[idx[4] : idx[4]+3]))
}
