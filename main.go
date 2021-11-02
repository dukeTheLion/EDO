package main

import (
	"fmt"
	"math"
)

var facts [40]int

func factorial(n int) (res int) {
	if facts[n] != 0 {
		res = facts[n]
		return res
	}

	if n > 0 {
		res = n * factorial(n-1)
		return res
	}

	return 1
}

func h(a float64, b float64, n int) float64 {
	return (b - a) / float64(n)
}

func divisor(a float64, b float64, n int) []float64 {
	aux := make([]float64, 0, n+1)
	step := h(a, b, n)

	aux = append(aux, a)

	for i := 1; i < n; i++ {
		aux = append(aux, a+step*float64(i))
	}

	aux = append(aux, b)

	return aux
}

func taylorSeries(x float64, xi float64, y ...float64) float64 {
	yn := y[0]

	for i := 1; i < len(y); i++ {
		aux := math.Pow(x-xi, float64(i))
		aux /= float64(factorial(i))
		aux *= y[i]

		yn += aux
	}

	return yn
}

func main() {
	yn := []float64{1, 4, 13, 40, 119}

	for i, v := range yn {
		fmt.Print("y")
		for j := 0; j < i; j++ {
			fmt.Print("´")
		}
		fmt.Printf("(0) = %.0f     ", v)
	}

	fmt.Printf("\nf(0.5) = %.10f\n", taylorSeries(0.5, 0, yn...))
}
