// Goは[正規表現](http://en.wikipedia.org/wiki/Regular_expression)を組み込みでサポートします。
// ここでは正規表現に関するいくつかの一般的な例を表示します。

package main

import "bytes"
import "fmt"
import "regexp"

func main() {

    // これは文字列がパターンにマッチするかどうかをテストします。
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    // 上の例では文字列のパターンを直接使いました。
    // 他の処理も実施するには最適化された`Regexp`構造体を`Compile`する必要があります。
    r, _ := regexp.Compile("p([a-z]+)ch")

    // この構造体では多くのメソッドが使えます。
    // これは先に見たマッチのテストです。
    fmt.Println(r.MatchString("peach"))

    // マッチした文字列を見つけます。
    fmt.Println(r.FindString("peach punch"))

    // 最初のマッチを見つけますが、マッチした文字列ではなく
    // マッチした文字列の最初と最後のインデックスを返します。
    fmt.Println(r.FindStringIndex("peach punch"))

    // `Submatch`には全体のパターンマッチの結果と
    // その中に含まれているサブマッチの結果を返します。
    // この例なら`p([a-z]+)ch`と`([a-z]+)`の両方の情報を返します。
    fmt.Println(r.FindStringSubmatch("peach punch"))

    // 同じようにマッチとサブマッチのインデックスの情報を返します。
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

    // `All`が付いている場合は最初のマッチだけではなく全てのマッチが対象になります。
    // この例ではすべてのマッチを見つけています。
    fmt.Println(r.FindAllString("peach punch pinch", -1))

    // この`All`は上と同じように他の関数にもあります。
    fmt.Println(r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))

    // 2つ目の引数に負でない整数を与えるとマッチの数を制限できます。
    fmt.Println(r.FindAllString("peach punch pinch", 2))

    // 上の例の関数は`MatchString`のような名前で文字列の引数を取りました。
    // `[]byte`の引数を取る関数もあります。関数名から`String`を削除します。
    fmt.Println(r.Match([]byte("peach")))

    // 正規表現の定数を作成する場合、`Compile`の変種である`MustCompile`を使います。
    // `Compile`はふたつの値を返すため定数には使えます。
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)

    // `regexp`パッケージは文字列の一部を他の文字列に置換するのにも使われます。
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    // `Func`を使えばマッチした文字列に対して引数に与えた関数を適用できます。
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}
