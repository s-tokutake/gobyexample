// 基本的なチャンネルの送受信はブロッキングです。
// しかし、`default`句のある`select`を使うことで
// _ノンブロッキング_な送受信やノンブロッキングで複数チャンネルを扱う`select`を実装できます。

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// これがノンブロッキングな受信です。`messages`の値が利用可能なら
	// `select`は`<-messages`の`case`を選択し値を取得します。
	// そうでないなら即座に`default`を選択します。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// ノンブロッキングな送信も同様です。ここでは`msg`は`messages`チャンネルに送信されません。
	// チャンネルがバッファなしであり、対応する受信がないからです。
	// それゆえ、`default`が選択されます。
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 複数の`case`を`default`の前に書くことでノンブロッキングで
	// 複数チャンネルを扱うselectを実装できます。
	// ここでは、`messages`と`signals`からの受信をノンブロッキングで行なっています。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
