// 多くのGoプログラムでファイルの読み書きは基本的なタスクです。
// まずはファイルの読み取りを紹介します。

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// ファイルの読み取りにはエラーのチェックが必要です。
// このヘルパ関数でエラーチェックを簡素にします。
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// もっとも基本的なファイル読み取りはファイルの内容全体を
	// メモリに読み込むことでしょう。
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// ファイルの読み込み方や読む部分を制御したい場合もあります。
	// この場合はまずファイルを`Open`して`os.File`を取得します。
	f, err := os.Open("/tmp/dat")
	check(err)

	// ファイルの先頭から数バイト読みます。
	// ５バイトまで読み、実際にどのくらい読んだのかを表示します。
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// ファイルの既知の位置を`Seek`してそこから`Read`します。
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// `io`パッケージはファイル読み込みに便利な関数を提供します。
	// 例えば上の例は`ReadAtLeast`を使うとより堅牢になります。
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 組み込みの巻き戻しはありません。`Seek(0, 0)`を使えば実現できます。
	_, err = f.Seek(0, 0)
	check(err)

	// `bufio`パッケージが実装するのはバッファ付きの読み取りで
	// 細かい多数の読み込みを効率的に処理できます。
	// また、別の便利な読み取りメソッドも提供します。
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// 完了したら閉じます(普通は`Open`の直後に`defer`を使って
	// 実行をスケジュールします)。
	f.Close()

}
