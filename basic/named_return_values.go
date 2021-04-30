package main

import "fmt"

// 型に指定した変数名に一致した変数を return ステートメント
// の横に書かなくても返せる
// 公式曰く、短い関数でのみ使用するべきと言っている。
// The official of go lang said that named return values should be used when a function is short.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
