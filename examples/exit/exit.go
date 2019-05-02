// `os.Exit`を使えば特定のステータスでプログラムを終了できます。

package main

import "fmt"
import "os"

func main() {

	// os.Exit`を使う場合、`defer`は実行_されません_。したがって、
	// この`fmt.Println`は実行されません。
	defer fmt.Println("!")

	// ステータス3でプログラムを終了します。
	os.Exit(3)
}

// Cとは違い、Goは`main`で整数の戻り値を使ってステータスを示しません。
// 0でないステータスでプログラムを終えたいなら、`os.Exit`を使う必要があります。
