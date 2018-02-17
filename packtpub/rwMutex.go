package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

type mapCounter struct{
	c map[int]int
	sync.RWMutex
}

func runReaders(m *mapCounter, n int){
	for{
		m.RLock()
		fmt.Println(m.c[rand.Intn(n)])
		m.RUnlock()
		time.Sleep(1 * time.Second)
	}
}

func runWriters(m *mapCounter, n int){
	for i:=0; i < n; i++{
		m.Lock()
		m.c[i] = i * 10
		m.Unlock()
		time.Sleep(1 * time.Second)
	}	
}


func main(){
	mc := mapCounter{ c: make(map[int]int) }
	go runWriters(&mc,10)
	go runReaders(&mc,10)
	go runReaders(&mc,10)
	time.Sleep(15 * time.Second)
}