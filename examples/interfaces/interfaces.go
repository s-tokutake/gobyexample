// _インターフェース_はメソッドのシグネチャの名前付きコレクションです。

package main

import "fmt"
import "math"

// これは図形のインターフェースです。
type geometry interface {
	area() float64
	perim() float64
}

// この例では`rect`型と`circle`型に上のインターフェースを実装します。
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Goでインターフェースを実装するには、
// 単にインターフェースのすべてのメソッドを実装すればいいだけです。
// ここでは`rect`に`geometry`を実装します。
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// circle`に実装します。
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 変数がインターフェース型の場合、その名前付きインターフェースのメソッドを呼び出せます。
// ここでは、この利点を生かしたどのような`geometry`にも使える`measure`関数を定義します。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// `circle`も`rect`も`geometry`インターフェースを実装しています。
	// したがって`measure`の引数として渡せます。
	measure(r)
	measure(c)
}
