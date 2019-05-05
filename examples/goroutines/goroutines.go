// _ゴルーチン_は軽量なスレッドの実行です。

package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// `f(s)`という関数呼び出しがあります。
	// この普通の呼び出し方では同期で実行されます。
	f("direct")

	// ゴルーチンで実行するなら`go f(s)`と書きます。
	// この場合、呼び出しをしたゴルーチンと平行で実行されます。
	go f("goroutine")

	// 無名関数の呼び出しもゴルーチンで開始できます。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 上のふたつの関数が非同期で別々のゴルーチンで実行されています。
	// ここで実行が止まります。
	// `Scanln`でプログラムが終了する前にキーを押下する必要があるためです。
	fmt.Scanln()
	fmt.Println("done")
}
