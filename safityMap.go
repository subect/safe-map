package main

import "sync"

type RWMap struct {
	sync.RWMutex
	m map[int]int
}

// NewRWMap 新建一个map
func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]int, n),
	}
}

// Get 从map读取值
func (m *RWMap) Get(n int) (int, bool) {
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[n]
	return v, existed
}

// Set 设置一个值
func (m *RWMap) Set(k int, v int) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

// Delete 删除一个键
func (m *RWMap) Delete(n int) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, n)
}

func (m *RWMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}
