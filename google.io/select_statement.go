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

func boring(msg string, quit chan bool) <- chan Message{
	c := make(chan Message)
	go func(){
		for i := 0; ; i++{
			waitForIt := make(chan bool)
			c <- Message{str: fmt.Sprintf("%s %d",msg, i), wait: waitForIt}:
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<- waitForIt
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
	for{
		select{
		case s := <- b: 
			fmt.Println(s.str)
		case <- time.After(5 * time.Second):
			fmt.Println("You're too slow.")
			return
		}
	}
	c := fanIn(boring("shakir", quit), boring("bilal", quit))
	for i := 0; i < 5; i++{
		msg1 := <- c; fmt.Println(msg1.str)
		msg2 := <- c; fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're boring, I'm leaving.")
}