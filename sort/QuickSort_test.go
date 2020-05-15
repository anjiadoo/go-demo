package sort

import (
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := rand.Perm(10)
	t.Log(arr)
	QuickSort(arr)
	t.Log(arr)
}
