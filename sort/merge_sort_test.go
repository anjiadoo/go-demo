package sort

import (
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := rand.Perm(10)
	t.Log(arr)
	MergeSort(arr)
	t.Log(arr)
}
