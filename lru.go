package dutils

import (
	"math"
)

type lruHeapItem struct {
	key string
	v   uint
}

type lruHeap struct {
	leng  uint
	value []lruHeapItem
	lru   *LRU
}

type LRU struct {
	heap  *lruHeap
	p     map[string]*uint
	count map[string]uint
	max   uint
}

func NewLRU(NMax uint) *LRU {
	lru := &LRU{
		heap: &lruHeap{
			leng:  0,
			value: make([]lruHeapItem, NMax+1),
		},
		count: make(map[string]uint, 0),
		p:     make(map[string]*uint, 0),
		max:   NMax,
	}
	lru.heap.lru = lru
	return lru
}

// Increase usage for given key
func (l *LRU) Inc(key string) {
	if l.p[key] == nil {
		if l.heap.leng >= l.max {
			l.drop()
			l.Inc(key)
			return
		}
		l.count[key] = 1
		l.heap.insert(lruHeapItem{
			key: key,
			v:   1,
		})
	} else {
		l.count[key] = l.count[key] + 1
		p := l.p[key]
		l.heap.value[*p].v += 1
		l.heap.downHeap(*p)
	}
}

// Get least recently used key
func (l *LRU) Least() string {
	return l.heap.value[1].key
}

func (l *LRU) drop() {
	key := l.heap.root()
	delete(l.count, key.key)
	delete(l.count, key.key)
	l.heap.removeRoot()
}

func (l *LRU) updatePostition(key string, i uint) {
	l.p[key] = &i
}

func (h *lruHeap) insert(x lruHeapItem) *lruHeap {
	h.leng = h.leng + 1
	h.value[h.leng] = x
	p := h.leng
	h.lru.updatePostition(x.key, p)
	h.upHeap(h.leng)
	return h
}

func (h *lruHeap) upHeap(i uint) *lruHeap {
	k := uint(math.Floor(float64(i) / 2))
	if i == 1 || h.value[k].v < h.value[i].v {
		return h
	}
	h.swap(i, k)
	h.upHeap(k)
	return h
}

func (h *lruHeap) swap(i, j uint) *lruHeap {
	t := h.value[i]
	h.value[i] = h.value[j]
	h.value[j] = t

	h.lru.updatePostition(h.value[i].key, i)
	h.lru.updatePostition(h.value[j].key, j)
	return h
}

func (h *lruHeap) downHeap(i uint) *lruHeap {
	m := i * 2
	if m > h.leng {
		return h
	}
	if m+1 < h.leng && h.value[m].v > h.value[m+1].v {
		m++
	}
	if h.value[m].v < h.value[i].v {
		// switch two node
		h.swap(m, i)
		h.downHeap(m)
		return h

	}
	return h
}

func (h *lruHeap) removeRoot() *lruHeap {
	h.swap(h.leng, 1)
	h.leng--
	if h.leng > 1 {
		h.downHeap(1)
	}
	return h
}

func (h *lruHeap) root() lruHeapItem {
	return h.value[1]
}
