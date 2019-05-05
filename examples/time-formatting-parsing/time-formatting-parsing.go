// Goはパターンベースのレイアウトで時間のフォーマットとパースを実現します。

package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    // ここでは基本的な時間のフォーマットを示します。
    // RFC3339に基づきます。対応するレイアウト定数を使っています。
    t := time.Now()
    p(t.Format(time.RFC3339))

    // 時間のパースは`Format`と同じレイアウト値を使います。
    t1, e := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
    p(t1)

    // `Format`と`Parse`は事例ベースレイアウトを使います。
    // これらのレイアウトとして多くの場合、`time`の定数を使います。
    // カスタムのレイアウトも使えます。
    // カスタムレイアウトの場合は`Mon Jan 2 15:04:05 MST 2006`を使って
    // 与えられた時間/文字列をフォーマット/パースするためのパターンを示します。
    // 年は2006であり時間は15、曜日はMonday、他の要素も正確にこの通りである必要があります。
    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2)

    // 純粋な数値表現のためには、timeの個別の値を取り出して
    // 標準的な文字列のフォーマッティングを使うこともできます。
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

    // 入力が不正な場合`Parse`はエラーを返しパースの問題を示します。
    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)
}
