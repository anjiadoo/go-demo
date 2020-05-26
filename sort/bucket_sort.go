package sort

// 桶排序

// 获取待排序数组中的最大值
func getMaxElement(array []int) int {
	maxElement := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > maxElement {
			maxElement = array[i]
		}
	}
	return maxElement
}

func BucketSort(array []int) {
	length := len(array)
	if length <= 1 {
		return
	}

	maxElement := getMaxElement(array)
	buckets := make([][]int, length) // 二维切片

	index := 0
	for i := 0; i < length; i++ {
		index = array[i] * (length - 1) / maxElement      // 桶序号
		buckets[index] = append(buckets[index], array[i]) // 加入对应的桶中
	}

	tmpPos := 0 // 标记数组位置
	for i := 0; i < length; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			QuickSort(buckets[i]) // 桶内做快速排序
			copy(array[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
}

// 桶排序简单实现
func BucketSortSimple(source []int) {
	if len(source) <= 1 {
		return
	}

	array := make([]int, getMaxElement(source)+1)
	for i := 0; i < len(source); i++ {
		array[source[i]] ++
	}

	arr := make([]int, 0)
	for i := 0; i < len(array); i++ {
		for array[i] != 0 {
			arr = append(arr, i)
			array[i] --
		}
	}
	copy(source, arr)
}
