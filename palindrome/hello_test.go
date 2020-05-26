package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	str := "abccba"
	t.Log(palindrome1(str,0,len(str)-1))

	str1 := ""
	t.Log(palindrome1(str1,0,len(str1)-1))

	str2 := "1"
	t.Log(palindrome1(str2,0,len(str2)-1))

	str3 := "1212"
	t.Log(palindrome1(str3,0,len(str3)-1))

	str4 := "121"
	t.Log(palindrome1(str4,0,len(str4)-1))
}
