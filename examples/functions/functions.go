// _関数_はGoの中心です。ここではいくつかの異なる例を
// 通じて関数を学びます。

package main

import "fmt"

// ふたつの`int`を受け取ってその合計を返す関数です。
func plus(a int, b int) int {

    // Goでは明示的なreturnが必要です。最後の式の値を
    // を自動的に返すことはしません。
    return a + b
}

// 同じ型の連続した複数のパラメータを取る場合は、
// 最後のパラメータまで型の名前を省略できます。
// 最後のパラメータで型を宣言します。
func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {

    // 関数呼び出しはおなじみの`name(args)`という形式です。
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}
