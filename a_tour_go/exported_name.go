package main

import (
	"fmt"
	"math"
)

func main() {
	// 初めが大文字で始める変数、関数は外部に公開される。
	// 補完を見ても、すべての候補が大文字で始まっている。
	// 明示的に export を書かなくて良いのは新鮮だ。
	fmt.Println(math.Pi)
}
