package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 5, 9, 6, 3, 7, 5, 10}
	fmt.Println("排序前：", arr)
	BubbleSort(arr, len(arr))
	fmt.Println("排序后：", arr)
}

func TestBubbleSort_copy(t *testing.T) {
	arr := []int{1, 5, 9, 6, 3, 7, 5, 10}
	fmt.Println("排序前：", arr)
	bubbleSort(arr)
	fmt.Println("排序后：", arr)
}

func TestInsertionSort(t *testing.T) {
	arr := []int{1, 5, 9, 6, 3, 7, 5, 10}
	fmt.Println("排序前：", arr)
	InsertionSort(arr, len(arr))
	fmt.Println("排序后：", arr)
}

func TestInsertionSort_copy(t *testing.T) {
	arr := []int{1, 5, 9, 6, 3, 7, 5, 10}
	fmt.Println("排序前：", arr)
	insertionSort(arr)
	fmt.Println("排序后：", arr)
}

func TestSelectionSort(t *testing.T) {
	arr := []int{1, 5, 9, 6, 3, 7, 5, 10}
	fmt.Println("排序前：", arr)
	SelectionSort(arr, len(arr))
	fmt.Println("排序后：", arr)
}

func TestSelectionSort_copy(t *testing.T) {
	arr := []int{1, 5, 9, 6, 3, 7, 5, 10}
	fmt.Println("排序前：", arr)
	selectionSort(arr)
	fmt.Println("排序后：", arr)
}
