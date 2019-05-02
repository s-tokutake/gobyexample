// [環境変数](https://ja.wikipedia.org/wiki/%E7%92%B0%E5%A2%83%E5%A4%89%E6%95%B0)
// は[Unixプログラムに構成情報を持ち込む](http://www.12factor.net/config)
// ための一般的な仕組みです。どのように環境変数の設定、取得、一覧をするのか見ていきましょう。

package main

import "os"
import "strings"
import "fmt"

func main() {

	// キー/値のペアを設定するには`os.Setenv`を使います。
	// キーの値を取り出すには`os.Getenv`を使います。
	// 環境にキーがなければ空の文字列を返します。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// 環境内のキー/値のペアをすべて一覧するには`os.Environ`を使います。
	// これは`KEY=value`という形式の文字列のスライスを返します。
	// `strings.Split`を使ってキーと値を取り出せます。
	// ここでは全てのキーを表示指定ます。
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
