package main

import "fmt"

func main() {
	fmt.Println("counting")

	// LIFO で実行される
	// 処理がスタックに積まれる
	// Python や JS にはない文法であるため、新鮮に感じる
	// なぜこの文法が存在しているのだろう？
	for i := 0; i < 10; i++ {
		//
		defer fmt.Println(i)
	}

	fmt.Println("done.")
}
