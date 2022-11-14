package main

import (
	"fmt"
)

func main() {
	//formula - C = K - 273

	fmt.Printf("O valor da conversão de 100 kelvin para %.2f celcius.", convert_kelvi_for_celsius(100))
	fmt.Println("")
	fmt.Printf("O valor da conversão de 200 kelvin para %.2f celcius.", convert_kelvi_for_celsius(200))
	fmt.Println("")
	fmt.Printf("O valor da conversão de 300 kelvin para %.2f celcius.", convert_kelvi_for_celsius(300))
	fmt.Println("")

}

func convert_kelvi_for_celsius(value float64) float64 {
	result := value - 273
	fmt.Print(result)
	return result
}
