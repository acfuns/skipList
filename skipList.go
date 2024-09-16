package main

import "math/rand"

const MAX_LEVEL = 16

type Node struct {
	val  int
	next []*Node
}

type SkipList struct {
	head  *Node
	level int
}

// 插入时随机生成一个level
func getRandomLevel() int {
	level := 0
	for rand.Float64() < 0.5 && level < MAX_LEVEL-1 {
		level++
	}
	return level
}

func NewSkipList() *SkipList {
	head := &Node{val: -1, next: make([]*Node, MAX_LEVEL)}
	return &SkipList{head: head, level: 0}
}

func (sl *SkipList) Search(val int) bool {
	current := sl.head
	for i := sl.level; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].val < val {
			current = current.next[i]
		}
	}
	current = current.next[0]
	return current != nil && current.val == val
}

func (sl *SkipList) Insert(val int) {
	update := make([]*Node, MAX_LEVEL)
	current := sl.head
	for i := sl.level; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].val < val {
			current = current.next[i]
		}
		update[i] = current
	}

	updateLevel := getRandomLevel()
	if updateLevel > sl.level {
		for i := sl.level + 1; i <= updateLevel; i++ {
			update[i] = sl.head
		}
		sl.level = updateLevel
	}

	newNode := &Node{val: val, next: make([]*Node, updateLevel+1)}

	for i := 0; i <= updateLevel; i++ {
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}
}

func (sl *SkipList) Remove(val int) {
	if !sl.Search(val) {
		return
	}
	update := make([]*Node, MAX_LEVEL)
	current := sl.head
	for i := sl.level; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].val < val {
			current = current.next[i]
		}
		if current.next[i] != nil && current.next[i].val == val {
			update[i] = current
		}
	}

	for i := 0; i <= sl.level; i++ {
		if update[i] != nil && update[i].next[i].val == val {
			update[i].next[i] = update[i].next[i].next[i]
		}
	}
}
