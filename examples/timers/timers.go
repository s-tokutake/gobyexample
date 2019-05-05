// Goのコードを未来のある時点でで実行したい場合があります。
// また、一定間隔で実行したい場合もあります。Goの組み込みの_timer_と_ticker_を
// 使うとこのふたつは簡単に実現できます。まずはtimerを紹介し、次に[tickers](tickers)を紹介します。

package main

import "time"
import "fmt"

func main() {

    // タイマーは未来の単一のイベントを表します。
    // タイマーにどのくらい待ってほしいかを指定すると、
    // タイマーはそのタイミングで通知をくれるチャンネルを提供してくれます。
    // ここでは2秒待ちます。
    timer1 := time.NewTimer(2 * time.Second)

    // `<-timer1.C`はタイマーのチャンネル`C`で
    // タイマーが時間が来たことを示す値を送信するまでブロックします。
    <-timer1.C
    fmt.Println("Timer 1 expired")

    // 単に待ちたいだけなら`time.Sleep`が使えます。
    // タイマーが便利な理由のひとつは時間が来る前にキャンセルできるからです。
    // ここでその例を示します。
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}
