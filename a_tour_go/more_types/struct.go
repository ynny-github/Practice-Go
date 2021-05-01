package main

import "fmt"

type Vertex struct {
	// ここにも大文字エクスポートの仕組みが働く
	// 小文字で宣言すると、class で protected, private のように
	// 変数を扱える
	X int
	Y int
}

func main() {
	// Vertex{1, 2} 位置で代入するなら、すべてのフィールドを初期化する必要あり
	// Vertex{X: 1, Y: 2} フィールドを指定して代入できる
	// Vertex{X: 1} フィールドしていないなら一部だけ初期化もできる

	v := Vertex{1, 2}
	// x の場合、x には参照できない
	fmt.Println(v.X, v.Y)

	p := &v
	// 構造体のポインタの場合、* は必須ではない
	p.X = 10
	(*p).Y = 11
	fmt.Println(v.X, v.Y)
}
