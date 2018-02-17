package main

import (
	"fmt"
	"sync"
)


type safeCounter struct{
	i int
	sync.Mutex
}

func (sc *safeCounter) Increment() {
	sc.Lock()
	defer sc.Unlock()
	sc.i ++;
}

func (sc *safeCounter) Decrement() {
	sc.Lock()
	defer sc.Unlock()
	sc.i --;
}

func (sc *safeCounter) render(){
	fmt.Println(sc.i)
}


func NewSafeCounter(startsfrom int) *safeCounter{
	return &safeCounter{ i: startsfrom }
}



var (
	counter1 = NewSafeCounter(100)
	counter2 = NewSafeCounter(101)
	counter3 = NewSafeCounter(102)
)

func main(){
	for i:=0;i<100;i++{
		counter1.Increment()
		counter1.render()
		counter1.Decrement()
		counter1.render()

		counter2.Increment()
		counter2.render()
		counter2.Decrement()
		counter2.render()

		counter3.Increment()
		counter3.render()
		counter3.Decrement()
		counter3.render()
	}
}