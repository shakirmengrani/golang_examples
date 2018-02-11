package main

import(
	"fmt"
	"time"
	"sync"
)

var wait_group sync.WaitGroup


func cleanup(){
	defer wait_group.Done()
	if r := recover(); r != nil{
		fmt.Println("Recovered panic: ", r)
	}
}

func say(s string){ 
	defer cleanup()
	for i := 0; i < 3; i++{
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("Oh no panic here !")
		}
	}
	
	
}

func main(){
	wait_group.Add(1)
	go say("Hey")
	wait_group.Add(1)
	go say("there")
	wait_group.Wait()
}