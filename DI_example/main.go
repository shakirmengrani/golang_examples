package main

import (
	"fmt"

	"./innerPkg"
)

var i int

func init() {
	fmt.Println("Init fired")
	i = 12
}

func main() {
	fmt.Println("Main fired")
	fmt.Println(fmt.Sprintf("The value of i is %d", i))
	di := InnerPkg.NewDI(&InnerPkg.Config{Name: "Hello", Port: 3000})
	fmt.Println(di)
}
