// _range_はさまざまなデータ構造の要素を反復処理します。
// すでに学習したデータ構造に対して`range`がどのように働くか紹介します。

package main

import "fmt"

func main() {

	// ここでは`range`を使ってスライス内の数を合計しています。
	// 配列でも同様のことができます。
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 配列とスライスに対する`range`はインデックスと値の両方を提供します。
	// 上の例ではインデックスは必要ないので、ブランク識別子`_`を使って無視しました。
	// インデックスを使いたい場合もあります。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// マップに対する`range`はキー/値のペアを反復します。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// マップのキーだけを反復することもできます。
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 文字列に対する`range`ではUnicodeのコードポイントを反復します。
	// 最初の値は`rune`の開始のバイトのインデックス、ふたつ目は`rune`そのものです。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
