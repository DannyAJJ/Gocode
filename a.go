package main

import (
	"fmt"
)

func main() {
	//area del rectangulo
	var base, altura float64
	base = 10
	altura = 15
	area := base * altura
	fmt.Printf("El area del rectangulo es %v\n", area)
	//area del Trapecio
	var base2, altura2, Base2 float64
	base2 = 10
	Base2 = 15
	altura2 = 20
	area2 := ((base2 + Base2) / 2) * altura2
	fmt.Printf("El area del Trapecio es %v\n", area2)
	//area del circulo
	var radio float64
	radio = 5
	const pi = 3.14159265359
	area3 := pi * (radio * radio)
	fmt.Printf("El area del circulo es %v\n", area3)
}
