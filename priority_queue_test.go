package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriorityQueueLen_Empty(t *testing.T) {
	pq := &PriorityQueue{
		items: make([]*Item, 0),
	}

	assert.Equal(t, 0, pq.Len(), "Length of empty queue should be 0")
}

func TestPriorityQueueLen_WithItems(t *testing.T) {
	pq := &PriorityQueue{
		items: []*Item{
			{Index: 1, Priority: 2, Value: "Val1"},
			{Index: 2, Priority: 3, Value: "Val2"},
		},
	}
	assert.Equal(t, 2, pq.Len(), "Length of queue with two items should be 2")
}

func TestPriorityQueueLess_WithItems(t *testing.T) {
	pq := &PriorityQueue{
		items: []*Item{
			{Index: 0, Priority: 5, Value: "Low"},
			{Index: 1, Priority: 15, Value: "High"},
			{Index: 2, Priority: 10, Value: "Medium"},
		},
	}

	assert.True(t, pq.Less(1, 0), "Should return true when first item has higher priority")
	assert.False(t, pq.Less(0, 1), "Should return false when first item has lower priority")
	assert.False(t, pq.Less(2, 2), "Should return false when priorities are equal")
}

func TestPriorityQueue_Swap(t *testing.T) {
	pq := &PriorityQueue{
		items: []*Item{
			{Index: 0, Priority: 1, Value: "First"},
			{Index: 1, Priority: 2, Value: "Second"},
			{Index: 2, Priority: 3, Value: "Third"},
		},
	}

	assert.Equal(t, "First", pq.items[0].Value, "The first item should initially be 'First'")
	assert.Equal(t, "Second", pq.items[1].Value, "The second item should initially be 'Second'")

	pq.Swap(0, 1)

	assert.Equal(t, "Second", pq.items[0].Value, "The first item should now be 'Second'")
	assert.Equal(t, "First", pq.items[1].Value, "The second item should now be 'First'")

	assert.Equal(t, 0, pq.items[0].Index, "The index of the first item should be updated to 0")
	assert.Equal(t, 1, pq.items[1].Index, "The index of the second item should be updated to 1")
}
