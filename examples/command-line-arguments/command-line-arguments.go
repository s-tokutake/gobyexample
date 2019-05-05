// [_コマンドライン引数_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)はプログラムの実行を
// パラメータ化する一般的な方法です。
// 例えば、`go run hello.go`は`run`と
// `hello.go`というふたつの引数を`go`プログラムに渡しています。

package main

import "os"
import "fmt"

func main() {

	// `os.Args`は生のコマンドライン引数へのアクセスを提供します。
	// このスライスの最初の値がプログラムへのパスであり、
	// `os.Args[1:]`にプログラムへの引数が含まれることに注意してください。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// 個別の引数はふつうのインデックスで取得できます。
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
