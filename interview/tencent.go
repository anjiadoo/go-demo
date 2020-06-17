package interview

import (
	"fmt"
	"sort"
)

// 腾讯面委题目1
// 不允许开辟堆栈去掉字符串里的空格

func _foo(str string) string {
	strByte := []byte(str)
	for i := 0; i < len(strByte); i++ {
		if strByte[i] == ' ' {
			strByte = append(strByte[:i], strByte[i+1:]...)
			i--
		}
	}
	return string(strByte)
}

// 腾讯面委题目2
// 求一个字符串的所有子集，如：abc -> [a,b,c,ab,ac,bc,abc]
func _foo1(str string) []string {
	// str := a
	// str := a b
	// str := ab c
	// str := abc d
	ret := []string{}

	for i := 0; i < len(str); i++ {
		tmp := []string{}
		for j := 0; j < len(ret); j++ {
			tmp = append(tmp, fmt.Sprintf("%s%s", ret[j], string(str[i])))
		}

		ret = append(ret, tmp...)
		ret = append(ret, string(str[i]))
	}
	sort.Strings(ret)
	return ret
}
