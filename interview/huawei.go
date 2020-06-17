package interview

func palindrome(str string) bool {
	if len(str) == 0 || len(str) == 1 {
		return true
	}

	i := 0
	j := len(str) - 1

	for {
		if i >= j {
			break
		}
		if str[i] == str[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func palindrome1(str string, i, j int) bool {
	if i >= j {
		return true
	}

	if str[i] == str[j] {
		start := i + 1
		end := j - 1
		return palindrome1(str, start, end)
	}
	return false
}
