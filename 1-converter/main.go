package main

import "fmt"

func main() {
	usdToEuro := 0.85
	usdToRub := 83.5
	EuroToRub := usdToRub / usdToEuro
	fmt.Print(EuroToRub)
}
