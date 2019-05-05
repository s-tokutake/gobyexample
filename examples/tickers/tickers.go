// [タイマー](timers)は未来のあるひとつの時点で何かを実行したいときに使います。
//  _ticker_は一定間隔で定期的に何かを実行したい場合に使います。
// ここでは停止するまで定期的に動作するtickerを紹介します。

package main

import "time"
import "fmt"

func main() {

    // tickerはタイマーと同様の仕組みを使います。値を送信するチャンネルです。
    // ここではチャンネルに`range`で反復し、500ミリ秒ごとに値を受信しています。
    ticker := time.NewTicker(500 * time.Millisecond)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()

    // タイマーと同じように停止できます。
    // 一度止めるとチャンネルから値を受信しなくなります。
    // ここでは1600ミリ秒後に停止しています。
    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    fmt.Println("Ticker stopped")
}
