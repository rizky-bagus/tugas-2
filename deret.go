package main

import "fmt"

func main() {

	var deret int

	fmt.Printf("Masukan junlah deret : ")
	fmt.Scanf("%d", &deret)

	fmt.Printf("Jumlah Deret : %d \n", deret)

	fmt.Printf("Deret Genap : ")
	for i := 0; i < deret*2; i++ {
		if i%2 == 0 {
			fmt.Printf(" %d ", i)
		}
	}

	fmt.Printf("\nDeret Ganjil : ")
	for i := 0; i < deret*2; i++ {
		if i%2 == 1 {
			fmt.Printf(" %d ", i)
		}
	}
	fmt.Println("")
}
