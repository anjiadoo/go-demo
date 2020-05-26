package sort

// QuickSort is quicksort methods for golang
func QuickSort(arr []int) {
	separateSort(arr, 0, len(arr)-1)
}

func separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	//i := partition(arr, start, end)
	i := partition_copy(arr, start, end)
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

func partition_copy(arr []int, start, end int) int {
	// 选取最后一个元素当对比数字
	pivot := arr[end]

	// point 分区点下标
	point := start
	for i := start; i < end; i++ {
		if arr[i] < pivot {
			if point != i {
				// 交换位置
				arr[point], arr[i] = arr[i], arr[point]
			}
			point++
		}
	}

	// 把所选的元素放到中间
	arr[point], arr[end] = arr[end], arr[point]
	return point
}
