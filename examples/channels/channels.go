// _チャンネル_は平行しているゴルーチン同士をつなぐパイプです。
// ひとつのゴルーチンがチャンネルに値を送信し、
// 別のゴルーチンがその値を受信できます。

package main

import "fmt"

func main() {

    // `make(chan val-type)`で新しいチャンネルを作ります。
    // チャンネルは送受信する値で型付けされます。
    messages := make(chan string)

    // チャンネルへの値の_送信_は`channel <-`という構文を使います。
    // ここでは新しいゴルーチンから`"ping"`を`messages`チャンネルに送信しています。
    go func() { messages <- "ping" }()

    // `<-channel`という構文でチャンネルから値を_受信_します。
    // ここでは、上で送信した`"ping"`メッセージを受信し、
    // プリントしています。
    msg := <-messages
    fmt.Println(msg)
}
