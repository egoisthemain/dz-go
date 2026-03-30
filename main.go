package main

import (
	"fmt"
)

type conversionMap map[string]map[string]float64

func main() {
	RubToUsd := 0.012
	RubToEur := 0.011

	convMap := conversionMap{
		"rub": map[string]float64{
			"usd": RubToUsd,
			"eur": RubToEur,
		},
		"usd": map[string]float64{
			"rub": 1 / RubToUsd,
			"eur": RubToEur / RubToUsd,
		},
		"eur": map[string]float64{
			"rub": 1 / RubToEur,
			"usd": RubToUsd / RubToEur,
		},
	}
	origCurrency, origMoney, targetCurrency := getUserInfo()
	//fmt.Printf("origCurrency = %s, origMoney = %.2f, targetCurrency = %s\n", origCurrency, origMoney, targetCurrency)
	convertValue := convert(origMoney, origCurrency, targetCurrency, &convMap)
	fmt.Printf("%.2f %s = %.2f %s", origMoney, origCurrency, convertValue, targetCurrency)
}

func checkCurrency(currency string) bool {
	values := []string{"rub", "usd", "eur"}
	for i := 0; i < 3; i++ {
		if currency == values[i] {
			return true
		}
	}
	fmt.Println("Неверный формат валюты!")
	return false
}

func getUserInfo() (string, float64, string) {
	var money float64
	var origCurrency, targetCurrency string
	for {
		fmt.Print("Введите исходную валюту(rub/usd/eur): ")
		fmt.Scan(&origCurrency)
		if !checkCurrency(origCurrency) {
			continue
		}
		break
	}
	for {
		fmt.Printf("Введите сумму в исходной валюте %s: ", origCurrency)
		_, err := fmt.Scan(&money)
		if err != nil {
			fmt.Println("Некорректная сумма!")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		if money < 0 {
			fmt.Println("Некорректная сумма!")
			continue
		}
		break
	}
	for {
		fmt.Print("Введите целевую валюту(rub/usd/eur): ")
		fmt.Scan(&targetCurrency)
		if origCurrency == targetCurrency {
			fmt.Println("Исходная и целевая валюта совпадают!")
			continue
		}
		if !checkCurrency(targetCurrency) {
			continue
		}
		break
	}
	return origCurrency, money, targetCurrency
}

func convert(money float64, origValue, targetValue string, convMap *conversionMap) float64 {
	return money * (*convMap)[origValue][targetValue]
}
