// Goはキャラクタ、文字列、数値、ブーリアンの_定数_をサポートします。

package main

import "fmt"
import "math"

// `const`は定数を宣言するのに使います。
const s string = "constant"

func main() {
	fmt.Println(s)

	// `const`文は`var`文が使えるところならどこでも使えます。
	const n = 500000000

	// 定数式任意の精度で計算を行います。
	const d = 3e20 / n
	fmt.Println(d)

	// 数値の定数は、例えば、明示的なキャストをして型を与えない限り、型がありません。
	fmt.Println(int64(d))

	// 型が必要な状況で数値を使うと型が与えられます。
	// 例えば、変数への割り当てや関数呼び出しです。
	// ここでの`math.Sin`は`float64`を期待しています。
	fmt.Println(math.Sin(n))
}
