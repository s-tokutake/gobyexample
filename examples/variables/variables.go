// Goでは、_変数_は明示的に宣言されます。コンパイラはこれを使って、
// 関数呼び出しの型の正しさをチェックなどをします。

package main

import "fmt"

func main() {

    // `var`はひとつ以上の変数を宣言します。
    var a = "initial"
    fmt.Println(a)

    // 複数の変数を一度に宣言できます。
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Goは初期化された変数の型を推論します。
    var d = true
    fmt.Println(d)

    // 初期値がセットされない変数は_ゼロ値_です。例えば、
    // `int`のゼロ値は`0`です。
    var e int
    fmt.Println(e)

    // `:=`という構文は変数の宣言と初期化の簡易表記です。
    // この例は、`var f string = "short"`と同じです。
    f := "short"
    fmt.Println(f)
}
