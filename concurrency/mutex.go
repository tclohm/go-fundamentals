package main

// mutex -- mutual exclusion
// provides a way to lock and unlock

import (
	"sync"
	"fmt"
)

type SyncedData struct {
	inner map[string]int
	mutex sync.Mutex
}

func (d *SyncedData) Insert(k string, v int) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.inner[k] = v
}

func (d *SyncedData) Get(k string) int {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.inner[k]
}

func main() {
	data := SyncedData{inner: make(map[string]int)}
	data.Insert("sample", 5)
	data.Insert("test", 2)
	fmt.Println(data.Get("sample"))
	fmt.Println(data.Get("test"))
}