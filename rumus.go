package main

import "fmt"

func main() {
	var panjang, lebar float64

	fmt.Printf("Masukan Satuan Panjang : ")
	fmt.Scanf("%f", &panjang)
	fmt.Printf("Masukan Satuan Lebar : ")
	fmt.Scanf("%f", &lebar)

	fmt.Printf("Panjang : %.2f \n", panjang)
	fmt.Printf("Lebar : %.2f \n", lebar)

	fmt.Printf("Luas Persegi Panjang : %.2f M² \n", lebar*panjang)
	fmt.Printf("Keliling Persegi Panjang : %.2f M \n", (lebar+panjang)*2)
}
