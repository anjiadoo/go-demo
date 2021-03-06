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

func _mergeSort(arr []int, n int) {
	if n <= 1 {
		return
	}
	_mergeSort1(arr, 0, n-1)
}

func _mergeSort1(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := start + (end-start)>>1
	_mergeSort1(arr, start, mid)
	_mergeSort1(arr, mid+1, end)
	_merge(arr, start, mid, end)
}

func _merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)

	i, j, k := start, mid+1, 0
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
