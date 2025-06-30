package main

import (
	"container/heap"
	"fmt"
	"math/rand"
)

// 小顶堆实现
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// 模拟 1 万个随机整数
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = rand.Intn(1000000)
	}
	
	h := &MinHeap{}
	heap.Init(h)
	
	for _, num := range nums {
		if h.Len() < 10 {
			heap.Push(h, num)
		} else if num > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, num)
		}
	}
	
	// 输出结果
	fmt.Println("最大的10个数：")
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
