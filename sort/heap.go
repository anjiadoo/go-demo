package sort

type Heap struct {
	arr   []int
	cap   int
	count int
}

//init heap
func NewHeap(capacity int) *Heap {
	heap := &Heap{}
	heap.cap = capacity
	heap.arr = make([]int, capacity+1)
	heap.count = 0
	return heap
}

//top-max heap -> heapify from down to up
func (heap *Heap) insert(data int) {
	//defensive
	if heap.count == heap.cap {
		return
	}

	heap.count++
	heap.arr[heap.count] = data

	//compare with parent node
	i := heap.count
	parent := i / 2
	for parent > 0 && heap.arr[parent] < heap.arr[i] {
		heap.arr[parent], heap.arr[i] = heap.arr[i], heap.arr[parent]
		i = parent
		parent = i / 2
	}
}

//heapfify from up to down
func (heap *Heap) removeMax() {

	//defensive
	if heap.count == 0 {
		return
	}

	//swap max and last
	heap.arr[1], heap.arr[heap.count] = heap.arr[heap.count], heap.arr[1]
	heap.count--

	//heapify from up to down
	heapifyUpToDown(heap.arr, heap.count)
}

//heapify
func heapifyUpToDown(arr []int, count int) {

	for i := 1; i <= count/2; {

		maxIndex := i
		if arr[i] < arr[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= count && arr[maxIndex] < arr[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		i = maxIndex
	}

}

/////////////////////////////////////////////////////////////////////////

//build a heap
func buidHeap(a []int) {
	n := len(a) - 1
	//heapify from the last parent node
	for i := n / 2; i >= 1; i-- {
		heapifyUpToDown1(a, i, n)
	}

}

//sort by ascend, a index begin from 1, has n elements
func sort(a []int) {
	buidHeap(a)
	//fmt.Println(a)
	k := len(a) - 1
	for k >= 1 {
		a[1], a[k] = a[k], a[1]
		heapifyUpToDown1(a, 1, k-1)
		k--
	}
}

//heapify from up to down , node index = top
func heapifyUpToDown1(a []int, top int, count int) {

	for i := top; i <= count/2; {

		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		a[i], a[maxIndex] = a[maxIndex], a[i]
		i = maxIndex
	}

}
