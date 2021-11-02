package main

import (
	"fmt"
	"math"
)

func factorial(n int)(result int) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func taylorSeries(x float64, xi float64, y ...float64) float64 {
	yn := y[0]

	for i := 1; i < len(y); i++ {
		aux := math.Pow(x - xi, float64(i))
		aux /= float64(factorial(i))
		aux *= y[i]

		yn += aux
	}

	return yn
}

func main() {
	fmt.Printf("%.40f", taylorSeries(0.5, 0, 1, 4, 13, 40, 119))

}
