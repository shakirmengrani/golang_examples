package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo (c chan int, number int){
	defer wg.Done()
	c <- number * 5
}


func main(){
	var chanLenght int = 10
	var fooVal = make(chan int, chanLenght)
	for i := 0; i < chanLenght; i++{
		wg.Add(1)
		go foo(fooVal, i)
	}
	wg.Wait()
	close(fooVal)
	for item := range fooVal{
		fmt.Println(item)
	}

}