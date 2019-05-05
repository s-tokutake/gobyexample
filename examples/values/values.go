// Goにはさまざまな値型があります。文字列、整数、浮動小数点数、ブーリアンなどです。
// いくつか基本的な例を示します。

package main

import "fmt"

func main() {

    // 文字列です。`+`で結合できます。
    fmt.Println("go" + "lang")

    // 整数と浮動小数点数です。
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // ブーリアンです。おなじみの演算子が使えます。
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
