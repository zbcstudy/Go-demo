package main

import (
	"fmt"
	"sort"
	"sync"
)

type Item interface {
	Less(than Item) bool
}

const (
	DefaultFreeListSize = 32
)

var (
	nilItems    = make(items, 16)
	nilChildren = make(children, 16)
)

type FreeList struct {
	mu       sync.Mutex
	freeList []*node
}

func NewFreeList(size int) *FreeList {
	return &FreeList{
		freeList: make([]*node, 0, size),
	}
}

func (f *FreeList) newNode() (n *node) {
	f.mu.Lock()
	index := len(f.freeList) - 1
	if index < 0 {
		f.mu.Unlock()
		return new(node)
	}
	n = f.freeList[index]
	f.freeList[index] = nil
	f.freeList = f.freeList[:index]
	f.mu.Unlock()
	return
}

// 新增节点
func (f *FreeList) freeNode(n *node) (out bool) {
	f.mu.Lock()
	if len(f.freeList) < cap(f.freeList) {
		f.freeList = append(f.freeList, n)
		out = true
	}
	f.mu.Unlock()
	return
}

type ItemIterator func(i Item) bool

func New(degree int) *BTree {
	return NewWithFreeList(degree, NewFreeList(DefaultFreeListSize))
}

func NewWithFreeList(degree int, f *FreeList) *BTree {
	if degree <= 1 {
		panic("bad degree")
	}
	return &BTree{
		degree: degree,
		cow:    &copyOnWriteContext{freeList: f},
	}
}

type items []Item

func (s *items) InsertAt(index int, item Item) {
	*s = append(*s, nil)
	if index < len(*s) {
		copy((*s)[index+1:], (*s)[index:])
	}
	(*s)[index] = item
}

func (s *items) removeAt(index int) Item {
	if index > len(*s) {
		fmt.Println("下标超过数组长度")
	}
	item := (*s)[index]
	copy((*s)[index:], (*s)[index+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return item
}

// remove and return last element
func (s *items) pop() (out Item) {
	// 最大的下标index
	index := len(*s) - 1
	out = (*s)[index]
	// 移除最后一个元素
	(*s)[index] = nil
	*s = (*s)[:index]
	return
}

func (s *items) truncate(index int) {
	var toClear items
	*s, toClear = (*s)[:index], (*s)[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, nilItems):]
	}
}

func (s items) find(item Item) (index int, found bool) {
	i := sort.Search(len(s), func(i int) bool {
		return item.Less(s[i])
	})
	if i > 0 && !s[i-1].Less(item) {
		return i - 1, true
	}
	return i, false
}

type children []*node

// 添加新的节点
func (s *children) insertAt(index int, n *node) {
	*s = append(*s, nil)
	if index < len(*s) {
		copy((*s)[index+1:], (*s)[index:])
	}
	(*s)[index] = n
}

func (s *children) removeAt(index int) *node {
	n := (*s)[index]
	copy((*s)[index:], (*s)[index+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return n
}

type node struct {
	items    items
	children children
	cow      *copyOnWriteContext
}

type copyOnWriteContext struct {
	freeList *FreeList
}

type BTree struct {
	degree int
	length int
	root   *node
	cow    *copyOnWriteContext
}
