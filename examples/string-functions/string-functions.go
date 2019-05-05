// `strings`パッケージは文字列関連の便利な関数を提供します。
// ここではいくつかの例を示します。
package main

import s "strings"
import "fmt"

// `fmt.Println`にエイリアスをつけます。たくさん使うからです。
var p = fmt.Println

func main() {

    // `strings`で使える関数のサンプルです。
    // 文字列オブジェクトのメソッドではなく`strings`パッケージの関数なので
    // 第一引数に文字列を渡す必要があります。
    // このパッケージのドキュメント[`strings`](http://golang.org/pkg/strings/)で他の関数も確認できます。
    p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
    p()

    // `strings`の関数ではないがここで紹介する価値があるのが、
    // バイトでの文字列の長さを取得する方法とインデックスでバイトを取得する方法です。
    p("Len: ", len("hello"))
    p("Char:", "hello"[1])
}

// `len`もインデックスもバイトレベルで働きます。
// GoはUTF-8でエンコードされた文字列を使うので多くの場合そのままで便利です。
// マルチバイトの文字を扱う場合はエンコーディングを意識した処理をした場合があります。
// 詳細は[strings, bytes, runes and characters in Go](https://blog.golang.org/strings)を見てください。
