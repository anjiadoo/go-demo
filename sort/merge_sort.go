package sort

func MergeSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	mergeSort(arr, 0, length-1)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func mergeSort_copy(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort_copy(arr, start, mid)
	mergeSort_copy(arr, mid+1, end)
	merge_copy(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] <= arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		k++
	}
	copy(arr[start:end+1], tmpArr)
}

func merge_copy(arr []int, start, mid, end int) {
	tmp := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
	}

	for ; i <= mid; k++ {
		tmp[k] = arr[i]
		i++
	}
	for ; j <= end; k++ {
		tmp[k] = arr[j]
		j++
	}
	copy(arr[start:end+1], tmp)
}
