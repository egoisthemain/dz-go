package main

import "fmt"

func main() {
	usdToEuro := 0.85
	usdToRub := 83.5
	EuroToRub := usdToRub / usdToEuro
	fmt.Print(EuroToRub)
}

func getUserInfo() {
	var money float64
	fmt.Print("Введите сумму в исходной валюте: ")
	fmt.Scan(&money)
}

func convert(money float64, origValue, targetValue string) {}
