// Goの状態管理の第一の仕組みはチャンネルを使った通信です。
// この例は前に[ワーカープール](worker-pools)で見ました。
// しかし、状態管理には他にも方法があります。
// ここでは`sync/atomic`パッケージを使って
// 複数のゴルーチンからアクセスされる_アトミックカウンタ_を示します。

package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	// 未割り当ての整数値を使って常に(常に正の整数)の
	// カウンタを表現します。
	var ops uint64

	// 平行の更新をシュミレートするため50のゴルーチンを開始し
	// 1ミリ秒に1回、カウンタをインクリメントします。
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// アトミックにカウンタをインクリメントするために
				// `&`構文を使って`AddUint64`に`ops`カウンタのメモリアドレスを
				// 渡します。
				atomic.AddUint64(&ops, 1)

				// インクリメント処理の間、少し待ちます。
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 1秒、opsのインクリメントが進むのを待ちます。
	time.Sleep(time.Second)

	// 他のゴルーチンによって更新されているカウンタに安全にアクセスするため
	// `LoadUint64`を使って`opsFinal`に現在の値のコピーを取り出します。
	// 上と同じようにこの関数にもメモリアドレス`&ops`を与え、値を取り出します。
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
