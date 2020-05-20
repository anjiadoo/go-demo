package sort

import (
	"math/rand"
	"testing"
)

func TestBucketSort(t *testing.T) {
	a := rand.Perm(100)
	t.Log(a)
	BucketSort(a)
	t.Log(a)
}

func TestBucketSortSimple(t *testing.T) {
	a := rand.Perm(100)
	t.Log(a)
	BucketSortSimple(a)
	t.Log(a)
}
