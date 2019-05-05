// Goは_複数戻り値_をサポートします。
// これはGoでは慣用的に使われます。
// 例えば、結果とエラーを両方返す場合です。

package main

import "fmt"

// この関数のシグネチャの`(int, int)`は
// ふたつの`int`を返す関数であることを示します。
func vals() (int, int) {
    return 3, 7
}

func main() {

    // _多重代入_でふたつ異なる戻り値を取得します。
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // 一部の戻り値だけ欲しいならブランク識別子`_`を使えます。
    _, c := vals()
    fmt.Println(c)
}
