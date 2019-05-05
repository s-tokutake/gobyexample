// <em>[律速](http://en.wikipedia.org/wiki/Rate_limiting)</em>はリソース利用の制御とサービスの質の維持にとって重要な仕組みです。
// Goはゴルーチンとチャンネルとtickers](tickers)を使って律速をエレガントにサポートします。

package main

import "time"
import "fmt"

func main() {

    // 最初は基本的な例です。処理するリクエスト数を制限したいとします。
    // ここではrequestsという名のチャンネルでリクエストを処理します。
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    // `limiter`チャンネルは200ミリ秒ごとに値を受信します。
    // これが律速の調整弁いなります。
    limiter := time.Tick(200 * time.Millisecond)

    // リクエストを処理する前のlimiter`チャンネルからの受信でブロックされることで
    // 200ミリ秒に1リクエスト処理するという制限がかかります。
    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }

    // 全体の律速は一定に保ったままリクエストが短い間バーストするのはを許容したい
    // 場合があります。これは`limiter`チャンネルをバッファありにすることで実現できます。
    // この`burstyLimiter`チャンネルは3回までバーストを許容します。
    burstyLimiter := make(chan time.Time, 3)

    // チャンネルを埋めて許容されたバーストを表します。
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    // 200ミリ秒ごとに上限の3つまで新しい値を`burstyLimiter`に追加します。
    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

    // ここで5つのリクエストをシュミレートします。最初の3つは
    // `burstyLimiter`のバースト能力の恩恵を受けます。
    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}
