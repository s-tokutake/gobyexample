// Goは[_無名関数_](https://ja.wikipedia.org/wiki/%E7%84%A1%E5%90%8D%E9%96%A2%E6%95%B0)をサポートします。
// これは<a href="https://ja.wikipedia.org/wiki/%E3%82%AF%E3%83%AD%E3%83%BC%E3%82%B8%E3%83%A3"><em>クロージャ</em></a>を作ります。
// 無名関数は名付けることなくインラインで関数を定義したい場合に便利です。

package main

import "fmt"

// この`intSeq`関数は別の関数を返します。
// `intSeq`の内部で無名で定義された関数です。
// この無名の関数は変数`i`を_閉じ込めて_クロージャを作ります。
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// `intSeq`を呼び出し、`nextInt`に設定します。
	// この関数値は`i`の値を捕捉しています。
	// この値は`nextInt`を呼び出すたびに更新されます。
	nextInt := intSeq()

	// `nextInt`を何回か呼び出してクロージャの挙動を確かめてみます。
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 状態が個別の関数でユニークであること確認するため、
	// 新しい関数値を作って実行してみます。
	newInts := intSeq()
	fmt.Println(newInts())
}
