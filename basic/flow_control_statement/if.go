package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// python のように括弧を書かなくてよく書きやすい
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// 読み慣れないが、スコープを限定できるメリットが有る
	// else までスコープの範囲に入る
	// else if でも使える
	// 続いている if ないなら使える
	if v := math.Pow(x, n); v < lim {
		return v
	} else if y := math.Pow(x, n); y > lim {
		fmt.Println(v)
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
