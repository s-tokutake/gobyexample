// _select_を使うと複数のチャンネル操作を待てます。
// selectとゴルーチン、チャンネルの組み合わせはGoの強力な特徴です。

package main

import "time"
import "fmt"

func main() {

    // この例ではふたつのチャンネルをselectします。
    c1 := make(chan string)
    c2 := make(chan string)

    // それぞれのチャンネルは値を受信します。
    // ブロッキングのRPC操作を平行ゴルーチンで実行しているのをシュミレートするため
    // いくらかの時間が経過したあと受信を行います。
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    // `select`を使って両方の値を同時に待ちます。
    // 値が到着したら表示します。
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
