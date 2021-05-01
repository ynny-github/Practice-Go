package main

import "fmt"

func main() {
	sum_1 := 0
	// C や Java の書き方に近い
	for i := 0; i < 10; i++ {
		sum_1 += i
	}
	fmt.Println(sum_1)

	// 初期化と後処理ステートメントは省略できる
	// while はなく、初期化と後処理がない for が while と等価である
	// ほぼ同じ動作で両方存在する意味がないため、シンプルでいい
	sum_2 := 1
	for sum_2 < 1000 {
		sum_2 += sum_2
	}
	fmt.Println(sum_2)

	// ループ条件も省略すると無限ループになる
	sum_3 := 1
	for {
		sum_3 += 1
		if sum_3 == 10 {
			break
		}
	}
}
