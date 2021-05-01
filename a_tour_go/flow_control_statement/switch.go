package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	// なぜ while を for に融合したのに、if と switch は共存している？
	// c のように break を明示しなくても、break してくれる
	// 逆に、c のようにわざと break せずに下の処理に進ませる技はできない。
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
