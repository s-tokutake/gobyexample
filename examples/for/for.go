// Goではループは`for`だけです。
// `for`ループのの基本的な使い方を説明します。

package main

import "fmt"

func main() {

    // もっとも基本的な単一の条件です。
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // おなじみの、初期化/継続条件/後続処理 の`for`ループです。
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    // 条件のない`for`は、`break`するか関数から`return`するまで繰り返されます。
    for {
        fmt.Println("loop")
        break
    }

    // `continue`を使えば次の反復に進めます。
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
