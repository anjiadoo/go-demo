package skiplist

import (
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T) {
	sl := NewSkipList()

	total := 100
	for i := 1; i < total; i++ {
		sl.Insert(fmt.Sprintf("val_%d", i), total-i)
	}

	t.Log(sl)
	//fn(t, sl.head.forwards[0])

	//t.Log(len(sl.head.forwards))
	//sl.Insert("adfdf", 9090)

	for i := 0; i < len(sl.head.forwards); i++ {
		t.Log(sl.head.forwards[i])
	}

	//t.Log(sl.head.forwards[1])
	//t.Log(sl.head.forwards[2])
	//t.Log(sl.head.forwards[3])
	//t.Log(sl.head.forwards[4])
	//t.Log(sl.head.forwards[5])

	//t.Log(sl.Find("val_46" ,54))
}

func fn(t *testing.T, sl *skipListNode) {
	if sl == nil {
		return
	}

	t.Log(sl.v, sl.score, sl.level)
	fn(t, sl.forwards[0])
}
