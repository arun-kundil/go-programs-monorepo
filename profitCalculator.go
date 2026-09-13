package main

import (
	"fmt"
)

func main() {
	fmt.Println("Profit Calculator")

	revenue := getInputData("Revenue : ")
	expenses := getInputData("Expenses : ")
	taxRate := getInputData("Tax Rate : ")

	fmt.Println("Calculating Expenses...")

	ebt, profit, taxRatio := profitCalculator(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax : %.1f\n", ebt)
	fmt.Printf("Net Profit : %.1f\n", profit)
	fmt.Printf("Tax Ratio : %.1f\n", taxRatio)

}

func getInputData(infoText string) float64 {

	var inputData float64
	fmt.Print(infoText)
	fmt.Scan(&inputData)

	return inputData
}

func profitCalculator(revenue, expenses, taxRate float64) (float64, float64, float64) {

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio

}
