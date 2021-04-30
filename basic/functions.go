package main

import "fmt"

// func <function name>(args) type {}
// func <function name>(args) (type, type...) {}
// args pattern 1
// arg1 type, arg2 type...

// args pattern 2
// If arg1 and arg2 are same type, you can write this.
// arg1, arg2 type

func add_pattern_1(x int, y int) int {
	return x + y
}

func add_pattern_2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add_pattern_1(1, 2))
	fmt.Println(add_pattern_2(1, 2))
}
