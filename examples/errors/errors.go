// Goでは慣用的に明示的な分離された戻り値でエラーを扱います。
// これはJavaやRubyのような例外を使う言語や
// C言語で使われることがあるオーバーロードされた単一の結果/エラー値
// とは対称的です。Goのアプローチはどの関数がエラーを返しそのエラーを
// エラーでない処理を行うのと同じ書き方で処理するのを簡単にします。

package main

import "errors"
import "fmt"

// 規約によりエラーは最後の戻り値になり
// 組み込みのインターフェースである`error`型を持ちます。
func f1(arg int) (int, error) {
    if arg == 42 {

        // `errors.New`は基本的な`error`値を作成します。
        // エラーメッセージを与えられます。
        return -1, errors.New("can't work with 42")

    }

    // エラーの位置が`nil`なのはエラーがないことを示します。
    return arg + 3, nil
}

// `Error()`メソッドを実装すれば`error`としてカスタムの型を定義できます。
//  ここでは上の例を変え、引数エラーであることを明示的に示すカスタム型を使います。
type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {

        // `&argError`という構文で新しい構造体を作ります。
        // そして、`arg`と`prob`フィールドに値を設定します。
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {

    // 下のふたつのループで上で定義したエラーを返す関数をテストします。
    // ループの中の`if`を使ったエラーのチェックはGoでは一般的なイディオムです。
    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    // カスタムエラーのデータをプログラムから使いたいなら
    // 型アサーション経由でカスタムエラー型のインスタンスを取得する必要があります。
    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
