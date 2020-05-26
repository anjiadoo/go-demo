package sort

/*
冒泡排序、插入排序、选择排序
*/

//冒泡排序，a是数组，n表示数组大小
func BubbleSort(arr []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		// 提前退出标志
		flag := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				//此次冒泡有数据交换
				flag = true
			}
		}
		// 如果没有交换数据，提前退出
		if !flag {
			break
		}
	}
}

func bubbleSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	for i := 0; i < length; i++ {
		// 提前退出标志
		flag := false
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// 此次冒泡有数据交换
				flag = true
			}
		}
		// 没有数据交换，提前退出
		if !flag {
			break
		}
	}
}

// 插入排序，a表示数组，n表示数组大小
func InsertionSort(arr []int, n int) {
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		value := arr[i]
		j := i - 1
		//查找要插入的位置并移动数据
		for ; j >= 0; j-- {
			if arr[j] > value {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = value
	}
}

func insertionSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	for i := 1; i < length; i++ {
		val := arr[i]
		j := i - 1
		// 查找要插入的位置
		for ; j >= 0; j-- {
			if arr[j] > val {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = val
	}
}

// 选择排序，a表示数组，n表示数组大小
func SelectionSort(arr []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		// 查找最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 交换
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func selectionSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	for i := 0; i < length; i++ {
		// 查找最小值
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 交换
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
