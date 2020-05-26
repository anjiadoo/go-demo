package sort

import (
	"math/rand"
	"testing"
)

func TestCountingSort(t *testing.T) {
	arr := rand.Perm(10)
	t.Log(arr)
	CountingSort(arr)
	t.Log(arr)

	t.Log(" ------------------- ")

	arr = rand.Perm(10)
	t.Log(arr)
	countingSort(arr)
	t.Log(arr)
}
