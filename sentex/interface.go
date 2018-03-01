package main

import "fmt"

type geo interface{
	Lat() float64
	Lng() float64
}

type location struct{
	name string
}

func (loc location)Lat() float64{
	return 0.0
}

func (loc location)Lng() float64{
	return 0.0
}


func get_lat_lng(g geo) (float64, float64) {
	return g.Lat(), g.Lng()
}

func main(){
	l := location{name:"Karachi"}
	fmt.Println(get_lat_lng(l))

}