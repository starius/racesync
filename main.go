package main

import "sync"

func main() {
	var m sync.Mutex
	m.Lock()
	m.Unlock()
}
