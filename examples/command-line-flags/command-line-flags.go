// [_コマンドラインフラグ_](http://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)はコマンドラインプログラムの
// オプションを定義する一般的な方法です。例えば、`wc -l`では`-l`がコマンドラインフラグです。

package main

// Goは`flag`パッケージを提供し、コマンドラインフラグのパースをサポートします。
// ここではこのパッケージを使ってサンプルのコマンドラインプログラムを実装します。
import "flag"
import "fmt"

func main() {

    // 基本的なフラグの宣言は文字列、整数、ブーリアンを利用します。
    // ここでは文字列のフラグ`word`を宣言し、デフォルトの値を`"foo"`にして、
    // 簡単な説明を設定しています。
    // この`flag.String`関数は文字列のポインタ(文字列値そのものではない)を返します。
    // このポインタの使い方は後述します。
    wordPtr := flag.String("word", "foo", "a string")

    // `word`と同じように`numb`と`fork`フラグを定義します。
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    // プログラムのどこかに宣言した既存の変数を使ってオプションを宣言することもできます。
    // フラグ宣言関数にポインタを渡す必要があることに注意してください。
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // すべてのフラグが宣言できたら`flag.Parse()`を呼び
    // コマンドラインのパースを実行します。
    flag.Parse()

    // ここではパースされたオプションと後続する可能性のある引数をダンプします。
    // `*wordPtr`というかたちでポインタをデリファレンスして実際の
    // オプション値を取得する必要があることに注意してください。
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}
