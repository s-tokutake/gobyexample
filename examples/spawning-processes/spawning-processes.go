// GoプログラムでGoではないプロセスを立ち上げる場合があります。
// 例えば、このサイトの構文ハイライトはGoのプログラムから[`pygmentize`](http://pygments.org/)プロセスを起動することで[実装](https://github.com/mmcgrana/gobyexample/blob/master/tools/generate.go)されています。
// ここではGoからプロセスを起動する方法についていくつかの例を示します。

package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {

	// 引数や入力を取らず、単純に標準出力に出力するだけのシンプルなコマンドから紹介します。
	// `exec.Command`を使って外部プロセスを表すオブジェクトを作成します。
	dateCmd := exec.Command("date")

	// `.Output`はコマンド実行の一般的なケースを扱います。
	// 完了を待ち、出力をまとめます。
	// エラーがなければ、`dateOut`にはバイト型の日時の情報が設定されます。
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// 次は少し込み入ったケースです。
	// `stdin`を使って外部プロセスにデータをパイプします。
	// そして`stdout`から結果を取得します。
	grepCmd := exec.Command("grep", "hello")

	// 明示的に入力/出力のパイプを捉えるようにします。
	// プロセスを開始し、入力を書き込んで結果の出力を読み取ります。
	// そして、プロセスが完了するのを待ちます。
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// 上の例ではエラーのチェックをしていませんが、
	// いつもの`if err != nil`パターンでチェックできます。
	// `StdoutPipe`の結果だけを取得することもできますが、
	// まったく同じ方法で`StderrPipe`も取得できます。
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 外部のコマンドを起動する場合はひとつのコマンドライン文字列を与えるのではなく、
	// コマンドと引数の配列を明示的に提供する必要があります。
	// 単一の文字列でコマンドを実行したいなら`bash`の`-c`オプションで実現できます。
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
