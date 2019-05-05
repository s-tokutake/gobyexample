// Goは構造体上で_メソッド_を宣言できます。

package main

import "fmt"

type rect struct {
    width, height int
}

// この`area`メソッドは`*rect`の_レシーバ型_を持ちます。
func (r *rect) area() int {
    return r.width * r.height
}

// メソッドはポインタレシーバ型でも値レシーバ型でも定義できます。
// これは値レシーバの例です。
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    // ここでふたつのメソッドを呼び出します。
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    // Goはメソッド呼び出しのとき、値とポインタを自動的に変換します。
    // メソッド呼び出しでのコピーを避けるため、または、
    // メソッドがレシーバとなる構造体を変えられるようにするために、
    // ポインタのレシーバ型を使うことができます。
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
