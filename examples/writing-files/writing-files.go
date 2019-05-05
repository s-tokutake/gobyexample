// Goのファイル書き込みは前に見た読み込みと同じようなパターンに従います。

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// まず文字列(つまり単なるバイト)をファイルにダンプします。
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// もっと細かく書き込みを行うには書き込み用のファイルを開きます。
	f, err := os.Create("/tmp/dat2")
	check(err)

	// ファイルを開いたら直後に`Close`をdefer付きで実行します。
	defer f.Close()

	// バイトのスライスを`Write`できます。
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// `WriteString`も可能です。
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// `Sync`を使うと安定的なストレージに書き込みをフラッシュします。
	f.Sync()

	// `bufio`は前に紹介したバッファ付き読み込みに加え
	// バッファ付きの書き込みを提供します。
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// `Flush`を使うとすべてのバッファされた操作が適用されます。
	w.Flush()

}
