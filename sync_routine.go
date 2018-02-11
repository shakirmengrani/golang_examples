package main

import(
	"fmt"
	"time"
	"sync"
)

var wait_group sync.WaitGroup

func say(s string){ 
	for i := 0; i < 3; i++{
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	wait_group.Done()
	
}

func main(){
	wait_group.Add(1)
	go say("Hey")
	wait_group.Add(1)
	go say("there")
	wait_group.Wait()
}