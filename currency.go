package main

import (
	"fmt"
	"strings"

	humanize "github.com/dustin/go-humanize"
)

var idr, usd float64

func main() {
	fmt.Printf("Input your numeric : ")
	fmt.Scanf("%f", &idr)
	usd = idr / 14251.55
	fmt.Printf("IDR : %s,00 \n", FormatRupiah(idr))
	fmt.Printf("USD : $ %.2f \n", usd)
}

func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return "Rp " + stringValue
}
