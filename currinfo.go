package main

import "fmt"

func PrintCurrencyInfo(curr, ctry, code string) {
	fmt.Println("The", ctry, "currency is the", curr, "(", code, ")")
}

// func main() {
// 	var currency string = "US Dollar"
// 	var country string = "United States"
// 	code := "USD"
// 	PrintCurrencyInfo(currency, country, code)
// }