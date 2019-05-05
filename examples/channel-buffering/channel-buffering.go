// デフォルトではチャンネルは_バッファなし_です。
// つまり、送信(`chan <-`)を受け付けるのは対応する受信(`<- chan`)が
// 送信された値を受け取る準備ができているときだけです。
// _バッファありチャンネル_の場合、対応する受信がなくても指定した数の値を
// 受け付けることができます。

package main

import "fmt"

func main() {

	// `make`でバッファが2の文字列のチャンネルを作ります。
	messages := make(chan string, 2)

	// バッファありなので対応する平行の受信がなくても
	// これらの送信ができます。
	messages <- "buffered"
	messages <- "channel"

	// 受信も普通にできます。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
