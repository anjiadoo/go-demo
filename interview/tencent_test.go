package interview

import "testing"

func Test_foo(t *testing.T) {
	str := "123   123   123   "
	t.Log(len(_foo(str)))
	t.Log((_foo(str)))
	//
	//str1 := ""
	//t.Log(len(_foo(str1)))
	//
	//str2 := "  ab  c  "
	//t.Log(len(_foo(str2)))
}

func Test_foo1(t *testing.T) {
	str := "12345"
	t.Log(len(_foo1(str)))
	t.Log((_foo1(str)))

	//str1 := ""
	//t.Log(len(_foo1(str1)))
	//t.Log((_foo1(str1)))
	//
	//str2 := "aaa"
	//t.Log(len(_foo1(str2)))
	//t.Log((_foo1(str2)))
}
