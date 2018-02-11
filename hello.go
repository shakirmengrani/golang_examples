package main
import "fmt"


func sayHello(name string) string{
	return "Hello, " + name
}

func main(){
	var x int = 15
	a := &x
	*a = 5
	*a = *a**a
	var n1,n2 float64 = 1.0, 2.0
	n3, n4 := 3.0, 4.0
	fmt.Println(a)
	fmt.Println(x)
	fmt.Println("Hello from Golang")
	fmt.Println(sayHello("Shakir"))
	fmt.Println( ((n1 + n2) * n3) / n4 )
}