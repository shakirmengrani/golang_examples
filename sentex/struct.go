package main

import (
	"fmt"
)

type car struct {
	year uint16
	modal string
	top_speed_kmh float64
}

func(c car) render(name string) string {
	return fmt.Sprintf("%s your car modal is %s and year is %d", name, c.modal, c.year)
}

func(c *car) setModal(name string){
	c.modal = name
}

func setModal(c car,name string) car{
	c.modal = name
	return c
}


func NewCar(year uint16, modal string, top_speed_kmh float64) car{
	return car{year:year, modal:modal, top_speed_kmh:top_speed_kmh}
}


func main(){
	a_car := car{year: 2018,modal: "Mehraan",top_speed_kmh: 225.0}
	fmt.Println(a_car.render("shakir"))
	a_car.setModal("city")
	fmt.Println(a_car.render("shakir"))
	a_car = setModal(a_car, "Corolla"	)
	fmt.Println(a_car.render("shakir"))
	a_car = NewCar(2014, "city", 225.0)
	fmt.Println(a_car.render("shakir"))
	
}