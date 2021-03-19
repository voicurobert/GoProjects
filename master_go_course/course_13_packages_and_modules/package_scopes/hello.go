package package_scopes

import "fmt"

const boilingPoint = 100

func Hello() {
	fmt.Println("Hi")
}

func toFahrenheit(t float64) float64 {
	return t*1.8 + 32
}
