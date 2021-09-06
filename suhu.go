package main

import "fmt"

func main() {
	var celcius, fahrenheit, kelvin float64

	fmt.Printf("Masukan satuan celcius : ")
	fmt.Scanf("%f", &celcius)
	fmt.Printf("Suhu Celcius : %.2f \n", celcius)

	fahrenheit = celcius*9/5 + 32
	fmt.Printf("Suhu Fahrenheit : %.2f \n", fahrenheit)

	kelvin = celcius + 273.15
	fmt.Printf("Suhu Kelvin : %.2f \n", kelvin)
}
