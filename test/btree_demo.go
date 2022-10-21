package main

import (
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

type children []*node

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
