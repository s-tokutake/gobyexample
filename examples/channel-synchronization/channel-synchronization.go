// チャンネルはゴルーチン間の同期実行に使えます。
// ここではゴルーチンの完了を待つ、受信によるブロッキングを使った例を示します。

package main

import "fmt"
import "time"

// これはゴルーチン内で実行する関数です。
// `done`チャンネルは別のゴルーチンに
// この関数が完了したことを通知します。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 値を送信して完了を通知します。
	done <- true
}

func main() {

	// ワーカーのゴルーチンを開始し、引数に通知用のチャンネルを渡します。
	done := make(chan bool, 1)
	go worker(done)

	// ワーカーからの通知を受信するまで待ちます。
	<-done
}
