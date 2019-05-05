// チャンネルを関数のパラメータに使う場合
// チャンネルが送信専用か受信専用なのかを定義できます。
// これによってプログラムがよりしっかりとタイプセーフになります。

package main

import "fmt"

// この`ping`関数は送信のみを行うチャンネルをパラメータとして受け取ります。
// このチャンネルで受信をしようとするとコンパイルエラーになります。
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// `pong`関数は受信のチャンネル(`pings`)と
// 送信のチャンネル(`pongs`)を受け付けます。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
