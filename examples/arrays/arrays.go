// Goでは_配列_は特定の長さの要素の番号付きのシーケンスです。

package main

import "fmt"

func main() {

	// ここでは`a`という配列を作成します。`a`は5つの整数を持ちます。
	// 要素の型と長さは両方とも配列の型の一部です。
	// デフォルトでは配列はゼロ値になります。つまり、`int`なら`0`です。
	var a [5]int
	fmt.Println("emp:", a)

	// インデックスを使って`array[index] = value`という構文で値をセットできます。
	// `array[index]`で値を取得します。
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 組み込みの`len`は配列の長さを返します。
	fmt.Println("len:", len(a))

	// 1行で配列の宣言と初期化をするにはこのような書き方をします。
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 配列は1次元ですが、多次元のデータ構造を構成できます。
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
