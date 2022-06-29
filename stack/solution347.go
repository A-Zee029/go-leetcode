package stack

import (
	"container/heap"
)

type freHeap [][2]int

func (h freHeap) Len() int {
	return len(h)
}

func (h freHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h freHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *freHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *freHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	h := &freHeap{}
	heap.Init(h)
	freMap := make(map[int]int)

	for _, num := range nums {
		freMap[num]++
	}
	for key, val := range freMap {
		heap.Push(h, [2]int{key, val})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).([2]int)[0])
	}
	return res
}

//
//func main() {
//	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
//}
