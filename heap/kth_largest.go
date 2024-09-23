package heap

import "container/heap"

type HeapNode struct {
	value     int
	heapIndex int
}

type KthLargest struct {
	minHeap *MinHeap
	k       int
}

type MinHeap []*HeapNode

func (m *MinHeap) Push(x any) {
	ele := x.(*HeapNode)
	ele.heapIndex = len(*m)
	*m = append(*m, ele)
}

func (m *MinHeap) Pop() any {
	list := *m
	ele := list[len(list)-1]
	*m = list[:len(list)-1]
	return ele
}

func (m MinHeap) Less(i, j int) bool {
	return m[i].value < m[j].value
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
	m[i].heapIndex = i
	m[j].heapIndex = j
}

func (m MinHeap) Len() int {
	return len(m)
}

func Constructor(k int, nums []int) KthLargest {
	var minHeap MinHeap
	inst := &KthLargest{
		minHeap: &minHeap,
		k:       k,
	}
	for _, num := range nums {
		inst.Add(num)
	}

	return *inst
}

func (this *KthLargest) Add(val int) int {
	if len(*this.minHeap) < this.k {
		heap.Push(this.minHeap, &HeapNode{value: val})
	} else {
		minHeap := *this.minHeap
		ele := minHeap[0]
		if val > ele.value {
			heap.Pop(this.minHeap)
			heap.Push(this.minHeap, &HeapNode{value: val})
		}
	}
	mH := *this.minHeap
	return mH[0].value
}
