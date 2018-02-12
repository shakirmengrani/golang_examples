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

func boring(msg string) <- chan Message{
	c := make(chan Message)
	go func(){
		for i := 0; ; i++{
			waitForIt := make(chan bool)
			c <- Message{str: fmt.Sprintf("%s %d",msg, i), wait: waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<- waitForIt
		}
	}()
	return c
}


func fanIn(input1, input2 <- chan Message) <- chan Message{
	c := make(chan Message)
	go func() { for { c <- <- input1 } }()
	go func() { for { c <- <- input2 } }()
	return c
}

func main(){
	c := fanIn(boring("shakir"), boring("bilal"))
	for i := 0; i < 5; i++{
		msg1 := <- c; fmt.Println(msg1.str)
		msg2 := <- c; fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're boring, I'm leaving.")
}