package main 

import "fmt"

func main(){
	var dict = make(map[string]string)
	dict["shakir"] = "hello Master"

	fmt.Println(dict["shakir"])
}