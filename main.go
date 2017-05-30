package main

import "net"
import "sync"

func main() {
	var m sync.Mutex
	m.Lock()
	m.Unlock()
	net.Pipe()
}
