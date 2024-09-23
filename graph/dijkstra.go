package graph

import (
	"container/heap"
	"math"
)

type Edge struct {
	parentNode int
	distance   int
	sourceNOde int
	heapIndex  int
}

type priorityQueue []*Edge

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i].distance < p[j].distance
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].heapIndex = i
	p[j].heapIndex = j
}

func (p *priorityQueue) Push(x any) {
	node := x.(*Edge)
	node.heapIndex = len(*p)
	*p = append(*p, node)
}

func (p *priorityQueue) Pop() any {
	list := *p
	leng := len(list)

	el := list[leng-1]
	*p = list[0 : leng-1]
	return el
}

func networkDelayTime(times [][]int, n int, k int) int {
	adjList := make(map[int][]*Edge)
	genAdjList(times, adjList)
	maxDelay := math.MinInt

	edge := &Edge{
		parentNode: k,
		distance:   0,
		sourceNOde: k,
		heapIndex:  0,
	}
	distanceMap := make(map[int]*Edge)

	distanceMap[k] = edge
	pq := priorityQueue{edge}
	heap.Init(&pq)

	for len(pq) != 0 {
		sourceVertex := heap.Pop(&pq).(*Edge)
		sourceEdge := *sourceVertex
		sourceDist := sourceEdge.distance
		for _, edge := range adjList[sourceEdge.sourceNOde] {
			name, distance := edge.sourceNOde, edge.distance

			if distanceMap[name] != nil && distanceMap[name].distance <= sourceDist+distance {
				continue
			}

			if distanceMap[name] != nil {
				distanceMap[name].distance = sourceDist + distance
				distanceMap[name].parentNode = sourceEdge.parentNode
				heap.Fix(&pq, distanceMap[name].heapIndex)
			} else {
				distanceMap[name] = &Edge{
					parentNode: sourceEdge.parentNode,
					distance:   distance + sourceDist,
					sourceNOde: name,
				}
				heap.Push(&pq, distanceMap[name])
			}
		}
	}

	for i := 1; i <= n; i++ {
		if distanceMap[i] == nil {
			return -1
		}

		maxDelay = max(maxDelay, distanceMap[i].distance)
	}

	return maxDelay
}

func genAdjList(times [][]int, adjList map[int][]*Edge) {
	for i := range times {
		sourceVertex := times[i][0]
		desVertex := times[i][1]
		distance := times[i][2]
		if list, found := adjList[sourceVertex]; !found {
			var list []*Edge
			list = append(list, &Edge{
				parentNode: sourceVertex,
				distance:   distance,
				sourceNOde: desVertex,
			})
			adjList[sourceVertex] = list
		} else {
			list = append(list, &Edge{
				parentNode: sourceVertex,
				distance:   distance,
				sourceNOde: desVertex,
			})
			adjList[sourceVertex] = list
		}
	}
}
