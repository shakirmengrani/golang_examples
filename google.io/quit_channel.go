package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Message struct{
	str string
	wait chan bool
}

func boring(msg string, quit chan string) <- chan Message{
	c := make(chan Message)
	go func(){
		for i := 0; ; i++{
			waitForIt := make(chan bool)
			select{
			case c <- Message{str: fmt.Sprintf("%s %d",msg, i), wait: waitForIt}:
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
				<- waitForIt
			case <- quit:
				quit <- "See you again"
				return
			}
		}
	}()
	return c
}


func fanIn(input1, input2 <- chan Message) <- chan Message{
	c := make(chan Message)
	go func() { 
		for {
			select{
			case s := <- input1: c <- s
			case s := <- input2: c <- s
			default: fmt.Println("no more value for communicate") 
			} 
		} 
	}()
	return c
}

func main(){
	quit := make(chan string)
	b := boring("Joe", quit)
	for i := 0; ; i++{
		s := <-b; fmt.Println(s.str)
		s.wait <- true
		if i == 5 {
			quit <- "Good Bye"
			fmt.Println(<- quit)
			return 
		}
		
	}
	
	fmt.Println("You're boring, I'm leaving.")
}