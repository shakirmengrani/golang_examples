package main

import "fmt"

func foo (c chan int, number int){
	c <- number * 5
}


func main(){
	var fooVal = make(chan int)
	go foo(fooVal, 3)
	go foo(fooVal, 5)
	val1, val2 := <- fooVal, <- fooVal
	fmt.Println(val1, val2)
}