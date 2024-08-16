package main

import "fmt"

func main() {

	const TempK = 373.2
	var tempC float64 = TempK - 273

	fmt.Printf("O ponto de ebulição da água em Kelvin é: %g e o ponto de ebulição da água em Celsius é %g", TempK, tempC)

}
