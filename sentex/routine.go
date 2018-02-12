package main

import(
	"fmt"
	"time"
)

func say(s string){
	for i := 0; i < 3; i++{
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}

func main(){
	go say("Hey")
	// say("There")
	go say("there")
	time.Sleep(time.Second)
}