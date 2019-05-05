// [前の](range)例では`for`と`range`の基本的なデータ構造に対する反復処理を紹介しました。
// チャンネルからの受信でもこの構文を使うことができます。

package main

import "fmt"

func main() {

	// `queue`チャンネルの2つの値を反復処理します。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// `range`を使って`queue`から受信した値を反復処理します。
	// 上でチャンネルを閉じているため2つの要素を受信して反復は終了します。
	for elem := range queue {
		fmt.Println(elem)
	}
}
