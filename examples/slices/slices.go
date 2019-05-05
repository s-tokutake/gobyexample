// _スライス_はGoの重要なデータ型です。
// シーケンスに対して配列よりも強力なインターフェースを提供します。

package main

import "fmt"

func main() {

    // 配列と違い、スライスは含む要素だけで型付けされます(要素数ではなく)。
    // 長さが0でない空のスライスを作るには組み込みの`make`を使います。
    // ここでは長さ`3`の`string`のスライス(初期値は0値)を作っています。
    s := make([]string, 3)
    fmt.Println("emp:", s)

    // 配列のように設定、取得ができます。
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // `len`はスライスお長さを返します。
    fmt.Println("len:", len(s))

    // これらの基本的な操作に加え、スライスにはいくつかの操作があり、
    // 配列よりもリッチな使い方ができます。
    // まず、組み込みの`append`です。
    // これはひとつ以上の新しい値を含むスライスを返します。
    // 新しいスライス値を取得する`append`の戻り値を
    // 受け取る必要があります。
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // スライスは`copy`できます。
    // ここでは`s`と同じ長さの空のスライス`c`に`s`をコピーします。
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // スライスは`slice[low:high]`という構文の"スライス"演算子を
    // サポートします。この例では`s[2]`, `s[3]`,`s[4]`を
    // 取得します。
    l := s[2:5]
    fmt.Println("sl1:", l)

    // このスライスは`s[5]`までを取得します(`s[5]`自体は含まない)。
    l = s[:5]
    fmt.Println("sl2:", l)

    // このスライスは`s[2]`から終わりまでを取得します(`s[2]`を含む)。
    l = s[2:]
    fmt.Println("sl3:", l)

    // 1行でスライスの宣言と初期化を行えます。
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // スライスを使って多次元のデータ構造を構成できます。
    // 多次元配列とは違い、内部のスライスの長さをさまざまです。
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
