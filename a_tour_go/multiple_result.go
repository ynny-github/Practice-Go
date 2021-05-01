package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// 必ず2変数で受け取らなければならない
	// python のように一変数で受け取ったら、配列として扱われない
	// ☓ c := swap("a", "b")
	a, b := swap("a", "b")

	fmt.Println(a)
	fmt.Println(b)
}
