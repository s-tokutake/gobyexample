// 前の例では[mutex](mutexes)を使って明示的にロックをして
// 共有する状態に対する複数ゴルーチンからのアクセスを同期しました。
// 別の選択肢としてゴルーチンとチャンネルの組み込みの同期機構を使っても
// 同じことを実現できます。このチャンネルを使ったアプローチは
// 通信をしてひとつだけのゴルーチンがデータを持つことでメモリの共有を実現する、
// というGoのアイディアと整合します。

package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

// この例では状態はひとつのゴルーチンに所有されます。
// これによって平行アクセスでデータが汚されないようになります。
// 状態を読み書きするためには、他のゴルーチンはオーナーのゴルーチンにメッセージを送信し
// 返答をもらいます。この`readOp`と`writeOp`という構造体は
// この読み書きの要求とオーナーゴルーチンの応答をカプセル化します。
type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func main() {

    // 前回同様、実行する操作の回数を数えます。
    var readOps uint64
    var writeOps uint64

    // `reads`と`writes`は読み書きの要求を発行するために他のゴルーチンから使われます。
    reads := make(chan *readOp)
    writes := make(chan *writeOp)

    // `state`を所有するゴルーチンです。前の例と同様`state`はマップです。
    // しかし、今回は状態を持つゴルーチン内で閉じています。
    // このゴルーチンは`reads`チャンネルと`writes`チャンネルに対して
    // selectを繰り返し、到着した要求に応答します。
    // 応答はまず、要求された処理を行い、その後に`resp`チャンネルに値を送信して
    // 成功したこと(`reads`のcase文の場合は望みの値)を通知します。
    go func() {
        var state = make(map[int]int)
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()

    // 100のゴルーチンを開始し`reads`チャンネル経由で状態を持つゴルーチンに対して読み取りを要求します。
    // それぞれの読み取りで`readOp`を作り`reads`チャンネル経由で送信し、
    // `resp`チャンネルから結果を受信します。
    for r := 0; r < 100; r++ {
        go func() {
            for {
                read := &readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // 書き込みも10のゴルーチンで行います。
    // 内容は読み取りと似ています。
    for w := 0; w < 10; w++ {
        go func() {
            for {
                write := &writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // ゴルーチンの処理を1秒待地ます。
    time.Sleep(time.Second)

    // 最後に処理回数を取り出して表示します。
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}
