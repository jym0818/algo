package main

import "fmt"

//拉链法解决哈希冲突

type Pair struct {
	key   int
	value string
	next  *Pair
}

type HashTable struct {
	size        int
	capacity    int
	buckets     []*Pair
	loadThres   float64
	extendRatio int
}

func NewHashTable(capacity int) *HashTable {
	if capacity <= 0 {
		capacity = 1
	}
	s := make([]*Pair, capacity)
	return &HashTable{
		size:        0,
		capacity:    capacity,
		buckets:     s,
		loadThres:   0.75,
		extendRatio: 2,
	}
}
func (h *HashTable) hashFunc(key int) int {
	return key % h.capacity
}
func (h *HashTable) loadFactor() float64 {
	return float64(h.size) / float64(h.capacity)
}

func (h *HashTable) Get(key int) string {
	index := h.hashFunc(key)
	cur := h.buckets[index]
	for cur != nil {
		if cur.key == key {
			return cur.value
		}
		cur = cur.next
	}
	return ""
}

// 添加或者更新一个值
func (h *HashTable) Put(key int, value string) {
	//触发扩容
	if h.loadFactor() >= h.loadThres {
		h.resize()
	}

	index := h.hashFunc(key)

	//遍历这个桶的链表如果有该值则更新
	bucket := h.buckets[index]
	for bucket != nil {
		if bucket.key == key {
			bucket.value = value
			return
		}
		bucket = bucket.next
	}
	pair := &Pair{key: key, value: value}
	pair.next = h.buckets[index]
	h.buckets[index] = pair
	h.size++

}
func (h *HashTable) resize() {
	//暂存原哈希表的桶数组
	oldBuckets := h.buckets

	newCapacity := h.capacity * h.extendRatio
	newBuckets := make([]*Pair, newCapacity)
	h.capacity = newCapacity
	h.buckets = newBuckets
	h.size = 0
	//  将键值对从原哈希表搬运至新哈希表
	for _, bucket := range oldBuckets {
		for bucket != nil {
			h.Put(bucket.key, bucket.value)
			bucket = bucket.next
		}
	}
}

func (h *HashTable) remove(key int) {
	index := h.hashFunc(key)
	cur := h.buckets[index]
	pre := (*Pair)(nil)
	for cur != nil {
		if cur.key == key {
			if pre == nil {
				h.buckets[index] = cur.next
			} else {
				pre.next = cur.next
			}
			h.size--
			return
		}
		pre = cur
		cur = cur.next
	}
}

func (h *HashTable) Print() {
	for i, bucket := range h.buckets {
		fmt.Printf("Slot %d: ", i)
		for bucket != nil {
			fmt.Printf("(%d: %s) -> ", bucket.key, bucket.value)
			bucket = bucket.next
		}
		fmt.Println("nil")
	}
}
func main() {
	h := NewHashTable(2)

	h.Put(1, "1")

	h.Put(2, "2")

	h.Put(5, "2")

	h.Print()
}
