package main

import (
	"fmt"
)

func main() {
	var char string
	var length int
	var counter = make(map[byte]int)

	fmt.Printf("Masukan Karakter : ")
	fmt.Scanf("%s", &char)

	length = len([]rune(char))

	fmt.Printf("Karakter : %s \n", char)
	fmt.Printf("Jumlah Karakter : %d \n", length)

	for i := 0; i < len(char); i++ {
		counter[char[i]]++
	}

	fmt.Println("Karakter Duplicate : ")

	for key, value := range counter {
		if value > 1 {
			fmt.Println(string(key), ":", uint32(value))
		}
	}
}
