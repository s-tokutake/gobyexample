// チャンネルを_閉じる_ことは、これ以上そのチャンネルを通じて
// 値が送信されることがない、ということを示します。
// これはチャンネルの受信側に完了を伝えるのに使えます。

package main

import "fmt"

// ここでは`jobs`チャンネルを使って`main()`ゴルーチンからワーカーの
// ゴルーチンへ処理が完了したことを通知します。
// ワーカーが処理しなければならないジョブがなくなったら
// ワーカーは`jobs`チャンネルを`close`します。
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    // これがワーカーのゴルーチンです。`j, more := <-jobs`で
    // `jobs`チャンネルから繰り返し受信をします。この特別なふたつの値の受信では、
    // `more`の値は`jobs`チャンネルが`close`されており
    // チャンネル内のすべての値が受信されていたらfalseを返します。
    // これを使ってワーカーがすべてのジョブを処理し終わったことを
    // `done`チャンネルに通知します。
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    // ここで`jobs`チャンネル経由でワーカーに
    // 3つのジョブを送信し、チャンネルを閉じます。
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    // 前の例で見た[同期](チャンネル同期)アプローチでワーカーを待ちます。
    <-done
}
