package sort

import (
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := rand.Perm(10)
	t.Log(arr)
	_mergeSort(arr, len(arr))
	t.Log(arr)
}
