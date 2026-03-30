package main

import (
	"container/heap"
)

type Segment struct {
	id    int
	value int64
	index int
	next  *Segment
	prev  *Segment
}

type Item struct {
	seg   *Segment
	cost  int64
	index int
	valid bool
}

type MinHeap []*Item

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].cost != h[j].cost {
		return h[i].cost < h[j].cost
	}
	return h[i].index < h[j].index
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func minimumPairRemoval(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	nodes := make([]*Segment, n)
	violations := 0
	for i, v := range nums {
		nodes[i] = &Segment{id: i, value: int64(v), index: i}
		if i > 0 {
			nodes[i].prev = nodes[i-1]
			nodes[i-1].next = nodes[i]
			if nodes[i-1].value > nodes[i].value {
				violations++
			}
		}
	}

	h := &MinHeap{}
	heap.Init(h)

	activeItems := make(map[*Segment]*Item)

	curr := nodes[0]
	for curr != nil && curr.next != nil {
		cost := curr.value + curr.next.value
		item := &Item{
			seg:   curr,
			cost:  cost,
			index: curr.index,
			valid: true,
		}
		activeItems[curr] = item
		heap.Push(h, item)
		curr = curr.next
	}

	ops := 0

	for violations > 0 {
		var bestItem *Item
		for h.Len() > 0 {
			bestItem = heap.Pop(h).(*Item)
			if bestItem.valid {
				break
			}
		}

		if bestItem == nil || !bestItem.valid {
			break
		}

		ops++

		segI := bestItem.seg
		segJ := segI.next

		if segI.prev != nil {
			if oldItem, exists := activeItems[segI.prev]; exists {
				oldItem.valid = false
			}
			if segI.prev.value > segI.value {
				violations--
			}
		}

		if segJ.next != nil {
			if oldItem, exists := activeItems[segJ]; exists {
				oldItem.valid = false
			}
			if segJ.value > segJ.next.value {
				violations--
			}
		}

		if segI.value > segJ.value {
			violations--
		}

		segI.value += segJ.value
		segI.next = segJ.next
		if segI.next != nil {
			segI.next.prev = segI
		}

		if segI.prev != nil {
			if segI.prev.value > segI.value {
				violations++
			}
			cost := segI.prev.value + segI.value
			newItem := &Item{
				seg:   segI.prev,
				cost:  cost,
				index: segI.prev.index,
				valid: true,
			}
			activeItems[segI.prev] = newItem
			heap.Push(h, newItem)
		}

		if segI.next != nil {
			if segI.value > segI.next.value {
				violations++
			}
			cost := segI.value + segI.next.value
			newItem := &Item{
				seg:   segI,
				cost:  cost,
				index: segI.index,
				valid: true,
			}
			activeItems[segI] = newItem
			heap.Push(h, newItem)
		}
	}

	return ops
}