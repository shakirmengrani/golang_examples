package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup


type syncCounter struct{
	c int
}


func(sc* syncCounter) set(n int){
	defer wg.Done()
	sc.c = n
	fmt.Println("Work done ",sc.c)
}

func main(){
	sc := new(syncCounter)
	for i := 0; i < 10; i++{
		wg.Add(1)
		go sc.set(i)
	}
	wg.Wait()
	fmt.Println(sc.c)
}