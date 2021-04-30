package main

// fomatter によって自動的に factored インポートステートメントに直される
// The import statement is fixed factored import statement by golang formatter.
// import "fmt"
// import "math" のようにも書ける
import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
