// Goは従来の`printf`の文字列フォーマッティングに対して素晴らしいサポートを提供します。
// ここではよくある文字列フォーマッティングの例を示します。

package main

import "fmt"
import "os"

type point struct {
    x, y int
}

func main() {

    // Goはいくつかの表示の"verbs"を提供しGoの一般的な値をフォーマットできるようにしています。
    // ここでは`point`構造体を表示しています。
    p := point{1, 2}
    fmt.Printf("%v\n", p)

    // 値が構造体の場合、`%+v`を使うと構造体のフィールド名が含まれます。
    fmt.Printf("%+v\n", p)

    // `%#v`は対象の値のGoの構文表示です。
    // つまり、対象の値を生成したソースコードスニペットを表示します。
    fmt.Printf("%#v\n", p)

    // 値の型を出力するなら`%T`を使います。
    fmt.Printf("%T\n", p)

    // ブーリアンの場合は直截的です。
    fmt.Printf("%t\n", true)

    // 整数のフォーマットにはたくさんのオプションがあります。
    // `%d`が標準的です。基数10でフォーマッティングします。
    fmt.Printf("%d\n", 123)

    // バイナリで表示するにはこれです。
    fmt.Printf("%b\n", 14)

    // これは与えられた整数に対応するキャラクタを表示します。
    fmt.Printf("%c\n", 33)

    // `%x`はHEXエンコーディングです。
    fmt.Printf("%x\n", 456)

    // 浮動小数点数にも多くのオプションがあります。
    // 基本的な小数表記なら`%f`を使います。
    fmt.Printf("%f\n", 78.9)

    // `%e`と`%E`は指数表記(わずかに異なるバージョン)で浮動小数点数を表します。
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)

    // 文字列は`%s`を使って基本的な出力をします。
    fmt.Printf("%s\n", "\"string\"")

    // To double-quote strings as in Go source, use `%q`.
    fmt.Printf("%q\n", "\"string\"")

    // As with integers seen earlier, `%x` renders
    // the string in base-16, with two output characters
    // per byte of input.
    fmt.Printf("%x\n", "hex this")

    // To print a representation of a pointer, use `%p`.
    fmt.Printf("%p\n", &p)

    // When formatting numbers you will often want to
    // control the width and precision of the resulting
    // figure. To specify the width of an integer, use a
    // number after the `%` in the verb. By default the
    // result will be right-justified and padded with
    // spaces.
    fmt.Printf("|%6d|%6d|\n", 12, 345)

    // You can also specify the width of printed floats,
    // though usually you'll also want to restrict the
    // decimal precision at the same time with the
    // width.precision syntax.
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

    // To left-justify, use the `-` flag.
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

    // You may also want to control width when formatting
    // strings, especially to ensure that they align in
    // table-like output. For basic right-justified width.
    fmt.Printf("|%6s|%6s|\n", "foo", "b")

    // To left-justify use the `-` flag as with numbers.
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

    // So far we've seen `Printf`, which prints the
    // formatted string to `os.Stdout`. `Sprintf` formats
    // and returns a string without printing it anywhere.
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)

    // You can format+print to `io.Writers` other than
    // `os.Stdout` using `Fprintf`.
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
