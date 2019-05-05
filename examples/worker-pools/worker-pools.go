// ゴルーチンとチャンネルを使った_ワーカープール_の実装を紹介します。

package main

import "fmt"
import "time"

// まずワーカーです。平行で実行します。`jobs`チャンネルからジョブを受信し、
// `results`チャンネルに結果を送信します。重いタスクをシュミレートするため、
// 処理の間に1秒待ちます。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// ワーカーのプールを使うためにジョブを送信し結果を集める必要があります。
	// このためにふたつのチャンネルを作ります。
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 3つのワーカーを立ち上げます。まだジョブがないのでブロックされます。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 5つの`jobs`を送信しチャンネルを`close`してこれ以上ジョブがないことを示します。
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 最後にすべての結果を受信します。
	for a := 1; a <= 5; a++ {
		<-results
	}
}
