package sort

func CountingSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	max := 0
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}

	tmp := make([]int, max+1)
	for i := range arr {
		tmp[arr[i]]++
	}
	for i := 1; i <= max; i++ {
		tmp[i] += tmp[i-1]
	}

	ret := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		index := tmp[arr[i]] - 1
		ret[index] = arr[i]
		tmp[arr[i]]--
	}
	copy(arr, ret)
}

func countingSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	max := 0
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}

	tmp := make([]int, max+1)
	for i := range arr {
		tmp[arr[i]]++
	}
	// 顺序求和
	for i := 1; i <= max; i++ {
		tmp[i] += tmp[i-1]
	}

	ret := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		idx := tmp[arr[i]] - 1
		ret[idx] = arr[i]
		tmp[arr[i]]--
	}
	copy(arr, ret)
}
