package sort

// QuickSort is quicksort methods for golang
func QuickSort(arr []int) {
	separateSort(arr, 0, len(arr)-1)
}

func separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	separateSort(arr, start, i-1)
	separateSort(arr, i+1, end)
}

func partition(arr []int, start, end int) int {
	// 选取最后一位当对比数字
	pivot := arr[end]

	// point 分区点
	var point = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(point == j) {
				// 交换位置
				arr[point], arr[j] = arr[j], arr[point]
			}
			point++
		}
	}

	arr[point], arr[end] = arr[end], arr[point]
	return point
}

func quickSort(arr []int) {
	_separateSort(arr, 0, len(arr)-1)
}

func _separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := _partition(arr, start, end)
	_separateSort(arr, start, mid-1)
	_separateSort(arr, mid+1, end)
}

func _partition(arr []int, start, end int) int {
	pivot := arr[end]

	var point = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if point != j {
				arr[j], arr[point] = arr[point], arr[j]
			}
			point++
		}
	}
	arr[point], arr[end] = arr[end], arr[point]
	return point
}
