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
