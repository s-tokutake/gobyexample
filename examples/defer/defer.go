// _defer_はある関数呼び出しがプログラムの実行の後に行われることを保証します。
// クリーアップ目的で使われるのが一般的です。
// `defer`がよく使われるのは他の言語で`ensure`や`finally`が使われる場合です。

package main

import "fmt"
import "os"

// ファイルを作成し書き込みをし、終わったらファイルを閉じたいとします。
// ここでは`defer`を使ってこれを実現します。
func main() {

    // `createFile`でファイルオブジェクトを取得した直後、
    // `closeFile`でファイルを閉じる処理をdeferを使って実行します。
    // これによって`closeFile`は`writeFile`が完了した後、This will be executed at the end
    // この関数(`main`)の最後に実行されます。
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
    fmt.Println("closing")
    f.Close()
}
