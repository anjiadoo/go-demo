package jz_offer

//输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

func exchange(nums []int) []int {
	var flag int
	for k := 0; k < len(nums); k++ {
		if nums[k]%2 == 0 {
			continue
		}
		if flag != k {
			nums[flag], nums[k] = nums[k], nums[flag]
		}
		flag++
	}
	return nums
}

func exchange1(nums []int) []int {
	head, tail := 0, len(nums)-1

	for head < tail {
		// 偶数-奇数
		if nums[head]%2 == 0 && nums[tail]%2 != 0 {
			nums[head], nums[tail] = nums[tail], nums[head]
			head++
			tail--
		} else {
			for nums[head]%2 != 0 && head < len(nums)-1 {
				head++
			}
			for nums[tail]%2 == 0 && tail > 0 {
				tail--
			}
		}
	}
	return nums
}
