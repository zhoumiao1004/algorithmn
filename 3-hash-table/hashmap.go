package main

import "container/list"

// 用拉链法解决哈希冲突的简化实现
type KVNode struct {
	key   int
	value int
}

// 底层 table 数组中的每个元素是一个链表
type ExampleChainingHashMap struct {
	table []*list.List
}

// 注意这里必须存储同时存储 key 和 value
// 因为要通过 key 找到对应的 value
func NewExampleChainingHashMap(capacity int) ExampleChainingHashMap {
	return ExampleChainingHashMap{
		table: make([]*list.List, capacity),
	}
}

func (h *ExampleChainingHashMap) hash(key int) int {
	return key % len(h.table)
}

// 查
func (h *ExampleChainingHashMap) Get(key int) (int, bool) {
	index := h.hash(key)
	if h.table[index] == nil {
		// 链表为空，说明 key 不存在
		return -1, false
	}

	for e := h.table[index].Front(); e != nil; e = e.Next() {
		node := e.Value.(KVNode)
		if node.key == key {
			return node.value, true
		}
	}

	// 链表中没有目标 key
	return -1, false
}

// 增/改
func (h *ExampleChainingHashMap) Put(key int, value int) {
	index := h.hash(key)
	if h.table[index] == nil {
		// 链表为空，新建一个链表，插入 key-value
		h.table[index] = list.New()
		h.table[index].PushBack(KVNode{key, value})
		return
	}

	for e := h.table[index].Front(); e != nil; e = e.Next() {
		node := e.Value.(KVNode)
		if node.key == key {
			// key 已经存在，更新 value
			node.value = value
			return
		}
	}

	// 链表中没有目标 key，添加新节点
	// 这里使用 addFirst 添加到链表头部或者 addLast 添加到链表尾部都可以
	// 因为 Java LinkedList 的底层实现是双链表，头尾操作都是 O(1) 的
	h.table[index].PushBack(KVNode{key, value})
}

// 删
func (h *ExampleChainingHashMap) Remove(key int) {
	index := h.hash(key)
	if h.table[index] == nil {
		return
	}

	for e := h.table[index].Front(); e != nil; e = e.Next() {
		node := e.Value.(KVNode)
		if node.key == key {
			// 如果 key 存在，则删除
			h.table[index].Remove(e)
			return
		}
	}
}
