// 前の例で[外部コマンドの起動](spawning-processes)を紹介しました。
// これは外部プロセスが実行中のGoプロセスにアクセスする必要がある場合に使います。
// しかし、単に既存のGoのプロセスを別の(Goではないかもしれない)プロセスに置き換えたいだけの場合もあります。
// この場合はGoの<a href="http://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>関数の実装を使います。

package main

import "syscall"
import "os"
import "os/exec"

func main() {

    // ここでは`ls`を実行します。Goには実行したいバイナリの絶対パスが必要です。
    // なので、`exec.LookPath`を使ってパス(おそらく`/bin/ls`)を特定します。
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

    // `Exec`にはスライスの引数(ひとつの長い文字列ではなく)を渡します。
    // `ls`にいくつかの引数を加えて実行します。最初の引数はプログラム名にします。
    args := []string{"ls", "-a", "-l", "-h"}

    // `Exec`には[環境変数](environment-variables)のセットも必要です。
    // ここでは現在の環境変数を利用します。
    env := os.Environ()

    // ここが実際の`syscall.Exec`の呼び出しです。
    // 成功すれば、このプロセスはここで終了し`/bin/ls -a -l -h`に置きかわります。
    // エラーが発生したら戻り値を取得します。
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}
