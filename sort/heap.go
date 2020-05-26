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
	heap.arr = make([]int, capacity)
	return heap
}

//top-max heap -> heapify from down to up
func (heap *Heap) insert(data int) {
	//defensive
	if heap.count == heap.cap {
		return
	}

	heap.arr[heap.count] = data
	heap.count++

	//compare with parent node
	i := heap.count - 1
	parent := i / 2
	for parent > 0 && heap.arr[parent] < heap.arr[i] {
		heap.arr[parent], heap.arr[i] = heap.arr[i], heap.arr[parent]
		i = parent
		parent = i / 2
	}
}

func (heap *Heap) getMax() int {
	if heap.count == 0 {
		return -1
	}
	return heap.arr[0]
}

//heapfify from up to down
func (heap *Heap) removeMax() {
	//defensive
	if heap.count == 0 {
		return
	}

	i := heap.count - 1

	//swap max and last
	heap.arr[0], heap.arr[i] = heap.arr[i], heap.arr[0]
	heap.count--

	//heapify from up to down
	heapifyUpToDown(heap.arr, heap.count)
}

//heapify
func heapifyUpToDown(arr []int, count int) {

	for i := 0; i <= (count-1)/2; {
		maxIndex := i
		if arr[i] < arr[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= (count-1) && arr[maxIndex] < arr[i*2+1] {
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

//堆排序
//s[0]不用，实际元素从角标1开始
//父节点元素大于子节点元素
//左子节点角标为2*k
//右子节点角标为2*k+1
//父节点角标为k/2
func HeapSort(s []int) {
	N := len(s) - 1 //s[0]不用，实际元素数量和最后一个元素的角标都为N
	//构造堆
	//如果给两个已构造好的堆添加一个共同父节点，
	//将新添加的节点作一次下沉将构造一个新堆，
	//由于叶子节点都可看作一个构造好的堆，所以
	//可以从最后一个非叶子节点开始下沉，直至
	//根节点，最后一个非叶子节点是最后一个叶子
	//节点的父节点，角标为N/2
	for k := N / 2; k >= 1; k-- {
		sink(s, k, N)
	}

	//下沉排序
	for N > 1 {
		swap(s, 1, N) //将大的放在数组后面，升序排序
		N--
		sink(s, 1, N)
	}
}

//下沉（由上至下的堆有序化）
func sink(s []int, k, N int) {
	for {
		i := 2 * k
		if i > N { //保证该节点是非叶子节点
			break
		}
		if i < N && s[i+1] > s[i] { //选择较大的子节点
			i++
		}
		if s[k] >= s[i] { //没下沉到底就构造好堆了
			break
		}
		swap(s, k, i)
		k = i
	}
}

func swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}
