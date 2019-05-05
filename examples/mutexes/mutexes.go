// 前の例では単純なカウンタの状態を[アトミックな操作](atomic-counters)で管理する方法を示しました。
// より複雑な状態の場合は<em>[mutex](http://en.wikipedia.org/wiki/Mutual_exclusion)</em>を使って複数のゴルーチンから
// 安全にデータにアクセスできます。

package main

import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)

func main() {

    // この例では`state`はマップです。
    var state = make(map[int]int)

    // `mutex`で`state`に対するアクセスを同期させます。
    var mutex = &sync.Mutex{}

    // 何回読み取りと書き込みをしたのかを追跡します。
    var readOps uint64
    var writeOps uint64

    // 100のゴルーチンを開始しstateに対して読み取りを繰り返し実行します。
    // 各ゴルーチンで1ミリ秒に1回読み取りを行います。
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {

                // 読み取り1回につき、アクセスのためのキーを取り出し、
                // `mutex`を`Lock()`して`state`に対する排他アクセスを保証し、
                // 選択したキーの値を読み、
                // `mutex`を`Unlock()`して、
                // `readOps`をインクリメントします。
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddUint64(&readOps, 1)

                // 読み取り処理の間に少し待ちます。
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // 10のゴルーチンを開始して書き込みをシュミレートします。
    // 処理のパターンは読み取りと同じです。
    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // 1秒間、待って10のゴルーチンが`state`と`mutex`に対する処理をするのを待ちます。
    time.Sleep(time.Second)

    // 最終の処理回数を取得し表示します。
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)

    // 最後に`state`をロックして状態を表示します。
    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}
