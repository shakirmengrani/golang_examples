package main

import (
	"fmt"
	"time"
	"math/rand"
)

func boring(msg string) <- chan string{
	c := make(chan string)
	go func(){
		for i := 0; ; i++{
			c <- fmt.Sprintf("%s %d",msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main(){
	// var c chan string
	// c = make(chan int)
	c := boring("boring!")
	shakir := boring("shakir")
	for i := 0; i < 5; i++{
		fmt.Printf("You say: %s\n",<-c)
		fmt.Printf("You say: %s\n",<-shakir)
	}
	fmt.Println("You're boring, I'm leaving.")
}