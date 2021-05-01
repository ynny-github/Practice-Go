package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	before_z := 1.1

	for before_z-z > 0.001 {
		// z は減少していくため、before のが大きい
		before_z = z
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println(math.Sqrt(2), Sqrt(2))
}
