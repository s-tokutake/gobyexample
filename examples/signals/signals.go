// Goで[Unixシグナル](http://en.wikipedia.org/wiki/Unix_signal)を賢く扱いたい場合があります。
// 例えば、`SIGTERM`を受信した場合、サーバをきれいに停止したり、
// `SIGINT`を受信したら入力の処理を停止するコマンドラインツールを作ったりする場合です。
// ここではチャンネルを使ってシグナルを処理する方法を示します。

package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

    // Goのシグナル通知は`os.Signal`値をチャンネルに送信することで実現します。
    // `os.Signal`を受け取るためのチャンネルを作ります。
    // (また、プログラムが終了できることを通知するチャンネルも作ります)
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    // `signal.Notify`でチャンネルを登録し、定義したシグナルの通知を受信するようにします。
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    // このゴルーチンがブロックしながらシグナルの受信をします。
    // シグナルを受信したら表示してプログラムが終了できることを通知します。
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    // 期待されるシグナル(`done`に値を送っている上のゴルーチンが示す)を得るまで
    // プログラムは停止し、その後、終了します。
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
