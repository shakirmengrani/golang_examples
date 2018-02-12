package main

import (
	"fmt"
	"time"
	"math/rand"
)


func boring_less_bore(msg string){
	for i := 0; ; i++{
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}


func boring_norm(msg string){
	for i := 0; ; i++{
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}


func main(){
	go boring_less_bore("boring!")
	fmt.Println("I'm listing..")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring, I'm leaving.")
}