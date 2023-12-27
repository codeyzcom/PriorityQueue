package main

import "sync"

type Item struct {
	Index    int
	Priority int
	Value    interface{}
}

type PriorityQueue struct {
	mutex sync.Mutex
	items []*Item
}

func (pq *PriorityQueue) Len() int {
	return len(pq.items)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.items[i].Priority > pq.items[j].Priority
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].Index = i
	pq.items[j].Index = j
}

func (pq *PriorityQueue) Push(i *Item) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	n := len(pq.items)
	i.Index = n
	pq.items = append(pq.items, i)
}

func (pq *PriorityQueue) Pop() (*Item, bool) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	old := pq.items
	n := len(old)
	if n < 1 {
		return nil, false
	}
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	pq.items = old[0 : n-1]
	return item, true
}
