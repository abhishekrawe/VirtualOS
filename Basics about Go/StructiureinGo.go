// go run StructiureinGo.go




package main 

import "fmt"

type Car struct {

	Model string 

	Mileage int 

	Brand string
}

func main() {
	var car = Car{"X4" ,  15 , "BMW" }
    car.Model = "X5"

	fmt.Println(car.Model)
}