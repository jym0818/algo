package main

import "fmt"

// Pair就是桶
type Pair struct {
	key   int
	value string
}

type HashTable struct {
	buckets []*Pair
}

func NewHashTable() *HashTable {
	//初始化100个桶
	sl := make([]*Pair, 100)
	return &HashTable{
		buckets: sl,
	}
}

// 哈希函数
func (h *HashTable) HashFunc(key int) int {
	return key % 100
}

/* 添加操作 */
func (h *HashTable) Put(key int, value string) {
	pair := &Pair{value: value, key: key}
	index := h.HashFunc(key)
	h.buckets[index] = pair
}

/*读取*/

func (h *HashTable) Get(key int) string {
	index := h.HashFunc(key)
	pair := h.buckets[index]
	if pair == nil {
		return "not found"
	}

	return pair.value

}

// 删除操作
func (h *HashTable) Delete(key int) {
	index := h.HashFunc(key)
	// 置为 nil ，代表删除
	h.buckets[index] = nil
}

/*打印hash*/
func (h *HashTable) Print() {
	for _, value := range h.buckets {
		if value != nil {
			fmt.Printf("%d->%s\n", value.key, value.value)
		}
	}
}

func main() {
	h := NewHashTable()
	h.Put(1000, "zhangsan")
	h.Put(1001, "lisi")
	h.Print()
}
