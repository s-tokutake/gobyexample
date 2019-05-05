// [_SHA1ハッシュ_](http://en.wikipedia.org/wiki/SHA-1)はバイナリや文字列の短い識別子を計算するのによく使われます。
// 例えば、[git](http://git-scm.com/)はSHA1を広く活用してファイルとディレクトリのバージョンを特定します。
// ここではGoでSHA1ハッシュを計算する方法を示します。

package main

// Goは`crypto/*`パッケージでいくつかのハッシュ関数を実装しています。
import "crypto/sha1"
import "fmt"

func main() {
    s := "sha1 this string"

    // ハッシュ生成のパターンは`sha1.New()`を実行し、
    // `sha1.Write(bytes)`を実行し、最後に`sha1.Sum([]byte{})`を実行します。
    // まず新しいハッシュを作成します。
    h := sha1.New()

    // `Write`はバイトを受け取ります。文字列`s`なら`[]byte(s)`でバイトにします。
    h.Write([]byte(s))

    // ここでバイトのスライスでハッシュ結果を取得します。
    // `Sum`の引数は既存のバイトスライスを追記するのに使われます。
    // 通常は不要です。
    bs := h.Sum(nil)

    // SHA1の値はHEXでで表示されます。例えば、gitのコミットです。
    // `%x`を使ってハッシュの結果を文字列で表示しています。
    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}
