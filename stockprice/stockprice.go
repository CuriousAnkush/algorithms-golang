package stockprice

import (
	"container/heap"
	"math"
)

type StockPrice struct {
	auditRecord     map[int]*Price
	maxTimeStamp    int
	minPriceTracker MinPriceTracker
	maxPriceTracker MaxPriceTracker
}

func Constructor() StockPrice {
	return StockPrice{
		auditRecord:  make(map[int]*Price),
		maxTimeStamp: math.MinInt,
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if oldPrice, found := this.auditRecord[timestamp]; !found {
		this.auditRecord[timestamp] = &Price{
			minHeapIndex: 0,
			maxHeapIndex: 0,
			value:        price,
		}
		heap.Push(&this.minPriceTracker, this.auditRecord[timestamp])
		heap.Push(&this.maxPriceTracker, this.auditRecord[timestamp])
	} else {
		oldPrice.value = price
		heap.Fix(&this.minPriceTracker, oldPrice.minHeapIndex)
		heap.Fix(&this.maxPriceTracker, oldPrice.maxHeapIndex)
	}
	if timestamp > this.maxTimeStamp {
		this.maxTimeStamp = timestamp
	}
}

func (this *StockPrice) Current() int {
	return this.auditRecord[this.maxTimeStamp].value
}

func (this *StockPrice) Maximum() int {
	return this.maxPriceTracker[0].value
}

func (this *StockPrice) Minimum() int {
	return this.minPriceTracker[0].value
}

type Price struct {
	minHeapIndex int
	maxHeapIndex int
	value        int
	heap.Interface
}

type MinPriceTracker []*Price

func (m *MinPriceTracker) Push(x any) {
	price := x.(*Price)
	price.minHeapIndex = len(*m)
	*m = append(*m, price)
}

func (m *MinPriceTracker) Pop() any {
	tracker := *m
	length := len(tracker)
	elem := tracker[length-1]
	*m = tracker[0 : length-1]
	return elem
}

func (m MinPriceTracker) Len() int {
	return len(m)
}

func (m MinPriceTracker) Less(i, j int) bool {
	tracker := m
	return tracker[i].value < tracker[j].value
}

func (m MinPriceTracker) Swap(i, j int) {
	tracker := m
	tracker[i], tracker[j] = tracker[j], tracker[i]
	tracker[i].minHeapIndex = i
	tracker[j].minHeapIndex = j
}

type MaxPriceTracker []*Price

func (m *MaxPriceTracker) Push(x any) {
	price := x.(*Price)
	price.maxHeapIndex = len(*m)
	*m = append(*m, price)
}

func (m *MaxPriceTracker) Pop() any {
	tracker := *m
	length := len(tracker)
	elem := tracker[length-1]
	tracker[0].maxHeapIndex = length - 1
	*m = tracker[0 : length-1]
	return elem
}

func (m MaxPriceTracker) Len() int {
	return len(m)
}

func (m MaxPriceTracker) Less(i, j int) bool {
	tracker := m
	return tracker[i].value > tracker[j].value
}

func (m MaxPriceTracker) Swap(i, j int) {
	tracker := m
	tracker[i], tracker[j] = tracker[j], tracker[i]
	tracker[i].maxHeapIndex = i
	tracker[j].maxHeapIndex = j
}
